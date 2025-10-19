# Queue Results

This folder contains the results from running Queue tests. (insert + claim rows with `FOR UPDATE SKIP LOCKED`).

## Design

A simple Queue design would use a `BOOLEAN is_read` field to denote whether a record is read (popped off the queue) or not.
In my early, early testing, I found this lowers performance. The EXPLAIN showed that each read query started paying a lot of latency to filter through many pages of already `is_read=true` rows.
Vacuum apparently takes time to prune them away, and until it does - reads were slowed. A more performant approach I found was to delete rows once read.  This keeps the table small and performant.

Since we don't necessarily want to omit the data (many companies have a policy of never deleting, at least not before a backup) - the benchmark moves the data to another table in a sort of archival step.
While deleting data may be the right approach for many, and the extra archival insert in the transaction does slow down performance, I found this tradeoff conservatively maintained the test scenario more realistic.

In my simple tests, I find this approach increases read throughput by **50%** at larger scale.

### Tables

We use a simple `queue` table. When an element is processed off the queue, it's moved into the archive table.

```sql
CREATE TABLE queue (
  id BIGSERIAL PRIMARY KEY,
  payload BYTEA NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
)

CREATE TABLE queue_archive (
  id BIGINT,
  payload BYTEA NOT NULL,
  created_at TIMESTAMP NOT NULL, -- ts the event was originally created at
  processed_at TIMESTAMP NOT NULL DEFAULT NOW() -- ts the event was processed at
)
```

### Writers

The benchmark runs `N` writer goroutines. These represent writer clients.
Each one simply loops and sequentially inserts a single random item into the table:

```sql
INSERT INTO queue (payload) VALUES ($1)
```

It only inserts one message per statement.

### Reads

The benchmark also runs `N` reader goroutines. Each reader loops and processes one message.
The processing is done inside a database transaction.

```sql
BEGIN;

SELECT id, payload, created_at
  FROM queue
  ORDER BY id
  FOR UPDATE SKIP LOCKED
  LIMIT 1;

-- Your business code "processes" the message. In the benchmark, it's a no-op.

DELETE FROM queue WHERE id = $1;

INSERT INTO queue_archive (id, payload, created_at, processed_at)
  VALUES ($1,$2,$3,NOW());

COMMIT;
```
Each reader again works only one message at a time per transaction.


## Throttle Writes

In my tests, I found that writes massively outpace reads. I would get results like ~14 MB/s writes and ~4 MB/s reads, which resulted in the system perpetually accumulating an ever-greater backlog.  I did not investigate the reason for this too deeply - it may be my design.
To avoid this, throughout the tests, I throttle the write rate to a certain rate in order to avoid it being unlimited.
I determine this goldilocks rate haphazardly by running a few tests at different rates until I find a reasonable one where reads/writes are kept steady at 1:1.
