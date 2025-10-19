# High E2E Latency Note

Some queue tests showed higher E2E latencies which I believe aren't an accurate representation.

In the pub-sub tests, I made readers start _before_ the writers and added [a small 1s sleep before starting the writers](../../pubsub/pubsub_runner.go):
```goland
time.Sleep(1000 * time.Millisecond)
	log.Printf("[pub info] spawned readers")
```
For the queue tests, though, I didn't do this. 😔

The result is that queue tests immediately spike queue depth at startup, because the writers manage to get a head start before the readers.
I believe the E2E latency is artificially high because of this flaw in the test. If you notice the detailed results in e.g. [the 96 vCPU RUN.md](96vcpu/single_node/run.md), you'll see the queue depth spikes only at **the beginning**.

In the first 15 seconds, it spikes to being 16,622 messages behind. After, it never goes consistently above 430 despite one small spike to 1366. 👇
At 20,000 msg/s, that's almost a full second of delay that the readers have to catch up **on top of** all new 20k messages coming per second.
This artificially increases the E2E latency to a number that's higher than what it should be.
```bash
ubuntu@ip-172-31-46-123:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres   --writers=100   --readers=200   --duration=120s   --payload=1024   --report=5s  --throttle_writes=20000 --mode=queue --tune-table-vacuum
# first 15 seconds:
[18:40:27] W: 23982/s R: 20657/s QDepth: 16622 # <--- HIGH
[18:40:32] W: 19978/s R: 20796/s QDepth: 12533 # <--- HIGH
[18:40:37] W: 20031/s R: 22072/s QDepth: 2329 # <--- high-ish
# the rest 105 seconds:
[18:40:42] W: 20014/s R: 20431/s QDepth: 242 Err(W/R): 0/0
[18:40:47] W: 20066/s R: 20036/s QDepth: 392 Err(W/R): 0/0
[18:40:52] W: 20020/s R: 20057/s QDepth: 207 Err(W/R): 0/0
[18:40:57] W: 20014/s R: 20036/s QDepth: 97 Err(W/R): 0/0
[18:41:02] W: 20020/s R: 20003/s QDepth: 179 Err(W/R): 0/0
[18:41:07] W: 20023/s R: 20042/s QDepth: 83 Err(W/R): 0/0
[18:41:12] W: 20025/s R: 19768/s QDepth: 1366 Err(W/R): 0/0
[18:41:17] W: 20016/s R: 20271/s QDepth: 89 Err(W/R): 0/0
[18:41:22] W: 20018/s R: 20018/s QDepth: 91 Err(W/R): 0/0
[18:41:27] W: 20020/s R: 20016/s QDepth: 111 Err(W/R): 0/0
[18:41:32] W: 20015/s R: 20003/s QDepth: 175 Err(W/R): 0/0
[18:41:37] W: 20021/s R: 20035/s QDepth: 106 Err(W/R): 0/0
[18:41:42] W: 20020/s R: 20025/s QDepth: 85 Err(W/R): 0/0
[18:41:47] W: 20017/s R: 20017/s QDepth: 87 Err(W/R): 0/0
[18:41:52] W: 20016/s R: 20014/s QDepth: 97 Err(W/R): 0/0
[18:41:57] W: 20018/s R: 20020/s QDepth: 89 Err(W/R): 0/0
[18:42:02] W: 20019/s R: 19951/s QDepth: 425 Err(W/R): 0/0
[18:42:07] W: 20014/s R: 20043/s QDepth: 284 Err(W/R): 0/0
[18:42:12] W: 20019/s R: 20013/s QDepth: 313 Err(W/R): 0/0
[18:42:17] W: 20021/s R: 19994/s QDepth: 447 Err(W/R): 0/0
```
