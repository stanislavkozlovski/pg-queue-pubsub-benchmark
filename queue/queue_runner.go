package queue

import (
	"context"
	"crypto/rand"
	"database/sql"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	c "main/common"
	"sync"
	"time"
)

type QueueConfig struct {
	c.BaselineConfig
	TuneTableVacuum bool // if true, alter vacuum params on queue table
}

type QueueBenchmarkRun struct {
	config *QueueConfig
	*c.BenchmarkRun
}

func NewQueueBenchmarkRun(cfg *QueueConfig, db *sql.DB, ctx context.Context, writeLimiter *rate.Limiter) (*QueueBenchmarkRun, error) {
	return &QueueBenchmarkRun{
		config: cfg,
		BenchmarkRun: c.NewBenchmarkRun(db,
			c.NewMetrics(cfg.Writers, cfg.Readers), ctx, writeLimiter),
	}, nil
}

func (br *QueueBenchmarkRun) Setup() error {
	queries := []string{
		"DROP TABLE IF EXISTS queue_archive",
		"DROP TABLE IF EXISTS queue",
		`CREATE TABLE queue (
			id BIGSERIAL PRIMARY KEY,
			payload BYTEA NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`,
		`CREATE TABLE queue_archive (
			id BIGINT,
			payload BYTEA NOT NULL,
			created_at TIMESTAMP NOT NULL,
			read_at TIMESTAMP NOT NULL DEFAULT NOW()
		)`,
		"CREATE INDEX idx_queue_id ON queue(id)",
	}
	for _, q := range queries {
		if _, err := br.Db.ExecContext(br.Ctx, q); err != nil {
			return fmt.Errorf("setup: %w", err)
		}
	}
	// optional tuning for queue table
	if br.config.TuneTableVacuum {
		tuneSQL := `
			ALTER TABLE queue SET (
				autovacuum_vacuum_scale_factor = 0.01,
				autovacuum_vacuum_insert_threshold = 1000,
				autovacuum_analyze_scale_factor = 0.05,
				fillfactor = 90
			)`
		if _, err := br.Db.ExecContext(br.Ctx, tuneSQL); err != nil {
			return fmt.Errorf("tune-table: %w", err)
		}
		log.Println("[queue info] applied aggressive autovacuum/fillfactor tuning to queue table")
	}
	return nil
}

func (br *QueueBenchmarkRun) Writer(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, _ := br.Db.Conn(br.Ctx)
	defer conn.Close()

	hist := br.Metrics.WriterHists[id]
	payload := make([]byte, br.config.PayloadSize)
	_, _ = rand.Read(payload)

	stmt, _ := conn.PrepareContext(br.Ctx, "INSERT INTO queue (payload) VALUES ($1)")
	defer stmt.Close()

	for {
		select {
		case <-br.Ctx.Done():
			return
		default:
			// throttle globally
			if br.WriteLimiter != nil {
				if err := br.WriteLimiter.Wait(br.Ctx); err != nil {
					return
				}
			}

			start := time.Now()
			if _, err := stmt.ExecContext(br.Ctx, payload); err != nil {
				br.Metrics.AggregateWriteErrors.Add(1)
			} else {
				br.Metrics.AggregateWritesCompleted.Add(1)
			}
			lat := time.Since(start).Nanoseconds()
			if err := hist.RecordValue(lat); err != nil {
				log.Printf("[writer %d] histogram record error: %v (lat=%dns)", id, err, lat)
			}
		}
	}
}

func (br *QueueBenchmarkRun) Reader(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, _ := br.Db.Conn(br.Ctx)
	defer conn.Close()

	selectHist := br.Metrics.ReaderReadHists[id]
	e2eHist := br.Metrics.ReaderE2EHists[id]

	const selSQL = `
		SELECT id, payload, created_at
		FROM queue
		ORDER BY id
		FOR UPDATE SKIP LOCKED
		LIMIT 1
	`
	selStmt, _ := conn.PrepareContext(br.Ctx, selSQL)
	defer selStmt.Close()

	delStmt, _ := conn.PrepareContext(br.Ctx,
		"DELETE FROM queue WHERE id = $1")
	defer delStmt.Close()

	insStmt, _ := conn.PrepareContext(br.Ctx,
		"INSERT INTO queue_archive (id, payload, created_at, read_at) VALUES ($1,$2,$3,NOW())")
	defer insStmt.Close()

	for {
		select {
		case <-br.Ctx.Done():
			return
		default:
			start := time.Now()
			tx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
			if err != nil {
				br.Metrics.AggregateReadErrors.Add(1)
				continue
			}

			var (
				id64    int64
				payload []byte
				created time.Time
			)
			// step 1: read + claim the row
			if err := tx.Stmt(selStmt).QueryRowContext(br.Ctx).Scan(&id64, &payload, &created); err != nil {
				_ = tx.Rollback()
				time.Sleep(200 * time.Microsecond)
				continue
			}

			// step 2: process (noop here)
			_ = payload

			// step 3a: delete
			if _, err := tx.Stmt(delStmt).ExecContext(br.Ctx, id64); err != nil {
				_ = tx.Rollback()
				br.Metrics.AggregateReadErrors.Add(1)
				continue
			}
			// step 3b: insert into archive
			if _, err := tx.Stmt(insStmt).ExecContext(br.Ctx, id64, payload, created); err != nil {
				_ = tx.Rollback()
				br.Metrics.AggregateReadErrors.Add(1)
				continue
			}

			if err := tx.Commit(); err != nil {
				br.Metrics.AggregateReadErrors.Add(1)
				continue
			}

			br.Metrics.AggregateReadsCompleted.Add(1)
			br.Metrics.AggregateUpdatesCompleted.Add(1)
			selectLatency := time.Since(start).Nanoseconds()
			if err := selectHist.RecordValue(selectLatency); err != nil {
				log.Printf("[reader %d] histogram record error: %v (lat=%dns)", id, err, selectLatency)
			}
			e2eLatency := time.Since(created).Nanoseconds()
			if err := e2eHist.RecordValue(e2eLatency); err != nil {
				log.Printf("[reader %d] histogram record e2e latency error: %v (lat=%dns)", id, err, e2eLatency)
			}
		}
	}
}

func (br *QueueBenchmarkRun) Run() {
	var wg sync.WaitGroup
	for i := 0; i < br.config.Writers; i++ {
		wg.Add(1)
		go br.Writer(i, &wg)
	}
	for i := 0; i < br.config.Readers; i++ {
		wg.Add(1)
		go br.Reader(i, &wg)
	}
	wg.Add(1)
	go br.Reporter(&wg)

	<-br.Ctx.Done()
	time.Sleep(50 * time.Millisecond)
	wg.Wait()
}
