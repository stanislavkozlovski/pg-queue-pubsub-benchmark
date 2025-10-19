package pubsub

import (
	"crypto/rand"
	"database/sql"
	"fmt"
	"github.com/lib/pq"
	"log"
	"sync"
	"time"
)

// this SQL ensures an insert with contiguous offsets
// the log_counter table row acts as the serialization point for writers
const insertWithReserveSQL = `
WITH reserve AS (
  UPDATE log_counter
  SET next_offset = next_offset + $1
  WHERE id = $3::int
  RETURNING (next_offset - $1) AS first_off
)

INSERT INTO topicpartition%d(c_offset, payload)
SELECT r.first_off + p.ord - 1, p.payload
FROM reserve r,
     unnest($2::bytea[]) WITH ORDINALITY AS p(payload, ord);
`

func (br *PubSubBenchmarkRun) Writer(id int, partitionID int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, _ := br.Db.Conn(br.Ctx)
	defer conn.Close()

	hist := br.Metrics.WriterHists[id]

	batchSize := br.config.WriteBatchSize
	if batchSize <= 0 {
		batchSize = 1
	}

	// prepare once on the base db; bind to tx with tx.Stmt(...) each loop
	insertSQL := fmt.Sprintf(insertWithReserveSQL, partitionID)
	stmt, err := br.Db.PrepareContext(br.Ctx, insertSQL)
	if err != nil {
		log.Printf("[writer %d] prepare failed: %v", id, err)
		return
	}
	defer stmt.Close()

	for {
		select {
		case <-br.Ctx.Done():
			return
		default:
			// throttle globally per batch
			if br.WriteLimiter != nil {
				if err := br.WriteLimiter.WaitN(br.Ctx, batchSize); err != nil {
					return
				}
			}

			payloads := buildPayloadBatch(batchSize, br.config.PayloadSize)

			tx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
			if err != nil {
				br.Metrics.AggregateWriteErrors.Add(int64(batchSize))
				continue
			}
			start := time.Now()
			_, execErr := tx.StmtContext(br.Ctx, stmt).ExecContext(
				br.Ctx,
				batchSize,          // $1 - the batch size
				pq.Array(payloads), // $2 :: bytea[] - the payload
				partitionID,        // $3 - the partition (table) ID
			)
			if execErr != nil {
				_ = tx.Rollback()
				br.Metrics.AggregateWriteErrors.Add(int64(batchSize))
				continue
			}
			if err := tx.Commit(); err != nil {
				br.Metrics.AggregateWriteErrors.Add(int64(batchSize))
				continue
			}

			// success - up the metrics
			br.Metrics.AggregateWritesCompleted.Add(int64(batchSize))
			if err := hist.RecordValue(time.Since(start).Nanoseconds()); err != nil {
				log.Printf("[writer %d] hist record err: %v", id, err)
			}
		}
	}
}

// helper: build N random payloads ([][]byte) of size payloadSize
func buildPayloadBatch(batchSize, payloadSize int) [][]byte {
	out := make([][]byte, batchSize)
	for i := 0; i < batchSize; i++ {
		buf := make([]byte, payloadSize)
		_, _ = rand.Read(buf)
		out[i] = buf
	}
	return out
}
