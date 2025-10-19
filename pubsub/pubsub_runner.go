package pubsub

import (
	"context"
	"database/sql"
	"fmt"
	"golang.org/x/time/rate"
	"log"
	c "main/common"
	"sync"
	"time"
)

type PubSubConfig struct {
	c.BaselineConfig

	// Readers denotes the number of subscribers PER group
	NumConsumerGroups int

	// The number of topicpartition tables to exist
	NumPartitions int

	// the max number of messages to read per request
	// in Kafka, the defaults are `fetch.max.bytes` (50MB) and `max.partition.fetch.bytes` (1MB)
	ReadBatchSize int

	// the number of messages per read/write request. In Kafka, linger.ms and batch.size (16 KB) control this per partition.
	// Its 16 KiB per producer per partition by default too, but increasing it is recommended in general
	WriteBatchSize int
}

type PubSubBenchmarkRun struct {
	config *PubSubConfig
	*c.BenchmarkRun
	PubSubMetrics *PubSubMetrics
}

func NewPubSubBenchmarkRun(cfg *PubSubConfig, db *sql.DB, ctx context.Context, writeLimiter *rate.Limiter) (*PubSubBenchmarkRun, error) {
	metrics := NewPubSubMetrics(cfg.NumConsumerGroups, cfg.Readers, cfg.NumPartitions)
	return &PubSubBenchmarkRun{
		config: cfg,
		BenchmarkRun: c.NewBenchmarkRun(db,
			c.NewMetrics(cfg.Writers, cfg.Readers), ctx, writeLimiter),
		PubSubMetrics: metrics,
	}, nil
}

// Run enforces 1 consumer per partition (per group).
// - readers: for each group, spawn exactly NumPartitions readers; reader `i` is pinned to partition `i` (1-based).
// - writers: round-robin assign partitions; log distribution after spawning.
func (br *PubSubBenchmarkRun) Run() {
	cfg := br.config

	// hard guard: 1 consumer per partition per group
	if cfg.Readers != cfg.NumPartitions {
		log.Fatalf("invalid config: Readers per group (%d) must equal NumPartitions (%d) to ensure 1 consumer per partition",
			cfg.Readers, cfg.NumPartitions)
	}

	var wg sync.WaitGroup

	// --- spawn readers, pinned 1:1 to partitions per group ---
	for groupID := 0; groupID < cfg.NumConsumerGroups; groupID++ {
		gm := br.PubSubMetrics.Groups[groupID]
		for p := 1; p <= cfg.NumPartitions; p++ {
			consumerID := p - 1 // keep 0-based consumer index if you rely on it elsewhere
			wg.Add(1)
			// groupID, gm, consumerID, partitionID, wg, kafkaSemantics
			go br.GroupMember(groupID, gm, consumerID, p, &wg)
		}
	}

	// --- spawn writers, round-robin partitions ---
	prodPerPart := make([]int, cfg.NumPartitions)
	for w := 0; w < cfg.Writers; w++ {
		pid := (w % cfg.NumPartitions) + 1 // 1-based
		prodPerPart[pid-1]++
		wg.Add(1)
		go br.Writer(w, pid, &wg)
	}
	time.Sleep(1000 * time.Millisecond)
	log.Printf("[pub info] spawned readers")

	// summary log: producers per partition in one line
	summary := ""
	for i, cnt := range prodPerPart {
		if i > 0 {
			summary += " "
		}
		summary += fmt.Sprintf("p%d=%d", i+1, cnt)
	}
	log.Printf("[pub info] producers per partition [%s]", summary)

	// reporter
	wg.Add(1)
	go br.Reporter(&wg)

	<-br.Ctx.Done()
	time.Sleep(50 * time.Millisecond)
	wg.Wait()
}
