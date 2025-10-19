# 4 vCPU Single-Node Results (3x 2-min tests)

See [run details](./runs.md) for the full benchmark run output.

- ✍️ write: 4.92 MiB/s (5,037 msg/s)
- 📖️ read: 24.59 MiB/s (25,183 msg/s)
- write latency:
  - p99: 38.66 ms
  - p95: 6.22 ms
  - p50: 5.28 ms
- read latency:
  - p99: 27.13 ms
  - p95: 4.67 ms
  - p50: 3.54 ms
- end-to-end latency:
  - p99: 60.09 ms
  - p95: 10.65 ms
  - p50: 8.76 ms
- bottleneck: number of clients; more clients raises CPU; this number of clients can't push more throughput.

- test parameters:
  - write batch size: 100 records
  - read batch size: 200 records
  - 4 tables (partitions)
  - 10 writers (2 writers per partition on average)
  - 5x read fanout - 5 consumer groups with 4 readers each
    - 20 reader total

- server: `c7i.xlarge`
    - gp3 25GB 9000 IOPS
    - Ubuntu 24.04
    - mostly default postgres settings;
      - only diff is partition tables are configured to vacuum analyze frequently (`SET (autovacuum_analyze_scale_factor = 0.05)`)
    - CPU usage: ~60% CPU
    - us-east-1a
- test driver client: `c7i.xlarge`
    - 10 writers
    - 20 readers
    - 1 KB payload
    - us-east-1a

# Run

```bash
./pg_msg_bench --host=$HOST --port=5432 --db=benchmark --user=postgres --password=postgres --report=5s \
  --write-batch 100 --read-batch 200 --writers=10 \
  --partitions 4 --mode=pubsub \
  --readers=4 --consumer-groups 5 \
  --duration=120s --throttle_writes=5000
```

Notes:
- Similar to the other tests, CPU creeps up on becoming a bottleneck because it can't handle a lot more writer/reader clients;
- In this test, I'm not sure more writer clients would have necessarily increased throughput. More partitions may have helped.
- More aggressive batching would have certainly boosted throughput too, but I kept within the defined 100/200msg batch sizes.
- In general, `iostat` again showed large write amplification on the disk. A ~5 MB/s write workload is writing ~20 MB/s to the disk. I can't explain it.
- 3000k IOPS on gp3 was low, so I increased it to 9000. `iostat` still showed high util percentage, but didn't seem to be anywhere close to bottleneck.

# 💸 Cost

The total cost of this setup is around **\$2418** per year.

- the server is around **\$1626** per year:
  - c7i.xlarge for \$1033.68/yr (at \$0.118/h 1yr reserved pricing)
- the EBS disk is around **\$792** per year
  - Assuming a ~12h message retention (then to s3 or whatever), we'd deploy a 450 GB disk (5 MB/s for 12h is 216GB, so good to deploy with some free space buffer)
    - a 450GB disk is \$432/yr (\$36/m at \$0.08/GB-month)
    - the extra 6000 IOPS is \$360/yr (\$30m at \$0.005/provisioned IOPS-month) 
