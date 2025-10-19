# 4 vCPU Single-Node Results (2x 15-min tests)

See [run details](./runs.md) for the full benchmark run output.

- ✍️ **write:** 2.82 MiB/s (2,885 msg/s)
- 📖️ **read:** 2.82 MiB/s (2,885 msg/s)
- **CPU:** ~60%
- write latency:
  - p99: 2.46 ms
  - p95: 2.27 ms
  - p50: 1.77 ms
- read latency:
  - p99: 4.20 ms
  - p95: 3.64 ms
  - p50: 3.37 ms
- end-to-end latency:
  - p99: 15.86 ms
  - p95: 5.74 ms
  - p50: 5.07 ms
- bottleneck: number of clients; more clients raises CPU; this number of clients can't push more throughput.

- server: `c7i.xlarge`
    - gp3 25GB 9000 IOPS
    - Ubuntu 24.04
    - absolute default postgres settings
    - CPU usage: ~60% CPU
    - us-east-1a
- test driver client: `c7i.xlarge`
    - 10 writers
    - 15 readers
    - 1 KB payload
    - 15 minute duration
    - us-east-1a

# Run

```bash
 ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  \
 --writers=10   --readers=15   --duration=1800s   --payload=1024   --report=5s \
 --throttle_writes=2900   --mode=queue
```

Notes:
- CPU bottlenecks due to number of readers/writers. With 50/50 read/write connections, it pegged to 100% CPU. At 100% CPU I could get ~4 MB/s out of the machine. Queue depth didn't increase either. But it was pegged at 100%, which is against my test goal. 
- Hence, at this 4 vCPU size without tuning or a connection pooler - we need less clients. We opted for 10 writers and 15 readers as a sweet spot. CPU kept at 60%.
- The bottleneck then became the number of **reader clients**. Since a reader client's p50 read latency is roughly 5.4ms, it can only read 185 messages a second. 15 clients can't read more than 3145 msg/s.
- In general, `iostat` showed large write amplification on the disk. A ~2.85 MB/s workload would write ~18 MB/s to the disk. I can't explain it. 
- 3000k IOPS on gp3 was low, so I increased it to 9000. `iostat` still showed high util percentage, but didn't seem to be anywhere close to bottleneck.

# 💸 Cost

The cost of the server is around **\$1626** per year:
- c7i.xlarge for \$1033.68/yr (at \$0.118/h 1yr reserved pricing)
- a 242GB gp3 disk for \$592.32/yr
  - gp3 storage costs is \$232.32/yr (19.36/m at $0.08/GB-month)
  - gp3 IOPS cost is  (6000 * 0.005/provisioned IOPS-month  is \$360/yr at \$30/m)
    We are assuming a ~12h message retention. The rest can be offloaded to S3 (or deleted). The S3 storage cost we don't count.
    In reality, data can be offloaded to S3 faster (e.g. by the hour).
