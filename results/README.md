# PostgreSQL Queue and Pub/Sub Benchmark Results

This directory contains comprehensive benchmark results for PostgreSQL used as a **queue** and **pub/sub** messaging system.

These benchmarks were conducted on AWS and are meant to accompany the article ["Kafka is Fast -- I'll use Postgres"](https://topicpartition.io/blog/postgres-pubsub-queue-benchmarks).

## Additional Information

- **[Methodology](./METHODOLOGY.md)** - Detailed explanation of the test methodology for both queue and pub/sub benchmarks
- **[Setup Instructions](./SETUP.md)** - How to set up the benchmark environment on AWS
- **[Queue Design Details](./queue/README.md)** - Implementation details and SQL queries for queue tests
- **[Pub/Sub Design Details](./pubsub/README.md)** - Implementation details and SQL queries for pub/sub tests

## Pub/Sub Results

The [Pub/Sub Test Methodology](./METHODOLOGY.md#pub-sub) ensures strict ordering per log, uses multiple logs (partitions), has a high read fan-out (5x, i.e. every byte written is read 5 times) and has at-least-once semantics.

| Configuration                                                                                    | Server | Write Rate | Read Rate (5x fanout) | E2E Latency (ms) | Write Latency (ms) | Read Latency (ms) | Partitions | PG Settings                                                                                                     | Bottleneck                                          |
|--------------------------------------------------------------------------------------------------|---------|------------|-----------------------|------------------|-------------------|------------------|------------|-----------------------------------------------------------------------------------------------------------------|-----------------------------------------------------|
| 🔥 **[4vCPU three-node sync replicated](./results/pubsub/4vcpu/three_node/4vcpu_replicated.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB) | 5,015 msg/s<br>4.90 MiB/s | 25,073 msg/s<br>24.49 MiB/s | p99: 185.37<br>p95: 12.35<br>p50: 8.89 | p99: 153.45<br>p95: 6.77<br>p50: 5.01 | p99: 56.69<br>p95: 4.91<br>p50: 3.89 | 4 | `FIRST 1 (pg2, pg3)` and `SET (autovacuum_analyze_scale_factor = 0.05)` on the partition tables                 | number of clients                                   |
| 🔥 **[4vCPU single-node](./results/pubsub/4vcpu/single_node/4vcpu.md)**                          | c7i.xlarge<br>(4 vCPU, 8 GiB) | 5,037 msg/s<br>4.92 MiB/s | 25,183 msg/s<br>24.59 MiB/s | p99: 60.09<br>p95: 10.65<br>p50: 8.76 | p99: 38.66<br>p95: 6.22<br>p50: 5.28 | p99: 27.13<br>p95: 4.67<br>p50: 3.54 | 4 | All defaults<br>(besides `SET (autovacuum_analyze_scale_factor = 0.05)` on the partition tables)                | number of clients                                   |
| 🔥 **[96vCPU single-node](./results/pubsub/96vcpu/single_node/96vcpu.md)**                       | c7i.24xlarge<br>(96 vCPU, 192 GiB) | 243k msg/s<br>238 MiB/s | 1.2M msg/s<br>1.16 GiB/s | p99: 853.37<br>p95: 242.13<br>p50: 23.43 | p99: 137.78<br>p95: 47.15<br>p50: 10.61 | p99: 24.61<br>p95: 19.50<br>p50: 10.45 | 30 | Modified (upscaled, huge_pages) <br> and `SET (autovacuum_analyze_scale_factor = 0.05)` on the partition tables | unsure, probably partitions (8 MiB/s per partition) |

## Queue Results

The [Queue Test Methodology](./METHODOLOGY.md#queues) uses a single record append/pop per statement and at-least-once semantics.

⚠️ The e2e p99 latency for the queue tests suffers from a flaw that inflates the number. See the [High Latency Note](./queue/IMPERFECTIONS.md#high-e2e-latency-note) on that.
The real number is likely significantly lower.

| Configuration | Server | PG Settings | Read/Write Rate | E2E Latency (ms)                                                                                                   | Write Latency (ms) | Read Latency (ms) | Bottleneck                                                                                                                                                                                    |
|--------------|---------|------------|-----------------|--------------------------------------------------------------------------------------------------------------------|-------------------|------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **[4vCPU three-node sync replicated](./results/queue/4vcpu/three_node/4vcpu_replicated.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB) | `FIRST 1 (pg2, pg3)` | 2,397 msg/s<br>2.34 MiB/s | p99:️ 920 ⚠️ [*(see flaw)*](./results/queue/IMPERFECTIONS.md#high-e2e-latency-note)<br>p95: 536<br>p50: 7          | p99: 3.3<br>p95: 2.6<br>p50: 2 | p99: 7.65<br>p95: 6.95<br>p50: 5.8 | number of clients                                                                                                                                                                             |
| 🔥 **[4vCPU single-node](./results/queue/4vcpu/single_node/4vcpu.md)** | c7i.xlarge<br>(4 vCPU, 8 GiB) | Default | 2,885 msg/s<br>2.82 MiB/s | p99: 15.86<br>p95: 5.74<br>p50: 5.07                                                                               | p99: 2.46<br>p95: 2.27<br>p50: 1.77 | p99: 4.20<br>p95: 3.64<br>p50: 3.37 | number of clients                                                                                                                                                                             |
| **[96vCPU single-node](./results/queue/96vcpu/single_node/96vcpu.md)** | c7i.24xlarge<br>(96 vCPU, 192 GiB) | Modified<br>(upscaled, huge_pages) | 20,144 msg/s<br>19.67 MiB/s | p99: 930.09 ⚠️ [*(see flaw)*](./results/queue/IMPERFECTIONS.md#high-e2e-latency-note)<br>p95: 709.01<br>p50: 12.59 | p99: 9.42<br>p95: 1.99<br>p50: 1.16 | p99: 22.59<br>p95: 14.93<br>p50: 6.27 | Not pushed further; Some bottleneck was probably hit due to single-table congestion; It's very likely multiple different queue tables would have scaled way beyond this number. | 
