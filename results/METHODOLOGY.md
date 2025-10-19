# The Results

I ran these tests in AWS us-east-2 under my personal account. I used spot instances. The full setup script is in [SETUP.md](./SETUP.md). If any test deviates from the test setup, it documents it in its respective folder. 

# General Test Methodology

Benchmarks are always imperfect (noisy neighbours, non-practical workload, etc.). Treat them as a single data point. Re-run them yourself to confirm, and vary the workload to match your perception of a realistic test.

This is what I followed when I ran these:

1. Aim for steady usage of ~50-60%. Ensure a run _never_ hits >90% CPU even during short-term spikes. Better yet - try to prevent large spikes.
2. Run for 2-15 minutes, depending on the workload. Do a few runs (3) and average the results.
3. Run with a small, practical and cheap EC2 instance (~4 vCPU) with all-default Postgres settings.
   - BONUS: Try a run against a multi-node cluster with synchronous replication enabled.
4. Run with a significantly larger machine (e.g. 96 vCPU) with extensive tuning to get a sense of the upper limits.

## Queues

- Always fetch single records per statement, as a queue most likely realistically would.
  - This is inefficient as it results in a ton of overhead per statement, and increases contention per message.
- At least once semantics

See the detailed SQL in the 👉 [README](./queue/README.md)

## Pub Sub

- Have high read fan-out (5x) as a pub-sub system would.
- Use multiple partitions (logs)
- Strict ordering, at least once semantics

See the detailed SQL in the 👉 [README](./pubsub/README.md)

## Replication

The 4vCPU replicated tests consist of 2 standby replicas, with one being synchronous and another being asynchronous.
This means that client writes block until the one `sync` replica successfully persists the change. The other one replicates asynchronously.
