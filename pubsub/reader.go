package pubsub

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

// consumers get the latest offset via the same log_counter table
// then atomically claim a range of messages via consumer_offsets
// This has at least once semantics
// It returns the start/end offset that the consumer just claimed.
// The set start_offset
const claimSQL = `
-- on first read, 0 is the highest committed offset (1(next_offset)-1)
WITH counter_tip AS (
  SELECT (next_offset - 1) AS highest_committed_offset
  FROM log_counter
  WHERE id = $3::int   -- partition id passed i
),

    
to_claim AS (
  SELECT
    c.group_id,
    c.next_offset      AS n0,   -- old start offset pointer before update
    -- take the min of the batch size or the current offset delta
    LEAST(
      $2::bigint,
      -- how many are available to claim?
      -- on first read: 1-1+1 = 1
      -- on second reads (imagine 100 new records are in): 101 - 2 + 1 = 100 delta
      GREATEST(0, (SELECT highest_committed_offset FROM counter_tip) - c.next_offset + 1)
    ) AS delta
  FROM consumer_offsets c
  WHERE c.group_id = $1::text AND c.topic_id = $3::int
  FOR UPDATE
),

upd AS (
  UPDATE consumer_offsets c
  -- on first read: 1 + 1 = 2
  -- on second read: 2 + 100 = 102
  SET next_offset = c.next_offset + t.delta
  FROM to_claim t
  WHERE c.group_id = t.group_id AND c.topic_id = $3::int
  -- on first read, this will be 1,0; 1 > 0, so it's an empty claim and there's nothing to read.
  -- if there was one record, it'd be 1,1, which would then have us fetch with ID between 1 and 1 -- i.e. we'd get that record
  RETURNING
    t.n0                        AS claimed_start_offset,  -- start = the old next_offset
    (c.next_offset - 1)          AS claimed_end_offset    -- end   = new pointer - 1
)

SELECT claimed_start_offset, claimed_end_offset
FROM upd;
`

// partitionID formatted in runtime
const readSQL = `SELECT c_offset, payload, created_at
		   	FROM topicpartition%d
		  	WHERE c_offset BETWEEN $1 AND $2
		  	ORDER BY c_offset`

// GroupMember runs the read loop for a single subscriber in a consumer group.
// groupID is the unique group identifier.
// kafkaSemantics toggles whether to use At-Least-Once and Strict Order in processing;
func (br *PubSubBenchmarkRun) GroupMember(groupID int, gm *GroupMetrics, consumerID int, partitionID int, wg *sync.WaitGroup) {
	defer wg.Done()

	conn, err := br.Db.Conn(br.Ctx)
	if err != nil {
		log.Printf("[consumer g%d r%d p%d] get conn err: %v", groupID, consumerID, partitionID, err)
		return
	}
	defer conn.Close()

	groupKey := fmt.Sprintf("g%d", groupID)

	err = br.ensureConsumerGroupRow(conn, groupKey, partitionID)
	if err != nil {
		panic("ensure group row failed: " + err.Error())
	}
	claimStmt, err := conn.PrepareContext(br.Ctx, claimSQL)
	if err != nil {
		panic("prepare claimStmt failed: " + err.Error())
	}
	defer claimStmt.Close()

	partitionReadSQL := fmt.Sprintf(readSQL, partitionID)
	readStmt, err := conn.PrepareContext(br.Ctx, partitionReadSQL)
	if err != nil {
		panic("prepare readStmt failed: " + err.Error())
	}
	defer readStmt.Close()

	for {
		select {
		case <-br.Ctx.Done():
			return
		default:
			br.kafkaSemanticRead(conn, gm, groupKey, consumerID, partitionID, claimStmt, readStmt)
		}
	}
}

// kafkaSemanticRead does the Offset Claim + Select in one transaction.
// This gives you at-least-once semantics and strict ordering.
// It doesn't make much sense to run a lot of consumers in a group with this mode.
func (br *PubSubBenchmarkRun) kafkaSemanticRead(conn *sql.Conn, gm *GroupMetrics, groupKey string, consumerID int, partitionID int, claimOffsetsStmt *sql.Stmt, readDataStmt *sql.Stmt) {
	tx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		gm.ReadErrors.Add(1)
		return
	}

	var startOff, endOff sql.NullInt64
	if err := tx.StmtContext(br.Ctx, claimOffsetsStmt).QueryRowContext(br.Ctx,
		groupKey, br.config.ReadBatchSize, partitionID).Scan(&startOff, &endOff); err != nil {
		_ = tx.Rollback()
		log.Printf("[consumer %s r%d p%d] claim err: %v (group=%s batch=%d)",
			groupKey, consumerID, partitionID, err, groupKey, br.config.ReadBatchSize)
		gm.ClaimErrors.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 800*time.Microsecond))
		return
	}

	if !startOff.Valid || !endOff.Valid || startOff.Int64 > endOff.Int64 {
		// nothing to claim
		_ = tx.Commit()
		gm.EmptyClaims.Add(1)
		time.Sleep(jitter(100*time.Microsecond, 2*time.Millisecond))
		return
	}

	// read the claimed range
	start := time.Now()
	rows, err := tx.StmtContext(br.Ctx, readDataStmt).QueryContext(br.Ctx, startOff.Int64, endOff.Int64)
	if err != nil {
		_ = tx.Rollback()
		gm.ReadErrors.Add(1)
		return
	}
	defer rows.Close()

	var count int64
	for rows.Next() {
		var off int64
		var payload []byte
		var created time.Time
		if err := rows.Scan(&off, &payload, &created); err != nil {
			log.Printf("[consumer %s r%d p%d] scan err: %v", groupKey, consumerID, partitionID, err)
			continue
		}
		count++

		// **processes the payload**

		e2eLat := time.Since(created).Nanoseconds()
		_ = gm.ReaderE2EHists[consumerID].RecordValue(e2eLat)
	}

	if err := tx.Commit(); err != nil {
		gm.ReadErrors.Add(1)
		return
	}

	// record metrics
	gm.ReadsCompleted.Add(count)
	gm.UpdatesCompleted.Add(count)
	gm.PolledRecords.Add(count)
	gm.Polls.Add(1)
	br.Metrics.AggregateReadsCompleted.Add(count)
	lat := time.Since(start).Nanoseconds()
	_ = gm.ReaderReadHists[consumerID].RecordValue(lat)
}

func jitter(min, max time.Duration) time.Duration {
	if max <= min {
		return min
	}
	// fast xorshift
	r := time.Now().UnixNano()
	r ^= r << 13
	r ^= r >> 7
	r ^= r << 17
	span := int64(max - min)
	if span <= 0 {
		return min
	}
	return min + time.Duration(r%span)
}

// EnsureConsumerGroupRow makes sure a row exists in consumer_offsets for this group.
// It creates one with next_offset = 1 if it doesn't already exist.
// Safe to call multiple times; concurrent calls will not conflict.
func (br *PubSubBenchmarkRun) ensureConsumerGroupRow(conn *sql.Conn, groupKey string, partitionID int) error {
	tx, err := conn.BeginTx(br.Ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() { _ = tx.Rollback() }() // rollback is no-op if committed

	_, err = tx.ExecContext(br.Ctx, `
INSERT INTO consumer_offsets (group_id, topic_id, next_offset)
VALUES ($1::text, $2::int, 1)
ON CONFLICT (group_id, topic_id) DO NOTHING;
	`, groupKey, partitionID)
	if err != nil {
		return fmt.Errorf("insert consumer_offsets: %w", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}
