# Run 1

```bash
ubuntu@ip-172-31-89-39:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=10   --readers=15   --duration=600s   --payload=1024   --report=5s   --throttle_writes=2400   --mode=queue
[20:39:22] W: 2873/s R: 2594/s QDepth: 1396 Err(W/R): 0/0
[20:39:27] W: 2401/s R: 2568/s QDepth: 564 Err(W/R): 0/0
[20:39:32] W: 2400/s R: 2510/s QDepth: 13 Err(W/R): 0/0
[20:39:37] W: 2400/s R: 2401/s QDepth: 10 Err(W/R): 0/0
[20:39:42] W: 2400/s R: 2399/s QDepth: 14 Err(W/R): 0/0
[20:39:47] W: 2400/s R: 2396/s QDepth: 35 Err(W/R): 0/0
[20:39:52] W: 2400/s R: 2404/s QDepth: 14 Err(W/R): 0/0
[20:39:57] W: 2400/s R: 2401/s QDepth: 12 Err(W/R): 0/0
[20:40:02] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:40:07] W: 2400/s R: 2399/s QDepth: 12 Err(W/R): 0/0
[20:40:12] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:40:17] W: 2400/s R: 2396/s QDepth: 31 Err(W/R): 0/0
[20:40:22] W: 2400/s R: 2404/s QDepth: 12 Err(W/R): 0/0
[20:40:27] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:40:32] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:40:37] W: 2400/s R: 2386/s QDepth: 81 Err(W/R): 0/0
[20:40:42] W: 2400/s R: 2414/s QDepth: 15 Err(W/R): 0/0
[20:40:47] W: 2400/s R: 2401/s QDepth: 12 Err(W/R): 0/0
[20:40:52] W: 2400/s R: 2398/s QDepth: 21 Err(W/R): 0/0
[20:40:57] W: 2400/s R: 2397/s QDepth: 39 Err(W/R): 0/0
[20:41:02] W: 2400/s R: 2405/s QDepth: 16 Err(W/R): 0/0
[20:41:07] W: 2399/s R: 2401/s QDepth: 10 Err(W/R): 0/0
[20:41:12] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:41:17] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:41:22] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:41:27] W: 2400/s R: 2396/s QDepth: 31 Err(W/R): 0/0
[20:41:32] W: 2400/s R: 2404/s QDepth: 12 Err(W/R): 0/0
[20:41:37] W: 2400/s R: 2393/s QDepth: 48 Err(W/R): 0/0
[20:41:42] W: 2400/s R: 2408/s QDepth: 11 Err(W/R): 0/0
[20:41:47] W: 2400/s R: 2399/s QDepth: 15 Err(W/R): 0/0
[20:41:52] W: 2400/s R: 2401/s QDepth: 10 Err(W/R): 0/0
[20:41:57] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:42:02] W: 2400/s R: 2395/s QDepth: 37 Err(W/R): 0/0
[20:42:07] W: 2400/s R: 2405/s QDepth: 11 Err(W/R): 0/0
[20:42:12] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:42:17] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:42:22] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:42:27] W: 2400/s R: 2400/s QDepth: 16 Err(W/R): 0/0
[20:42:32] W: 2400/s R: 2400/s QDepth: 15 Err(W/R): 0/0
[20:42:37] W: 2400/s R: 2400/s QDepth: 16 Err(W/R): 0/0
[20:42:42] W: 2400/s R: 2302/s QDepth: 509 Err(W/R): 0/0
[20:42:47] W: 2400/s R: 2499/s QDepth: 13 Err(W/R): 0/0
[20:42:52] W: 2400/s R: 2400/s QDepth: 16 Err(W/R): 0/0
[20:42:57] W: 2400/s R: 2400/s QDepth: 16 Err(W/R): 0/0
[20:43:02] W: 2399/s R: 2401/s QDepth: 9 Err(W/R): 0/0
[20:43:07] W: 2400/s R: 2399/s QDepth: 14 Err(W/R): 0/0
[20:43:12] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:43:17] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:43:22] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:43:27] W: 2399/s R: 2399/s QDepth: 13 Err(W/R): 0/0
[20:43:32] W: 2401/s R: 2401/s QDepth: 12 Err(W/R): 0/0
[20:43:37] W: 2400/s R: 2398/s QDepth: 22 Err(W/R): 0/0
[20:43:42] W: 2400/s R: 2245/s QDepth: 796 Err(W/R): 0/0
[20:43:47] W: 2400/s R: 2557/s QDepth: 11 Err(W/R): 0/0
[20:43:52] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:43:57] W: 2400/s R: 2399/s QDepth: 13 Err(W/R): 0/0
[20:44:02] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:44:07] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:44:12] W: 2400/s R: 2400/s QDepth: 14 Err(W/R): 0/0
[20:44:17] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:44:22] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:44:27] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:44:32] W: 2400/s R: 2401/s QDepth: 8 Err(W/R): 0/0
[20:44:37] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:44:42] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:44:47] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:44:52] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:44:57] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:45:02] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:45:07] W: 2400/s R: 2398/s QDepth: 24 Err(W/R): 0/0
[20:45:12] W: 2400/s R: 2402/s QDepth: 13 Err(W/R): 0/0
[20:45:17] W: 2400/s R: 2401/s QDepth: 9 Err(W/R): 0/0
[20:45:22] W: 2400/s R: 2399/s QDepth: 13 Err(W/R): 0/0
[20:45:27] W: 2400/s R: 2401/s QDepth: 10 Err(W/R): 0/0
[20:45:32] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:45:37] W: 2400/s R: 2384/s QDepth: 89 Err(W/R): 0/0
[20:45:42] W: 2400/s R: 2416/s QDepth: 10 Err(W/R): 0/0
[20:45:47] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:45:52] W: 2400/s R: 2400/s QDepth: 9 Err(W/R): 0/0
[20:45:57] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:46:02] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:46:07] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:46:12] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:46:17] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:46:22] W: 2400/s R: 2399/s QDepth: 13 Err(W/R): 0/0
[20:46:27] W: 2400/s R: 2401/s QDepth: 11 Err(W/R): 0/0
[20:46:32] W: 2400/s R: 2399/s QDepth: 14 Err(W/R): 0/0
[20:46:37] W: 2400/s R: 2391/s QDepth: 55 Err(W/R): 0/0
[20:46:42] W: 2400/s R: 2364/s QDepth: 238 Err(W/R): 0/0
[20:46:47] W: 2400/s R: 2386/s QDepth: 308 Err(W/R): 0/0
[20:46:52] W: 2335/s R: 1935/s QDepth: 2309 Err(W/R): 0/0
[20:46:57] W: 2394/s R: 2516/s QDepth: 1699 Err(W/R): 0/0
[20:47:02] W: 2406/s R: 2507/s QDepth: 1193 Err(W/R): 0/0
[20:47:07] W: 2400/s R: 2492/s QDepth: 734 Err(W/R): 0/0
[20:47:12] W: 2400/s R: 2486/s QDepth: 305 Err(W/R): 0/0
[20:47:17] W: 2400/s R: 2458/s QDepth: 15 Err(W/R): 0/0
[20:47:22] W: 2400/s R: 2400/s QDepth: 14 Err(W/R): 0/0
[20:47:27] W: 2400/s R: 2397/s QDepth: 29 Err(W/R): 0/0
[20:47:32] W: 2400/s R: 2399/s QDepth: 33 Err(W/R): 0/0
[20:47:37] W: 2400/s R: 2404/s QDepth: 11 Err(W/R): 0/0
[20:47:42] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:47:47] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:47:52] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:47:57] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:48:02] W: 2400/s R: 2401/s QDepth: 11 Err(W/R): 0/0
[20:48:07] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:48:12] W: 2400/s R: 2400/s QDepth: 10 Err(W/R): 0/0
[20:48:17] W: 2400/s R: 2400/s QDepth: 11 Err(W/R): 0/0
[20:48:22] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:48:27] W: 2400/s R: 2400/s QDepth: 13 Err(W/R): 0/0
[20:48:32] W: 2400/s R: 2400/s QDepth: 12 Err(W/R): 0/0
[20:48:37] W: 2400/s R: 2389/s QDepth: 67 Err(W/R): 0/0
[20:48:42] W: 2400/s R: 2353/s QDepth: 300 Err(W/R): 0/0
[20:48:47] W: 1754/s R: 1744/s QDepth: 351 Err(W/R): 0/0
[20:48:52] W: 2361/s R: 1952/s QDepth: 2400 Err(W/R): 0/0
[20:48:57] W: 2400/s R: 2527/s QDepth: 1768 Err(W/R): 0/0
[20:49:02] W: 2400/s R: 2511/s QDepth: 1213 Err(W/R): 0/0
[20:49:07] W: 2400/s R: 2483/s QDepth: 800 Err(W/R): 0/0
[20:49:12] W: 2400/s R: 2480/s QDepth: 401 Err(W/R): 0/0

=== Summary ===
Total Writes: 1438586
Total Reads: 1438527
Total Updates: 1438527
Write Errors: 0
Read Errors: 4
Avg Write Throughput: 2397.64 rows/sec
Avg Read Throughput: 2397.55 rows/sec

Write Latencies (INSERT only):
  P50: 2.029567ms
  P95: 2.629631ms
  P99: 3.307519ms

Read Latencies (txn: SELECT+DELETE+INSERT):
  P50: 5.808127ms
  P95: 6.959103ms
  P99: 7.655423ms

End-to-End Latencies (created_at → consumed):
  P50: 7.184383ms
  P95: 536.870911ms
  P99: 920.125439ms

2025/10/05 20:49:17 queue benchmark complete
```
