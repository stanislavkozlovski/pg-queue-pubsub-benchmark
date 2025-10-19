# 🐘 postgres-queue-benchmarks

This repository contains code for benchmarking PostgreSQL as a **queue** and a Kafka-like **pub/sub** messaging system.

It also contains results from a few runs.

It is meant to accompany my article ["Kafka is Fast -- I'll use Postgres"](https://topicpartition.io/blog/postgres-pubsub-queue-benchmarks).

# AWS Benchmark Results

I spent ~$50 and a week+ to run benchmarks on AWS. The extensive results are in the [./results](./results) folder. Here are the highlights:


## Pub/Sub Results

The [Pub/Sub Test Methodology](./results/METHODOLOGY.md#pub-sub) ensures strict ordering per log, uses multiple logs (partitions), has a high read fan-out (5x, i.e. every byte written is read 5 times) and has at-least-once semantics.

| Configuration | Server | Write Rate | Read Rate (5× fanout) | E2E Latency (ms) |
|----------------|---------|-------------|------------------------|------------------|
| 🔥 **[4vCPU three-node sync replicated](./results/pubsub/4vcpu/three_node/4vcpu_replicated.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB)<br>4 partitions | 5,015 msg/s <br>4.90 MiB/s | 25,073 msg/s <br>24.49 MiB/s | p99: 185.37 <br>p95: 12.35 <br>p50: 8.89 |
| 🔥 **[4vCPU single-node](./results/pubsub/4vcpu/single_node/4vcpu.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB)<br>4 partitions | 5,037 msg/s <br>4.92 MiB/s | 25,183 msg/s <br>24.59 MiB/s | p99: 60.09 <br>p95: 10.65 <br>p50: 8.76 |
| 🔥 **[96vCPU single-node](./results/pubsub/96vcpu/single_node/96vcpu.md)** | c7i.24xlarge<br>(96 vCPU, 192 GiB)<br>30 partitions | 243 k msg/s <br>238 MiB/s | 1.2 M msg/s <br>1.16 GiB/s | p99: 853.37 <br>p95: 242.13 <br>p50: 23.43 |

See detailed results here 👉 [results/README.md](./results/README.md#pubsub-results)

## Queue Results

The [Queue Test Methodology](./results/METHODOLOGY.md#queues) uses a single record append/pop per statement and at-least-once semantics.

⚠️ The e2e p99 latency for the queue tests suffers from a flaw that inflates the number. See the [High Latency Note](./results/queue/IMPERFECTIONS.md#high-e2e-latency-note) on that.
The real number is likely significantly lower.

| Configuration | Server | Read/Write Rate | E2E Latency (ms)                                                                                            |
|----------------|---------|------------------|-------------------------------------------------------------------------------------------------------------|
| **[4vCPU three-node sync replicated](./results/queue/4vcpu/three_node/4vcpu_replicated.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB) | 2,397 msg/s <br>2.34 MiB/s | p99: 920 ⚠️ [*see flaw*](./results/queue/IMPERFECTIONS.md#high-e2e-latency-note) <br>p95: 536 <br>p50: 7    |
| 🔥 **[4vCPU single-node](./results/queue/4vcpu/single_node/4vcpu.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB) | 2,885 msg/s <br>2.82 MiB/s | p99: 15.86 <br>p95: 5.74 <br>p50: 5.07                                                                      |
| **[96vCPU single-node](./results/queue/96vcpu/single_node/96vcpu.md)** | c7i.24xlarge<br>(96 vCPU, 192 GiB) | 20,144 msg/s <br>19.67 MiB/s | p99: 930 ⚠️ [*see flaw*](./results/queue/IMPERFECTIONS.md#high-e2e-latency-note) <br>p95: 709 <br>p50: 12.6 |

See detailed results here 👉 [results/README.md](./results/README.md#queue-results)

## Results Conclusion

> **Disclaimer:** I am not a Postgres expert. I have experience with distributed infra software (Kafka) and try to think from first principles. (I also use a lot of ChatGPT to get ideas and confirm behavior).
>
> It is very likely there are some mistakes in these benchmarks - please feel encouraged to point them out to me. I'm interested in getting to the truth rather than proving a point. (make sure you read my article to understand my full stance on how PG competes with other systems though)



----

## Build
```
go build -o pg_msg_bench .
```

## Usage

```
./pg_msg_bench [flags]
```

### Flags

#### Connection Flags
- `-host` - PostgreSQL host (default: `localhost`)
- `-port` - PostgreSQL port (default: `5432`)
- `-db` - Database name (default: `benchmark`)
- `-user` - Database user (default: `postgres`)
- `-password` - Database password (default: `""`)

#### General Benchmark Flags
- `-mode` - Benchmark mode: `queue` or `pubsub` (default: `queue`)
- `-writers` - Number of writer goroutines (default: `4`)
- `-readers` - Number of reader goroutines (default: `4`)
- `-duration` - Test duration (default: `30s`)
- `-payload` - Payload size in bytes (default: `1024`)
- `-report` - Report interval (default: `5s`)
- `-throttle_writes` - Throttle writer rows/sec, 0 = unlimited (default: `0`)

#### Queue Mode Flags
- `-tune-table-vacuum` - Apply aggressive autovacuum/fillfactor settings to the queue table (default: `false`)

#### Pub/Sub Mode Flags
- `-consumer-groups` - Number of consumer groups for pub-sub. Each group has `-readers` goroutines consuming for it. (default: `1`)
- `-partitions` - Number of logs/partitions for pub-sub. A partition is associated with only one reader goroutine per consumer group. (default: `10`)
- `-read-batch` - Max number of messages to read per request for pub-sub (default: `100`)
- `-write-batch` - Number of messages per write request for pub-sub (default: `100`)
