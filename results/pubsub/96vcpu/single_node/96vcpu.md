# 96 vCPU Single-Node Results (3x 2-min tests)

See [run details](./runs.md) for the full benchmark run output.

- ✍️ **write**: 238 MiB/s (243k msg/s)
- 📖️ **read**: 1.16 GiB/s (1.2M msg/s)
- write latency:
  - p99: 137.78 ms
  - p95: 47.15 ms
  - p50: 10.61 ms
- read latency:
  - p99: 24.61 ms
  - p95: 19.50 ms
  - p50: 10.45 ms
- end-to-end latency:
  - p99: 853.37 ms
  - p95: 242.13 ms
  - p50: 23.43 ms
- bottleneck: This is probably the number of partitions. 8 MiB/s of writes per partition (8000 msg) sounds like a potential limit of a single table with this design
    - Reads could have been scaled much further I think. (up until 100% CPU). I tried with 10x fanout and had no issues, but figured I don't need to push that extreme.

- test parameters:
  - write batch size: 200 records
  - read batch size: 200 records
  - 30 tables (partitions)
  - 100 writers (~3.33 writers per partition on average)
  - 5x read fanout - 5 consumer groups with 30 readers each
    - 150 reader total


- server: `c7i.24xlarge` (96 vCPU, 192 GiB RAM)
    - io2 50GB /w 12_000 IOPS
    - Ubuntu 24.04
    - modified (upscaled) Postgres settings with huge_pages enabled
    - CPU usage: **~10%** CPU (completely unused)
    - us-east-1a

- test driver client: `c7i.4xlarge`
    - 100 writers
    - 200 readers
    - 1 KB payload
    - 2 minute duration
    - us-east-1a


In this test, it was time to reconfigure postgres. It's well known that its default settings are very conservative:
> The Postgres defaults are very conservative. It will start even on a tiny machine like Raspberry Pi. Which is great! The flip side is it’s terrible for typical database servers which tend to have much more RAM/CPU.
> To get good performance on such larger machines, you need to adjust a couple parameters (shared_buffers, max_wal_size, …).

We are running on a 96 vCPU machine with 128GB of RAM after all. I talked to ChatGPT a ton and it gave me the following configs:
- [postgres config](./modified_postgresql.conf) - it suggested I substantially increase some settings like `shared_buffers` and `max_worker_processes`.
- [huge pages](./change_configs.md#5-disable-transparent-huge-pages-in-the-kernel) - it suggested I modify the kernel's huge page count and configurations
The way I applied these changes can be found in the [change configs script](./change_configs.md).

# Run

```bash
./pg_msg_bench --host=$HOST --port=5432 --db=benchmark --user=postgres --password=postgres \
  --report=5s  --write-batch=200 --read-batch=200 \
   --writers=100 --readers=30 --consumer-groups=5 \
   --mode=pubsub --partitions=30 \
   --duration=120s --throttle_writes=250000
```