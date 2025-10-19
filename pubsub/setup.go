package pubsub

import (
	"fmt"
	"log"
)

// Setup prepares the schema for a pub/sub style benchmark.
//
// It creates N independent "topic partitions", each with:
//   - A dedicated table (topicpartition1, topicpartition2, … topicpartitionN)
//   - A log_counter row to track the next offset to assign
//
// The schema is:
//
//	topicpartitionN:   the actual log table, append-only (c_offset, payload, created_at)
//	log_counter:       global table w/ one row per partition, tracking next_offset
//	consumer_offsets:  consumer groups' positions, keyed by (group_id, topic_id)
//
// Example: Setup(10) creates 10 topicpartition tables + 10 log_counter rows
//
//	consumer_offsets then tracks group progress per topicpartition
func (br *PubSubBenchmarkRun) Setup(numPartitions int) error {
	queries := []string{
		// Drop global tables first
		"DROP TABLE IF EXISTS consumer_offsets CASCADE;",
		"DROP TABLE IF EXISTS log_counter CASCADE;",
	}

	// Drop all topicpartitionN tables if they exist
	for i := 1; i <= numPartitions; i++ {
		queries = append(queries, fmt.Sprintf("DROP TABLE IF EXISTS topicpartition%d CASCADE;", i))
	}

	// log_counter: one row per topicpartition (id = partition id)
	queries = append(queries, `
CREATE TABLE log_counter (
  id           INT PRIMARY KEY,     -- partition id
  next_offset  BIGINT NOT NULL      -- next offset to assign
);`)

	// Create each topicpartitionN table + its log_counter row
	for i := 1; i <= numPartitions; i++ {
		queries = append(queries, fmt.Sprintf(`
CREATE TABLE topicpartition%d (
  id          BIGSERIAL PRIMARY KEY,        -- physical row id
  c_offset    BIGINT UNIQUE NOT NULL,       -- strictly increasing offset (indexed by UNIQUE)
  payload     BYTEA NOT NULL,               -- message payload
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);`, i))

		// Initialize log_counter row for this partition
		queries = append(queries,
			fmt.Sprintf("INSERT INTO log_counter(id, next_offset) VALUES (%d, 1);", i),

			// More aggressive analyze so the planner sees fresh stats
			fmt.Sprintf(`ALTER TABLE public.topicpartition%d SET (
  autovacuum_analyze_scale_factor = 0.05
);`, i))
	}

	// consumer_offsets: now tracks offsets per (group_id, topic_id)
	// instead of one offset per group across all topics
	queries = append(queries, `
CREATE TABLE consumer_offsets (
  group_id     TEXT NOT NULL,     -- consumer group identifier
  topic_id     INT  NOT NULL,     -- partition id (matches log_counter.id / topicpartitionN; 1-based)
  next_offset  BIGINT NOT NULL DEFAULT 1, -- next offset this group should claim
  PRIMARY KEY (group_id, topic_id)
);`)

	// Execute everything
	for _, q := range queries {
		if _, err := br.Db.ExecContext(br.Ctx, q); err != nil {
			return fmt.Errorf("setup: %w (query=%s)", err, q)
		}
	}

	log.Printf("[pub info] successfully created %d topicpartitions w/ counters + consumer_offsets", numPartitions)
	return nil
}
