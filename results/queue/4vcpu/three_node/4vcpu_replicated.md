# 4 vCPU Tri-Node Results (1x 10-min test)

See [run details](./run.md) for the full benchmark run output.

- ✍️ **write:** 2.34 MiB/s (2,397 msg/s)
- 📖️ **read:** 2.34 MiB/s (2,397 msg/s)
- **CPU:** ~65%
- write latency:
    - p99: 3.3 ms
    - p95: 2.6 ms
    - p50: 2 ms
- read latency:
    - p99: 7.65 ms
    - p95: 6.95 ms
    - p50: 5.8 ms
- end-to-end latency:
    - p99: 920 ms (see [High E2E Latency Note](../../IMPERFECTIONS.md#high-e2e-latency-note))
    - p95: 536 ms
    - p50: 7 ms

- primary server: `c7i.xlarge`
    - gp3 25GB 9000 IOPS
    - Ubuntu 24.04
    - few custom postgres settings
    - CPU usage: ~60% CPU
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
    - 15 readers
    - 1 KB payload
    - 10 minute duration
    - us-east-1a

Notes:
- As expected, throughput and latency were impacted. Our p50 read nearly 2x'd given the replication (3.37ms vs 5.8ms here)
  - This limits the max throughput we can achieve with the given set of clients and CPU usage.
  - See the single-node test for more notes regarding that [here](../single_node/4vcpu.md).
- This replication setup in Postgres is the equivalent of something like an RF=2.5 in Kafka.
  - The client receives a response when the one `sync` replica confirms the change.
  - The other `potential` replica is replicating asynchronously (not blocking the write path) but will become promoted to `sync` if the sync one was to die. 

# Run

```bash
./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres \
   --writers=10   --readers=15   --duration=600s   --payload=1024   --report=5s   \
   --throttle_writes=2400   --mode=queue
```


# 💸 Cost

The cost of the servers and EBS is around **\$4878** per year. (see [4vCPU README](../single_node/4vcpu.md#-cost) for in-depth calculations)
The networking cost for replication is **\$2883** per year (2.34 MiB/s replicated 2x across availability zones).
Assuming your clients are in the same availability zone as your primary, the total cost would be **\$7761** per year.