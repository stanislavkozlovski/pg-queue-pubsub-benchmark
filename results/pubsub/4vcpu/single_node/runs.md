
# Run 1
```bash
buntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s    --write-batch 100   --read-batch 200     --writers=10   --readers=4   --consumer-groups 5     --mode=pubsub --partitions 4 --duration=120s --throttle_writes=5000
2025/10/03 12:30:15 [pub info] successfully created 4 topicpartitions w/ counters + consumer_offsets
2025/10/03 12:30:15 [pub info] producers per partition [p1=3 p2=3 p3=2 p4=2]
[12:30:20] W: 5980/s TotalW:29900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5960     29800      0        5145     104.9    100     
  1      5960     29800      0        5040     106.0    100     
  2      5960     29800      0        5070     104.9    100     
  3      5960     29800      0        5123     104.9    100     
  4      5960     29800      0        5090     106.0    100     
  TotalR: 29800/s QueueDepth(min=100, max=100)

[12:30:25] W: 5000/s TotalW:54900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     54800      0        10447    102.6    100     
  1      5000     54800      0        10249    103.2    100     
  2      5000     54800      0        10291    102.6    100     
  3      5000     54800      0        10379    102.6    100     
  4      5000     54800      0        10321    103.2    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:30] W: 5000/s TotalW:79900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     79800      0        14715    101.8    100     
  1      5000     79800      0        14465    102.2    100     
  2      5000     79800      0        14483    101.8    100     
  3      5000     79800      0        14636    101.8    100     
  4      5000     79800      0        14551    102.2    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:35] W: 5000/s TotalW:104900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     104800     0        19453    101.4    100     
  1      5000     104800     0        19109    101.6    100     
  2      5000     104800     0        19114    101.4    100     
  3      5000     104800     0        19342    101.4    100     
  4      5000     104800     0        19230    101.6    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:40] W: 5000/s TotalW:129900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     129800     0        24882    101.1    100     
  1      5000     129800     0        24431    101.3    100     
  2      5000     129800     0        24443    101.1    100     
  3      5000     129800     0        24737    101.1    100     
  4      5000     129800     0        24596    101.3    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:45] W: 5000/s TotalW:154900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     154800     0        29516    100.9    100     
  1      5000     154800     0        28953    101.1    100     
  2      5000     154800     0        28967    100.9    100     
  3      5000     154800     0        29314    100.9    100     
  4      5000     154800     0        29179    101.1    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:50] W: 5000/s TotalW:179900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     179800     0        33672    100.8    100     
  1      5000     179800     0        33039    101.0    100     
  2      5000     179800     0        33049    100.8    100     
  3      5000     179800     0        33446    100.8    100     
  4      5000     179800     0        33304    101.0    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:30:55] W: 5000/s TotalW:204900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     204800     0        39027    100.7    100     
  1      5000     204800     0        38293    100.8    100     
  2      5000     204800     0        38343    100.7    100     
  3      5000     204800     0        38762    100.7    100     
  4      5000     204800     0        38626    100.8    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:00] W: 5000/s TotalW:229900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     229800     0        44151    100.6    100     
  1      5000     229800     0        43348    100.7    100     
  2      5000     229800     0        43413    100.6    100     
  3      5000     229800     0        43875    100.6    100     
  4      5000     229800     0        43721    100.7    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:05] W: 5000/s TotalW:254900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     254800     0        48391    100.6    100     
  1      5000     254800     0        47516    100.7    100     
  2      5000     254800     0        47585    100.6    100     
  3      5000     254800     0        48105    100.6    100     
  4      5000     254800     0        47929    100.7    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:10] W: 5000/s TotalW:279900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     279800     0        53269    100.5    100     
  1      5000     279800     0        52299    100.6    100     
  2      5000     279800     0        52386    100.5    100     
  3      5000     279800     0        52930    100.5    100     
  4      5000     279800     0        52745    100.6    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:15] W: 5000/s TotalW:304900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     304800     0        58575    100.5    100     
  1      5000     304800     0        57499    100.6    100     
  2      5000     304800     0        57606    100.5    100     
  3      5000     304800     0        58178    100.5    100     
  4      5000     304800     0        57988    100.6    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:20] W: 5000/s TotalW:329900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     329800     0        61540    100.6    100     
  1      5000     329800     0        60408    100.7    100     
  2      5000     329800     0        60530    100.7    100     
  3      5000     329800     0        61108    100.6    100     
  4      5000     329800     0        60930    100.7    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:25] W: 5000/s TotalW:354900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     354800     0        65806    100.5    100     
  1      5000     354800     0        64632    100.6    100     
  2      5000     354800     0        64741    100.7    100     
  3      5000     354800     0        65341    100.5    100     
  4      5000     354800     0        65158    100.7    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:30] W: 5000/s TotalW:379900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     379800     0        71096    100.5    100     
  1      5000     379800     0        69810    100.6    100     
  2      5000     379800     0        69949    100.6    100     
  3      5000     379800     0        70575    100.5    100     
  4      5000     379800     0        70389    100.6    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:35] W: 5000/s TotalW:404900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     404800     0        75995    100.5    100     
  1      5000     404800     0        74606    100.5    100     
  2      5000     404800     0        74766    100.6    100     
  3      5000     404800     0        75433    100.5    100     
  4      5000     404800     0        75225    100.6    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:40] W: 5000/s TotalW:429900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     429800     0        80205    100.4    100     
  1      5000     429800     0        78716    100.5    100     
  2      5000     429800     0        78900    100.6    100     
  3      5000     429800     0        79612    100.4    100     
  4      5000     429800     0        79393    100.5    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:45] W: 5000/s TotalW:454900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     454800     0        85110    100.4    100     
  1      5000     454800     0        83545    100.5    100     
  2      5000     454800     0        83724    100.5    100     
  3      5000     454800     0        84461    100.4    100     
  4      5000     454800     0        84238    100.5    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:31:50] W: 5000/s TotalW:479900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     479800     0        90437    100.4    100     
  1      5000     479800     0        88745    100.5    100     
  2      5000     479800     0        88955    100.5    100     
  3      5000     479800     0        89741    100.4    100     
  4      5000     479800     0        89493    100.5    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:31:55] W: 5000/s TotalW:504900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     504800     0        94687    100.4    100     
  1      5000     504800     0        92892    100.4    100     
  2      5000     504800     0        93106    100.5    100     
  3      5000     504800     0        93932    100.4    100     
  4      5000     504800     0        93669    100.5    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

[12:32:00] W: 5000/s TotalW:529900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     529800     0        99155    100.4    100     
  1      5000     529800     0        97271    100.4    100     
  2      5000     529800     0        97501    100.5    100     
  3      5000     529800     0        98344    100.4    100     
  4      5000     529800     0        98061    100.4    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:32:05] W: 5000/s TotalW:554900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     554800     0        104454   100.3    100     
  1      5000     554800     0        102477   100.4    100     
  2      5000     554800     0        102710   100.4    100     
  3      5000     554800     0        103589   100.3    100     
  4      5000     554800     0        103294   100.4    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

[12:32:10] W: 5000/s TotalW:579900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     579800     0        109223   100.3    100     
  1      5000     579800     0        107155   100.4    100     
  2      5000     579800     0        107412   100.4    100     
  3      5000     579800     0        108306   100.3    100     
  4      5000     579800     0        107998   100.4    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

2025/10/03 12:32:15 [consumer g2 r2 p3] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 12:32:15 [consumer g4 r1 p2] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 12:32:15 [consumer g0 r1 p2] claim err: context deadline exceeded (group=g0 batch=200)

=== Summary ===
Total Writes: 604400
Total Reads: 3022000
Total Updates: 0
Write Errors: 0
Read Errors: 0
Avg Write Throughput: 5036.67 rows/sec
Avg Read Throughput: 25183.33 rows/sec

Write Latencies (INSERT only):
  P50: 5.267455ms
  P95: 6.184959ms
  P99: 43.450367ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 0s
  P95: 0s
  P99: 0s

End-to-End Latencies (created_at → consumed):
  P50: 0s
  P95: 0s
  P99: 0s


=== PubSub Summary ===
Delivery Semantics: at-least-once (strict ordering)
Number of partitions: 4
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      113381
  Avg Poll Size:     100.32
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.493887ms
    P95: 4.599807ms
    P99: 9.019391ms

  End-to-End Latency (created_at → consumed):
    P50: 8.667135ms
    P95: 10.534911ms
    P99: 86.441983ms

-- Consumer Group 1 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      0
  Empty Claims:      111241
  Avg Poll Size:     100.37
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.567615ms
    P95: 4.653055ms
    P99: 9.355263ms

  End-to-End Latency (created_at → consumed):
    P50: 8.781823ms
    P95: 10.649599ms
    P99: 86.966271ms

-- Consumer Group 2 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      111502
  Avg Poll Size:     100.40
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.551231ms
    P95: 4.612095ms
    P99: 8.527871ms

  End-to-End Latency (created_at → consumed):
    P50: 8.740863ms
    P95: 10.600447ms
    P99: 83.165183ms

-- Consumer Group 3 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      0
  Empty Claims:      112417
  Avg Poll Size:     100.32
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.532799ms
    P95: 4.607999ms
    P99: 8.855551ms

  End-to-End Latency (created_at → consumed):
    P50: 8.757247ms
    P95: 10.616831ms
    P99: 85.590015ms

-- Consumer Group 4 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      112107
  Avg Poll Size:     100.38
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.516415ms
    P95: 4.595711ms
    P99: 8.814591ms

  End-to-End Latency (created_at → consumed):
    P50: 8.716287ms
    P95: 10.543103ms
    P99: 83.361791ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 3022000
  Total Read Errors:     0
  Total Updates:         3022000
  Total Claim Errors:    3
  Total Empty Claims:    560648
  Avg Poll Size:         100.36
  Avg Throughput:        25183.33 reads/sec

  Read Latency (claim txn):
    P50: 3.532799ms
    P95: 4.616191ms
    P99: 8.888319ms

  End-to-End Latency (created_at → consumed):
    P50: 8.732671ms
    P95: 10.584063ms
    P99: 85.917695ms

2025/10/03 12:32:15 pubsub benchmark complete
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ 
```

