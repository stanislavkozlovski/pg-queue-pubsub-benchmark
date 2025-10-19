# Pub Sub

To emulate Pub/Sub messaging in Postgres, we do the following:
- one source of truth `datalog` table contains all the data
- writes constantly append to the `datalog` table in batches of 50 messages
- readers read batches sequentially from the table `SELECT * where id > 50 AND id < 100`;
- readers keep tabs on what the latest read ID is, and consume from that ID one batch forward
- readers are split into consumer groups. Each group runs N readers and shares their latest read ID
- each group maintains a single row in a `consumer_offsets` table which denotes their latest read ID
- readers within a group lock that offset row and optimistically update it `UPDATE start_offset, end_offset VALUES (51, 100) WHERE group='group1'`
  - (this is at-most-once. for at-least-once, we could use a claim table to store claims and retry if failed)

The read fanout scales by adding more groups. One group should have enough readers to be able to keep up with the tail of the log.

The pseudocode is the following:

#### The Tables
```SQL
-- === core log ===
CREATE TABLE topicpartition (
  id          BIGSERIAL PRIMARY KEY,                  -- physical row id (not used for reads)
  offset      BIGINT UNIQUE NOT NULL,                 -- gapless, strictly increasing, read-key
  payload     BYTEA NOT NULL,
  created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

-- fast range scans on the logical log
CREATE UNIQUE INDEX idx_topicpartition_offset ON topicpartition(offset);


-- === txn-safe global counter for writers ===
CREATE TABLE log_counter (
                             id           INT PRIMARY KEY CHECK (id = 1),
                             next_offset  BIGINT NOT NULL
);
INSERT INTO log_counter(id, next_offset) VALUES (1, 1);

-- === per-group progress ===
CREATE TABLE consumer_offsets (
  group_id     TEXT PRIMARY KEY,
  next_offset  BIGINT NOT NULL DEFAULT 1,             -- next offset to claim/read
);
```
#### Writer
We use a separate table to ensure offsets are gapless. The naive way of using a PG id can fail because it can have gaps (eg if tx fails)
```SQL
-- writer (you know batch_size + values at insert time)
BEGIN;

WITH reserve AS (
UPDATE log_counter
SET next_offset = next_offset + :batch_size
WHERE id = 1
    RETURNING (next_offset - :batch_size) AS first_off
)
INSERT INTO topicpartition(offset, payload)
SELECT r.first_off + p.ord - 1, p.payload
FROM reserve r,
     unnest(:payloads) WITH ORDINALITY AS p(payload, ord);  -- ord = 1..batch_size

COMMIT;
```

#### Reader

KEY: We don't expect any gaps in offsets because of the writer's TX

What we do is claim a range of offsets inside Postgres and atomically update the last consumer group offset.
We claim either the batch size, or up to the latest written offset (whichever is smaller).
We get the latest written offset from the log_counter table
```
const group = "g1"
const batch = B

loop:
  -- inputs: :consumer_group_id, :max_batch_size
  BEGIN;
  
  WITH counter_tip AS (
    -- O(1): writers advance this in their tx; readers just peek
    SELECT (next_offset - 1) AS highest_committed_offset
    FROM log_counter
    WHERE id = 1
  ),
  claimed_range AS (
    UPDATE consumer_offsets c
    SET next_offset = c.next_offset
                    + LEAST(
                        :max_batch_size,
                        GREATEST(
                          0,
                          (SELECT highest_committed_offset FROM counter_tip) - c.next_offset + 1
                        )
                      )
    WHERE c.group_id = :consumer_group_id
    RETURNING
      /* compute the exact claimed window */
      (next_offset - LEAST(
         :max_batch_size,
         GREATEST(0, (SELECT highest_committed_offset FROM counter_tip) - next_offset + :max_batch_size)
       )) AS claimed_start_offset,
      (next_offset - 1) AS claimed_end_offset
  )
  SELECT claimed_start_offset, claimed_end_offset
  FROM claimed_range;
  
  COMMIT;


  continue;
```
