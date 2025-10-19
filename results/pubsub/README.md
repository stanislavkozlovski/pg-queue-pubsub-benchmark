# PubSub Results

This folder contains the results from running PubSub tests..

## Design

The PubSub table design is inspired by Kafka.
It uses gapless, monotonically-increasing offsets and supports at-least-once/at-most-once processing. This test in particular uses at-least-once semantics due to the way reads are handled, but neither solution should impact the benchmark results.

The design uses three tables:
- `topic_partition` - the data tables where the data is written to, and read from. Multiple of these can exist.
- `log_counter` - a table used to synchronize and track the latest offset per `topic_partition` table. Writers increment the latest offset here while writing.
- `consumer_offsets` - a table used to track the latest offset read per consumer group for a particular topic partition.

It does NOT use the `NOTIFY/LISTEN` statements. Those statements are [lossy](https://lobste.rs/s/rk3eft/choose_postgres_queue_technology#c_iix4ph). They mainly serve an optimization role in Postgres which reduces the time to read a new record (i.e. instead of polling every 500ms, immediately get notified).

### Tables

```sql
CREATE TABLE log_counter (
  id           INT PRIMARY KEY,     -- topicpartition table name id
  next_offset  BIGINT NOT NULL      -- next offset to assign
);

for i in NUM_PARTITIONS:
  CREATE TABLE topicpartition%d (
    id          BIGSERIAL PRIMARY KEY,
    -- strictly increasing offset (indexed by UNIQUE)
    c_offset    BIGINT UNIQUE NOT NULL,
    payload     BYTEA NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
  );
  INSERT INTO log_counter(id, next_offset) VALUES (%d, 1);

CREATE TABLE consumer_offsets (
  group_id     TEXT NOT NULL,     -- consumer group identifier
  -- topic-partition id (matches log_counter.id / topicpartitionN)
  topic_id     INT  NOT NULL,
  -- next offset the consumer group should claim
  next_offset  BIGINT NOT NULL DEFAULT 1,
  PRIMARY KEY (group_id, topic_id)
);
```

### Writes

Each writer goroutine loops and atomically inserts `$BATCH_SIZE` records in a given topic partition table, while updating the latest offset:
```sql
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
```
Each writer is sticky to one topic-partition table in the lifetime of the test. In aggregate, all writers produce to all tables.
Each message written therefore has a unique monotonically-increasing offset number associated with it.

### Reads

Reads are conceptually very similar to Kafka:
- Readers consume the `topicpartition` table(s) sequentially. They start from the lowest offset and read progressively.
- Readers are split into consumer groups. Each group makes separate, independent reads and progress on the `topicpartition` tables.
- Each consumer group contains 1 reader per `topicpartition` table.
- Readers store their progress in a `consumer_offsets` table, with a row for each `topicpartition,group` pair.
- Each reader updates the latest processed offset (claiming the records via a lock on the group's offset marker), selects the records and processes them inside a single transaction.

Each reader is assigned a particular consumer group and partition. The group as a whole reads all partitions while each reader in the group reads only one partition at a time.
The reader loops, opens a transaction, optimistically claims `$BATCH_SIZE` records (by advancing the offset mark beyond them), selects them and processes the records.
If successful, it commits the transaction and through that advances the offset for the group.

It is a pull-based read (just like Kafka), rather than push-based. If the reader has no records to poll, it sleeps for a bit.