# Run 2
```bash
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s    --write-batch 100   --read-batch 200     --writers=10   --readers=4   --consumer-groups 5     --mode=pubsub --partitions 4 --duration=120s --throttle_writes=5000
2025/10/03 12:33:25 [pub info] successfully created 4 topicpartitions w/ counters + consumer_offsets
2025/10/03 12:33:25 [pub info] producers per partition [p1=3 p2=3 p3=2 p4=2]
[12:33:30] W: 5980/s TotalW:29900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5980     29900      0        5202     102.7    0       
  1      5960     29800      0        5192     101.4    100     
  2      5960     29800      0        5272     101.7    100     
  3      5980     29900      0        5232     101.0    0       
  4      5960     29800      0        5100     102.4    100     
  TotalR: 29839/s QueueDepth(min=0, max=100)

[12:33:35] W: 5000/s TotalW:54900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     54800      0        10050    101.5    100     
  1      5020     54900      0        10044    100.7    0       
  2      5000     54800      0        10179    100.9    100     
  3      4980     54800      0        10100    100.6    100     
  4      5020     54900      0        9871     101.3    0       
  TotalR: 25001/s QueueDepth(min=0, max=100)

[12:33:40] W: 5000/s TotalW:79900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     79800      0        14311    101.0    100     
  1      4980     79800      0        14290    100.5    100     
  2      5020     79900      0        14475    100.6    0       
  3      5020     79900      0        14355    100.4    0       
  4      4980     79800      0        14057    100.9    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:33:45] W: 5000/s TotalW:104900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     104800     0        19368    100.8    100     
  1      5000     104800     0        19292    100.4    100     
  2      5000     104900     0        19543    100.5    0       
  3      4980     104800     0        19414    100.3    100     
  4      5000     104800     0        19006    100.7    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[12:33:50] W: 5000/s TotalW:129900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     129800     0        24718    100.6    100     
  1      5000     129800     0        24634    100.3    100     
  2      4980     129800     0        24930    100.4    100     
  3      5020     129900     0        24754    100.2    0       
  4      5000     129800     0        24249    100.5    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:33:55] W: 5000/s TotalW:154900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     154800     0        28958    100.5    100     
  1      5000     154800     0        28872    100.3    100     
  2      5000     154800     0        29219    100.3    100     
  3      5000     154900     0        29002    100.2    0       
  4      5000     154800     0        28406    100.5    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:34:00] W: 5000/s TotalW:179900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     179900     0        33500    100.4    0       
  1      5020     179900     0        33400    100.2    0       
  2      5020     179900     0        33813    100.3    0       
  3      5000     179900     0        33545    100.2    0       
  4      5020     179900     0        32870    100.4    0       
  TotalR: 25080/s QueueDepth(min=0, max=0)

[12:34:05] W: 5000/s TotalW:204900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     204800     0        38796    100.4    100     
  1      4980     204800     0        38715    100.2    100     
  2      4980     204800     0        39195    100.2    100     
  3      4980     204800     0        38876    100.1    100     
  4      4980     204800     0        38123    100.3    100     
  TotalR: 24900/s QueueDepth(min=100, max=100)

[12:34:10] W: 5000/s TotalW:229900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     229900     0        43476    100.3    0       
  1      5000     229800     0        43366    100.2    100     
  2      5020     229900     0        43943    100.2    0       
  3      5000     229800     0        43569    100.1    100     
  4      5000     229800     0        42728    100.3    100     
  TotalR: 25040/s QueueDepth(min=0, max=100)

[12:34:15] W: 5000/s TotalW:254900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     254800     0        47648    100.3    100     
  1      5000     254800     0        47563    100.2    100     
  2      4980     254800     0        48183    100.2    100     
  3      5000     254800     0        47742    100.1    100     
  4      5000     254800     0        46841    100.3    100     
  TotalR: 24960/s QueueDepth(min=100, max=100)

[12:34:20] W: 5000/s TotalW:279900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     279800     0        52713    100.3    100     
  1      5000     279800     0        52615    100.1    100     
  2      5000     279800     0        53322    100.2    100     
  3      5020     279900     0        52825    100.1    0       
  4      5020     279900     0        51821    100.3    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[12:34:25] W: 5000/s TotalW:304900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     304800     0        57827    100.3    100     
  1      5020     304900     0        57729    100.1    0       
  2      5020     304900     0        58518    100.2    0       
  3      4980     304800     0        57967    100.1    100     
  4      5000     304900     0        56843    100.2    0       
  TotalR: 25019/s QueueDepth(min=0, max=100)

[12:34:30] W: 5000/s TotalW:329900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     329900     0        61125    100.2    0       
  1      4980     329800     0        61021    100.1    100     
  2      5000     329900     0        61861    100.2    0       
  3      5020     329900     0        61294    100.1    0       
  4      4980     329800     0        60090    100.2    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:34:35] W: 5000/s TotalW:354900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     354900     0        65739    100.2    0       
  1      5000     354800     0        65633    100.1    100     
  2      4980     354800     0        66535    100.1    100     
  3      4980     354800     0        65933    100.1    100     
  4      5000     354800     0        64632    100.2    100     
  TotalR: 24959/s QueueDepth(min=0, max=100)

[12:34:40] W: 5000/s TotalW:379900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     379800     0        71010    100.2    100     
  1      5000     379800     0        70855    100.1    100     
  2      5020     379900     0        71840    100.1    0       
  3      5000     379800     0        71180    100.1    100     
  4      5000     379800     0        69762    100.2    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:34:45] W: 5000/s TotalW:404900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     404900     0        75534    100.2    0       
  1      5020     404900     0        75393    100.1    0       
  2      5000     404900     0        76448    100.1    0       
  3      5000     404800     0        75735    100.1    100     
  4      5020     404900     0        74226    100.2    0       
  TotalR: 25060/s QueueDepth(min=0, max=100)

[12:34:50] W: 5000/s TotalW:429900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     429900     0        79714    100.2    0       
  1      5000     429900     0        79577    100.1    0       
  2      4980     429800     0        80676    100.1    100     
  3      5020     429900     0        79907    100.1    0       
  4      4980     429800     0        78322    100.2    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[12:34:55] W: 5000/s TotalW:454900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     454900     0        84943    100.2    0       
  1      4980     454800     0        84801    100.1    100     
  2      5020     454900     0        85956    100.1    0       
  3      5000     454900     0        85133    100.1    0       
  4      5000     454800     0        83464    100.2    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:35:00] W: 5000/s TotalW:479900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     479900     0        89603    100.2    0       
  1      5000     479800     0        89462    100.1    100     
  2      4980     479800     0        90669    100.1    100     
  3      5000     479900     0        89802    100.1    0       
  4      5020     479900     0        88046    100.1    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[12:35:05] W: 5000/s TotalW:504900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     504900     0        93760    100.2    0       
  1      5020     504900     0        93633    100.1    0       
  2      5020     504900     0        94897    100.1    0       
  3      5000     504900     0        93984    100.1    0       
  4      5000     504900     0        92168    100.1    0       
  TotalR: 25040/s QueueDepth(min=0, max=0)

[12:35:10] W: 5000/s TotalW:529900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     529900     0        98546    100.2    0       
  1      4980     529800     0        98400    100.1    100     
  2      5000     529900     0        99721    100.1    0       
  3      5000     529900     0        98773    100.1    0       
  4      4980     529800     0        96853    100.1    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[12:35:15] W: 5000/s TotalW:554900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     554900     0        103771   100.1    0       
  1      5020     554900     0        103615   100.1    0       
  2      5000     554900     0        105000   100.1    0       
  3      5000     554900     0        104009   100.1    0       
  4      5020     554900     0        101995   100.1    0       
  TotalR: 25040/s QueueDepth(min=0, max=0)

[12:35:20] W: 5000/s TotalW:579900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     579900     0        108149   100.1    0       
  1      4980     579800     0        108007   100.1    100     
  2      5000     579900     0        109427   100.1    0       
  3      4980     579800     0        108408   100.1    100     
  4      4980     579800     0        106310   100.1    100     
  TotalR: 24940/s QueueDepth(min=0, max=100)

2025/10/03 12:35:25 [consumer g0 r0 p1] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 12:35:25 [consumer g2 r0 p1] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 12:35:25 [consumer g2 r2 p3] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 12:35:25 [consumer g4 r1 p2] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 12:35:25 [consumer g1 r1 p2] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 12:35:25 [consumer g4 r0 p1] claim err: context deadline exceeded (group=g4 batch=200)

=== Summary ===
Total Writes: 604400
Total Reads: 3022000
Total Updates: 0
Write Errors: 100
Read Errors: 0
Avg Write Throughput: 5036.67 rows/sec
Avg Read Throughput: 25183.33 rows/sec

Write Latencies (INSERT only):
  P50: 5.296127ms
  P95: 6.189055ms
  P99: 28.622847ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 0s
  P95: 0s
  P99: 0s

End-to-End Latencies (created_at → consumed):
  P50: 0s
  P95: 0s
  P99: 0s


=== PubSub Summary ===
Delivery Semantics: at-least-once (strict ordering)
Number of partitions: 4
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      112341
  Avg Poll Size:     100.13
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.530751ms
    P95: 4.644863ms
    P99: 25.640959ms

  End-to-End Latency (created_at → consumed):
    P50: 8.724479ms
    P95: 10.567679ms
    P99: 37.060607ms

-- Consumer Group 1 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      112221
  Avg Poll Size:     100.07
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.508223ms
    P95: 4.673535ms
    P99: 25.133055ms

  End-to-End Latency (created_at → consumed):
    P50: 8.732671ms
    P95: 10.526719ms
    P99: 37.814271ms

-- Consumer Group 2 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      2
  Empty Claims:      113692
  Avg Poll Size:     100.08
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.475455ms
    P95: 4.624383ms
    P99: 25.411583ms

  End-to-End Latency (created_at → consumed):
    P50: 8.708095ms
    P95: 10.485759ms
    P99: 36.470783ms

-- Consumer Group 3 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      0
  Empty Claims:      112630
  Avg Poll Size:     100.05
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.506175ms
    P95: 4.681727ms
    P99: 26.083327ms

  End-to-End Latency (created_at → consumed):
    P50: 8.724479ms
    P95: 10.608639ms
    P99: 35.422207ms

-- Consumer Group 4 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      2
  Empty Claims:      110453
  Avg Poll Size:     100.12
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.575807ms
    P95: 4.669439ms
    P99: 25.411583ms

  End-to-End Latency (created_at → consumed):
    P50: 8.806399ms
    P95: 10.616831ms
    P99: 35.389439ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 3022000
  Total Read Errors:     0
  Total Updates:         3022000
  Total Claim Errors:    6
  Total Empty Claims:    561337
  Avg Poll Size:         100.09
  Avg Throughput:        25183.33 reads/sec

  Read Latency (claim txn):
    P50: 3.520511ms
    P95: 4.665343ms
    P99: 25.526271ms

  End-to-End Latency (created_at → consumed):
    P50: 8.740863ms
    P95: 10.567679ms
    P99: 36.274175ms

2025/10/03 12:35:25 pubsub benchmark complete
```

