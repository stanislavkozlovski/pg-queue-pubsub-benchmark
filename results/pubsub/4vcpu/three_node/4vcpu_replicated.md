# 4 vCPU Tri-Node Results (2x 5-min test)

See [run details](./run.md) for the full benchmark run output.

- ✍️ write: 4.90 MiB/s (5,015 msg/s)
- 📖️ read: 24.49 MiB/s (25,073 msg/s)
- write latency:
  - p99: 153.45 ms
  - p95: 6.77 ms
  - p50: 5.01 ms
- read latency:
  - p99: 56.69 ms
  - p95: 4.91 ms
  - p50: 3.89 ms
- end-to-end latency:
  - p99: 185.37 ms
  - p95: 12.35 ms
  - p50: 8.89 ms

- test parameters:
  - write batch size: 100 records
  - read batch size: 200 records
  - 4 tables (partitions)
  - 10 writers (2 writers per partition on average)
  - 5x read fanout - 5 consumer groups with 4 readers each
    - 20 reader total

- primary server: `c7i.xlarge`
  - gp3 25GB 9000 IOPS
  - Ubuntu 24.04
  - few custom postgres settings
  - CPU usage: ~65% CPU
  - us-east-1a

The custom postgres settings:
```bash
wal_sync_method=open_datasync # minor potential optimization I toyed around with. Not certain it helped at all
wal_compression=on
hot_standby=on
hot_standby_feedback=on
synchronous_standby_names='FIRST 1 (pg2, pg3)'
wal_log_hints=on # guarantees correctness in failover
max_worker_processes=32
max_parallel_workers=32
max_parallel_workers_per_gather=16
```

- pg2 streaming server: `c7i.xlarge`
    - same gp3 25GB 9000 IOPS, Ubuntu 24.04, few custom postgres settings
    - `sync` replica (synchronously replicating and waiting to ack)
    - us-east-1b
- pg3 streaming server: `c7i.xlarge`
    - same gp3 25GB 9000 IOPS, Ubuntu 24.04, few custom postgres settings
    - `potential` replica (**a**synchronously replicating but on standby to become `sync` if the other one dies)
    - us-east-1c
- test driver client: `c7i.xlarge`
    - 10 writers
    - 20 readers
    - 1 KB payload
    - 5 minute duration
    - us-east-1a

Notes:
- As expected, latency was impacted. Our p99 e2e latency 3x'd (60ms vs 185ms), but interestingly the p50 didn't budge.
- Throughput was not impacted. ~4.9 MiB/s was sustained!
- This replication setup in Postgres is the equivalent of something like a RF=2.5 in Kafka.
  - The client receives a response when the one `sync` replica confirms the change.
  - The other `potential` replica is replicating asynchronously (not blocking the write path) but will become promoted to `sync` if the sync one was to die.

# Run

```bash
./pg_msg_bench --host=$HOST --port=5432 --db=benchmark --user=postgres --password=postgres --report=5s \
  --write-batch 100 --read-batch 200 \
  --writers=10 --readers=4 --consumer-groups 5 \
  --mode=pubsub --partitions 4 \
  --duration=300s --throttle_writes=5000
```

# 💸 Cost

The total cost of this setup is around **\$11,514** per year. 

- The networking cost for replication is **\$6036** per year. (4.9 MiB/s replicated 2x across availability zones).
  - Assuming your clients are in the same availability zone as your primary (or the DB they're writing to/reading from), no extra networking costs exist.
- The compute and disk cost is **\$5478** per year (see [4vCPU README](../single_node/4vcpu.md#-cost) for in-depth calculations per node; we run 3 nodes)