# Run 3
```bash
ubuntu@ip-172-31-88-36:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s    --write-batch 100   --read-batch 200     --writers=10   --readers=4   --consumer-groups 5     --mode=pubsub --partitions 4 --duration=120s --throttle_writes=5000
2025/10/03 12:35:55 [pub info] successfully created 4 topicpartitions w/ counters + consumer_offsets
2025/10/03 12:35:55 [pub info] producers per partition [p1=3 p2=3 p3=2 p4=2]
[12:36:00] W: 5980/s TotalW:29900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5960     29800      0        4173     104.6    100     
  1      5960     29800      0        4169     104.2    100     
  2      5960     29800      0        4143     103.5    100     
  3      5960     29800      0        4115     104.6    100     
  4      5960     29800      0        4145     104.9    100     
  TotalR: 29799/s QueueDepth(min=100, max=100)

[12:36:05] W: 5000/s TotalW:54900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     54800      0        9450     102.4    100     
  1      5000     54800      0        9474     102.2    100     
  2      5000     54800      0        9365     101.9    100     
  3      5000     54800      0        9314     102.4    100     
  4      5000     54800      0        9433     102.6    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

[12:36:10] W: 5000/s TotalW:79900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     79800      0        14291    101.7    100     
  1      5000     79800      0        14316    101.5    100     
  2      5000     79800      0        14126    101.3    100     
  3      5000     79800      0        14065    101.7    100     
  4      5000     79800      0        14247    101.8    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:15] W: 5000/s TotalW:104900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     104800     0        18430    101.3    100     
  1      5000     104800     0        18451    101.2    100     
  2      5000     104800     0        18237    101.0    100     
  3      5000     104800     0        18157    101.3    100     
  4      5000     104800     0        18355    101.4    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:20] W: 5000/s TotalW:129900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     129800     0        23296    101.0    100     
  1      5000     129800     0        23311    100.9    100     
  2      5000     129800     0        23028    100.8    100     
  3      5000     129800     0        22952    101.0    100     
  4      5000     129800     0        23205    101.1    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:25] W: 5000/s TotalW:154900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     154800     0        28515    100.8    100     
  1      5000     154800     0        28514    100.8    100     
  2      5000     154800     0        28167    100.7    100     
  3      5000     154800     0        28083    100.8    100     
  4      5000     154800     0        28426    100.9    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:30] W: 4980/s TotalW:179800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     179800     0        32663    100.7    0       
  1      5000     179800     0        32660    100.7    0       
  2      5000     179800     0        32275    100.6    0       
  3      5000     179800     0        32183    100.7    0       
  4      5000     179800     0        32556    100.8    0       
  TotalR: 25000/s QueueDepth(min=0, max=0)

[12:36:35] W: 5020/s TotalW:204900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     204800     0        37065    100.6    100     
  1      5000     204800     0        37025    100.6    100     
  2      5000     204800     0        36621    100.5    100     
  3      5000     204800     0        36513    100.6    100     
  4      5000     204800     0        36911    100.7    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:40] W: 4980/s TotalW:229800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     229800     0        42334    100.6    0       
  1      5000     229800     0        42303    100.5    0       
  2      5000     229800     0        41799    100.4    0       
  3      5000     229800     0        41671    100.6    0       
  4      5000     229800     0        42158    100.6    0       
  TotalR: 24999/s QueueDepth(min=0, max=0)

[12:36:45] W: 5020/s TotalW:254900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     254800     0        47005    100.5    100     
  1      5000     254800     0        46986    100.5    100     
  2      5000     254800     0        46422    100.4    100     
  3      5000     254800     0        46295    100.5    100     
  4      5000     254800     0        46828    100.6    100     
  TotalR: 25002/s QueueDepth(min=100, max=100)

[12:36:50] W: 5000/s TotalW:279900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     279800     0        51070    100.5    100     
  1      5000     279800     0        51056    100.5    100     
  2      5000     279800     0        50430    100.4    100     
  3      5000     279800     0        50309    100.5    100     
  4      5000     279800     0        50859    100.5    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:36:55] W: 5000/s TotalW:304900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     304800     0        56015    100.4    100     
  1      5000     304800     0        56036    100.4    100     
  2      5000     304800     0        55325    100.3    100     
  3      5000     304800     0        55178    100.4    100     
  4      5000     304800     0        55832    100.5    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:37:00] W: 5000/s TotalW:329900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     329800     0        61158    100.4    100     
  1      5000     329800     0        61200    100.4    100     
  2      5000     329800     0        60412    100.3    100     
  3      5000     329800     0        60255    100.4    100     
  4      5000     329800     0        61008    100.4    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:37:05] W: 5000/s TotalW:354900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     354800     0        65274    100.4    100     
  1      5000     354800     0        65354    100.4    100     
  2      5000     354800     0        64518    100.3    100     
  3      5000     354800     0        64333    100.4    100     
  4      5000     354800     0        65108    100.4    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

[12:37:10] W: 4980/s TotalW:379800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     379800     0        69761    100.3    0       
  1      5000     379800     0        69831    100.3    0       
  2      5000     379800     0        68951    100.3    0       
  3      5000     379800     0        68746    100.3    0       
  4      5000     379800     0        69542    100.4    0       
  TotalR: 25000/s QueueDepth(min=0, max=0)

[12:37:15] W: 5020/s TotalW:404900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     404800     0        74922    100.3    100     
  1      5000     404800     0        75035    100.3    100     
  2      5000     404800     0        74043    100.2    100     
  3      5000     404800     0        73813    100.3    100     
  4      5000     404800     0        74664    100.3    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:37:20] W: 5000/s TotalW:429900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     429800     0        79498    100.3    100     
  1      5000     429800     0        79607    100.3    100     
  2      5000     429800     0        78567    100.2    100     
  3      5000     429800     0        78309    100.3    100     
  4      5000     429800     0        79193    100.3    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:37:25] W: 5000/s TotalW:454900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     454800     0        83648    100.3    100     
  1      5000     454800     0        83741    100.3    100     
  2      5000     454800     0        82675    100.2    100     
  3      5000     454800     0        82413    100.3    100     
  4      5000     454800     0        83303    100.3    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:37:30] W: 5000/s TotalW:479900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     479800     0        88809    100.3    100     
  1      5000     479800     0        88915    100.3    100     
  2      5000     479800     0        87772    100.2    100     
  3      5000     479800     0        87440    100.3    100     
  4      5000     479800     0        88441    100.3    100     
  TotalR: 25002/s QueueDepth(min=100, max=100)

[12:37:35] W: 5000/s TotalW:504900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     504800     0        91393    100.3    100     
  1      5000     504800     0        91499    100.3    100     
  2      5000     504800     0        90318    100.2    100     
  3      5000     504800     0        89968    100.3    100     
  4      5000     504800     0        90991    100.3    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:37:40] W: 5000/s TotalW:529900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     529800     0        95545    100.2    100     
  1      5000     529800     0        95665    100.2    100     
  2      5000     529800     0        94425    100.2    100     
  3      5000     529800     0        94066    100.2    100     
  4      5000     529800     0        95121    100.3    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[12:37:45] W: 5000/s TotalW:554900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     554800     0        100224   100.2    100     
  1      5000     554800     0        100365   100.2    100     
  2      5000     554800     0        99039    100.2    100     
  3      5000     554800     0        98681    100.2    100     
  4      5000     554800     0        99754    100.3    100     
  TotalR: 24999/s QueueDepth(min=100, max=100)

[12:37:50] W: 5000/s TotalW:579900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     579800     0        105452   100.2    100     
  1      5000     579800     0        105623   100.2    100     
  2      5000     579800     0        104184   100.2    100     
  3      5000     579800     0        103811   100.2    100     
  4      5000     579800     0        104928   100.2    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

2025/10/03 12:37:55 [consumer g0 r0 p1] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 12:37:55 [consumer g1 r0 p1] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 12:37:55 [consumer g3 r0 p1] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 12:37:55 [consumer g1 r2 p3] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 12:37:55 [consumer g3 r1 p2] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 12:37:55 [consumer g3 r2 p3] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 12:37:55 [consumer g1 r1 p2] claim err: context deadline exceeded (group=g1 batch=200)

=== Summary ===
Total Writes: 604400
Total Reads: 3022000
Total Updates: 0
Write Errors: 0
Read Errors: 0
Avg Write Throughput: 5036.67 rows/sec
Avg Read Throughput: 25183.33 rows/sec

Write Latencies (INSERT only):
  P50: 5.263359ms
  P95: 6.299647ms
  P99: 43.909119ms

Read Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):
  P50: 0s
  P95: 0s
  P99: 0s

End-to-End Latencies (created_at → consumed):
  P50: 0s
  P95: 0s
  P99: 0s


=== PubSub Summary ===
Delivery Semantics: at-least-once (strict ordering)
Number of partitions: 4
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      1
  Empty Claims:      109816
  Avg Poll Size:     100.22
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.553279ms
    P95: 4.722687ms
    P99: 46.759935ms

  End-to-End Latency (created_at → consumed):
    P50: 8.773631ms
    P95: 10.706943ms
    P99: 57.081855ms

-- Consumer Group 1 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      3
  Empty Claims:      109968
  Avg Poll Size:     100.22
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.553279ms
    P95: 4.706303ms
    P99: 46.432255ms

  End-to-End Latency (created_at → consumed):
    P50: 8.765439ms
    P95: 10.747903ms
    P99: 57.769983ms

-- Consumer Group 2 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      0
  Empty Claims:      108486
  Avg Poll Size:     100.17
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.581951ms
    P95: 4.710399ms
    P99: 47.120383ms

  End-to-End Latency (created_at → consumed):
    P50: 8.839167ms
    P95: 10.878975ms
    P99: 57.966591ms

-- Consumer Group 3 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      3
  Empty Claims:      108108
  Avg Poll Size:     100.22
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.598335ms
    P95: 4.759551ms
    P99: 47.513599ms

  End-to-End Latency (created_at → consumed):
    P50: 8.855551ms
    P95: 10.846207ms
    P99: 59.375615ms

-- Consumer Group 4 --
  Reads Completed:   604400
  Read Errors:       0
  Updates Completed: 604400
  Claim Errors:      0
  Empty Claims:      109226
  Avg Poll Size:     100.23
  Avg Throughput:    5036.67 reads/sec

  Read Latency (claim txn):
    P50: 3.555327ms
    P95: 4.694015ms
    P99: 46.071807ms

  End-to-End Latency (created_at → consumed):
    P50: 8.781823ms
    P95: 10.846207ms
    P99: 59.932671ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 3022000
  Total Read Errors:     0
  Total Updates:         3022000
  Total Claim Errors:    7
  Total Empty Claims:    545604
  Avg Poll Size:         100.21
  Avg Throughput:        25183.33 reads/sec

  Read Latency (claim txn):
    P50: 3.567615ms
    P95: 4.722687ms
    P99: 46.989311ms

  End-to-End Latency (created_at → consumed):
    P50: 8.806399ms
    P95: 10.805247ms
    P99: 58.064895ms

2025/10/03 12:37:55 pubsub benchmark complete
```
