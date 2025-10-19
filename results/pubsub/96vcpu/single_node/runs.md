# Run 1
```bash
ubuntu@ip-172-31-46-123:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  --report=5s    --write-batch=200   --read-batch=200     --writers=100   --readers=30   --consumer-groups=5        --mode=pubsub --partitions=30 --duration=120s --throttle_writes=250000
2025/10/03 18:26:58 [pub info] successfully created 30 topicpartitions w/ counters + consumer_offsets
2025/10/03 18:26:59 [pub info] spawned readers
2025/10/03 18:26:59 [pub info] producers per partition [p1=4 p2=4 p3=4 p4=4 p5=4 p6=4 p7=4 p8=4 p9=4 p10=4 p11=3 p12=3 p13=3 p14=3 p15=3 p16=3 p17=3 p18=3 p19=3 p20=3 p21=3 p22=3 p23=3 p24=3 p25=3 p26=3 p27=3 p28=3 p29=3 p30=3]
[18:27:04] W: 342554/s TotalW:1712800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      311995   1560000    0        1728     200.0    152800  
  1      329035   1645200    0        1602     200.0    67600   
  2      328275   1641400    0        2142     200.0    71400   
  3      326075   1630400    0        1704     200.0    82400   
  4      322395   1612000    0        1847     200.0    100800  
  TotalR: 1617773/s QueueDepth(min=67600, max=152800)

[18:27:09] W: 253283/s TotalW:2979200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      264323   2881600    0        5378     200.0    97600   
  1      263123   2960800    0        6457     200.0    18400   
  2      262443   2953600    0        7561     200.0    25600   
  3      267043   2965600    0        6353     200.0    13600   
  4      266523   2944600    0        6501     200.0    34600   
  TotalR: 1323457/s QueueDepth(min=13600, max=97600)

[18:27:14] W: 250280/s TotalW:4230600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      262880   4196000    0        11030    200.0    34600   
  1      252400   4222800    0        12899    200.0    7800    
  2      254640   4226800    0        14597    200.0    3800    
  3      251960   4225400    0        13173    200.0    5200    
  4      256000   4224600    0        13252    200.0    6000    
  TotalR: 1277881/s QueueDepth(min=3800, max=34600)

[18:27:19] W: 249800/s TotalW:5479600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      253880   5465400    0        16733    200.0    14200   
  1      250640   5476000    0        19135    200.0    3600    
  2      249560   5474600    0        21452    200.0    5000    
  3      249640   5473600    0        19720    200.0    6000    
  4      249240   5470800    0        19597    200.0    8800    
  TotalR: 1252962/s QueueDepth(min=3600, max=14200)

[18:27:24] W: 249999/s TotalW:6729600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251319   6722000    0        24197    200.0    7600    
  1      248639   6719200    0        27279    200.0    10400   
  2      249959   6724400    0        29983    200.0    5200    
  3      249399   6720600    0        27999    200.0    9000    
  4      250879   6725200    0        28099    200.0    4400    
  TotalR: 1250196/s QueueDepth(min=4400, max=10400)

[18:27:29] W: 250346/s TotalW:7981400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250506   7974600    0        30838    200.0    6800    
  1      251466   7976600    0        34530    200.0    4800    
  2      250546   7977200    0        37806    200.0    4200    
  3      250986   7975600    0        35549    200.0    5800    
  4      249866   7974600    0        35396    200.0    6800    
  TotalR: 1253372/s QueueDepth(min=4200, max=6800)

[18:27:34] W: 249527/s TotalW:9229000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248807   9218600    0        36275    200.0    10400   
  1      249007   9221600    0        40399    200.0    7400    
  2      248447   9219400    0        44147    200.0    9600    
  3      248727   9219200    0        41581    200.0    9800    
  4      249567   9222400    0        41423    200.0    6600    
  TotalR: 1244554/s QueueDepth(min=6600, max=10400)

[18:27:39] W: 249887/s TotalW:10478400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      247727   10457200   0        39915    200.0    21200   
  1      248807   10465600   0        44475    200.0    12800   
  2      249767   10468200   0        48505    200.0    10200   
  3      249767   10468000   0        45837    200.0    10400   
  4      248727   10466000   0        45520    200.0    12400   
  TotalR: 1244797/s QueueDepth(min=10200, max=21200)

[18:27:44] W: 250360/s TotalW:11730200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251880   11716600   0        45848    200.0    13600   
  1      251920   11725200   0        50937    200.0    5000    
  2      249880   11717600   0        55512    200.0    12600   
  3      250920   11722600   0        52609    200.0    7600    
  4      251560   11723800   0        52426    200.0    6400    
  TotalR: 1256159/s QueueDepth(min=5000, max=13600)

[18:27:49] W: 249840/s TotalW:12979400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251400   12973600   0        51411    200.0    5800    
  1      249800   12974200   0        56999    200.0    5200    
  2      251080   12973000   0        62086    200.0    6400    
  3      249840   12971800   0        58657    200.0    7600    
  4      250120   12974400   0        58554    200.0    5000    
  TotalR: 1252242/s QueueDepth(min=5000, max=7600)

[18:27:54] W: 250320/s TotalW:14231000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249600   14221600   0        57359    200.0    9400    
  1      248560   14217000   0        63469    200.0    14000   
  2      250040   14223200   0        69065    200.0    7800    
  3      250520   14224400   0        65401    200.0    6600    
  4      249440   14221600   0        65156    200.0    9400    
  TotalR: 1248158/s QueueDepth(min=6600, max=14000)

[18:27:59] W: 249396/s TotalW:15478000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249636   15469800   0        64757    200.0    8200    
  1      249916   15466600   0        71554    200.0    11400   
  2      247716   15461800   0        77413    200.0    16200   
  3      248116   15465000   0        74040    200.0    13000   
  4      249516   15469200   0        73395    200.0    8800    
  TotalR: 1244900/s QueueDepth(min=8200, max=16200)

[18:28:04] W: 248444/s TotalW:16720200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248204   16710800   0        71890    200.0    9400    
  1      248964   16711400   0        78990    200.0    8800    
  2      249884   16711200   0        85476    200.0    9000    
  3      249204   16711000   0        81921    200.0    9200    
  4      248364   16711000   0        81070    200.0    9200    
  TotalR: 1244622/s QueueDepth(min=8800, max=9400)

[18:28:09] W: 251959/s TotalW:17980000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251319   17967400   0        75159    200.0    12600   
  1      251679   17969800   0        82595    200.0    10200   
  2      252159   17972000   0        89460    200.0    8000    
  3      252199   17972000   0        85599    200.0    8000    
  4      250279   17962400   0        84942    200.0    17600   
  TotalR: 1257637/s QueueDepth(min=8000, max=17600)

[18:28:14] W: 250321/s TotalW:19231600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251001   19222400   0        81678    200.0    9200    
  1      250161   19220600   0        89507    200.0    11000   
  2      249441   19219200   0        97110    200.0    12400   
  3      249281   19218400   0        92650    200.0    13200   
  4      251681   19220800   0        92047    200.0    10800   
  TotalR: 1251563/s QueueDepth(min=9200, max=13200)

[18:28:19] W: 249159/s TotalW:20477400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249799   20471400   0        88318    200.0    6000    
  1      249479   20468000   0        96506    200.0    9400    
  2      249479   20466600   0        104903   200.0    10800   
  3      249999   20468400   0        99864    200.0    9000    
  4      248959   20465600   0        99218    200.0    11800   
  TotalR: 1247715/s QueueDepth(min=6000, max=11800)

[18:28:24] W: 250678/s TotalW:21730800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250078   21721800   0        93934    200.0    9000    
  1      251558   21725800   0        102649   200.0    5000    
  2      251797   21725600   0        111468   200.0    5200    
  3      251358   21725200   0        106111   200.0    5600    
  4      251518   21723200   0        105561   200.0    7600    
  TotalR: 1256308/s QueueDepth(min=5000, max=9000)

[18:28:29] W: 249523/s TotalW:22978400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250043   22972000   0        100531   200.0    6400    
  1      249363   22972600   0        109820   200.0    5800    
  2      249123   22971200   0        118859   200.0    7200    
  3      249803   22974200   0        113252   200.0    4200    
  4      250123   22973800   0        112697   200.0    4600    
  TotalR: 1248457/s QueueDepth(min=4200, max=7200)

[18:28:34] W: 250279/s TotalW:24229800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249799   24221000   0        107069   200.0    8800    
  1      250119   24223200   0        117063   200.0    6600    
  2      250039   24221400   0        126580   200.0    8400    
  3      249279   24220600   0        120539   200.0    9200    
  4      249599   24221800   0        119894   200.0    8000    
  TotalR: 1248836/s QueueDepth(min=6600, max=9200)

[18:28:39] W: 249440/s TotalW:25477000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248120   25461600   0        110851   200.0    15400   
  1      249360   25470000   0        121524   200.0    7000    
  2      249360   25468200   0        131167   200.0    8800    
  3      249240   25466800   0        124950   200.0    10200   
  4      248920   25466400   0        124428   200.0    10600   
  TotalR: 1245001/s QueueDepth(min=7000, max=15400)

[18:28:44] W: 250875/s TotalW:26731400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252875   26726000   0        116436   200.0    5400    
  1      250955   26724800   0        127492   200.0    6600    
  2      251915   26727800   0        137604   200.0    3600    
  3      251755   26725600   0        131138   200.0    5800    
  4      251995   26726400   0        130691   200.0    5000    
  TotalR: 1259496/s QueueDepth(min=3600, max=6600)

[18:28:49] W: 249845/s TotalW:27980600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249325   27972600   0        123145   200.0    8000    
  1      249645   27973000   0        134867   200.0    7600    
  2      249125   27973400   0        145095   200.0    7200    
  3      248525   27968200   0        138648   200.0    12400   
  4      248805   27970400   0        138042   200.0    10200   
  TotalR: 1245424/s QueueDepth(min=7200, max=12400)

[18:28:54] W: 249797/s TotalW:29229600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249757   29221400   0        129022   200.0    8200    
  1      249957   29222800   0        141351   200.0    6800    
  2      249557   29221200   0        152178   200.0    8400    
  3      250597   29221200   0        145568   200.0    8400    
  4      250237   29221600   0        144560   200.0    8000    
  TotalR: 1250106/s QueueDepth(min=6800, max=8400)

2025/10/03 18:28:56 [consumer g2 r3 p4] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g3 r9 p10] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g4 r4 p5] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g3 r10 p11] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g0 r11 p12] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g4 r3 p4] scan err: context deadline exceeded
2025/10/03 18:28:56 [consumer g3 r14 p15] claim err: context deadline exceeded (group=g3 batch=200)

=== Summary ===
Total Writes: 29551000
Total Reads: 147701600
Total Updates: 0
Write Errors: 3800
Read Errors: 0
Avg Write Throughput: 246258.33 rows/sec
Avg Read Throughput: 1230846.67 rows/sec

Write Latencies (INSERT only):
  P50: 11.001855ms
  P95: 50.397183ms
  P99: 157.810687ms

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
Number of partitions: 30
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   29540800
  Read Errors:       21
  Updates Completed: 29540800
  Claim Errors:      0
  Empty Claims:      130463
  Avg Poll Size:     200.00
  Avg Throughput:    246173.33 reads/sec

  Read Latency (claim txn):
    P50: 11.034623ms
    P95: 20.414463ms
    P99: 25.985023ms

  End-to-End Latency (created_at → consumed):
    P50: 25.608191ms
    P95: 383.254527ms
    P99: 1.434451967s

-- Consumer Group 1 --
  Reads Completed:   29540800
  Read Errors:       21
  Updates Completed: 29540800
  Claim Errors:      0
  Empty Claims:      142923
  Avg Poll Size:     200.00
  Avg Throughput:    246173.33 reads/sec

  Read Latency (claim txn):
    P50: 10.461183ms
    P95: 19.267583ms
    P99: 24.068095ms

  End-to-End Latency (created_at → consumed):
    P50: 23.740415ms
    P95: 205.258751ms
    P99: 611.844095ms

-- Consumer Group 2 --
  Reads Completed:   29541200
  Read Errors:       17
  Updates Completed: 29541200
  Claim Errors:      0
  Empty Claims:      153730
  Avg Poll Size:     200.00
  Avg Throughput:    246176.67 reads/sec

  Read Latency (claim txn):
    P50: 10.051583ms
    P95: 19.267583ms
    P99: 24.346623ms

  End-to-End Latency (created_at → consumed):
    P50: 23.560191ms
    P95: 215.089151ms
    P99: 709.885951ms

-- Consumer Group 3 --
  Reads Completed:   29538800
  Read Errors:       19
  Updates Completed: 29538800
  Claim Errors:      1
  Empty Claims:      147119
  Avg Poll Size:     200.00
  Avg Throughput:    246156.67 reads/sec

  Read Latency (claim txn):
    P50: 10.346495ms
    P95: 19.316735ms
    P99: 24.248319ms

  End-to-End Latency (created_at → consumed):
    P50: 23.822335ms
    P95: 213.254143ms
    P99: 537.919487ms

-- Consumer Group 4 --
  Reads Completed:   29540000
  Read Errors:       22
  Updates Completed: 29540000
  Claim Errors:      0
  Empty Claims:      146229
  Avg Poll Size:     200.00
  Avg Throughput:    246166.67 reads/sec

  Read Latency (claim txn):
    P50: 10.436607ms
    P95: 19.546111ms
    P99: 24.379391ms

  End-to-End Latency (created_at → consumed):
    P50: 23.937023ms
    P95: 222.560255ms
    P99: 874.512383ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 147701600
  Total Read Errors:     100
  Total Updates:         147701600
  Total Claim Errors:    1
  Total Empty Claims:    720464
  Avg Poll Size:         200.00
  Avg Throughput:        1230846.67 reads/sec

  Read Latency (claim txn):
    P50: 10.461183ms
    P95: 19.578879ms
    P99: 24.641535ms

  End-to-End Latency (created_at → consumed):
    P50: 24.100863ms
    P95: 230.293503ms
    P99: 851.443711ms

2025/10/03 18:28:56 pubsub benchmark complete
```
# Run 2
```bash
ubuntu@ip-172-31-46-123:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  --report=5s    --write-batch=200   --read-batch=200     --writers=100   --readers=30   --consumer-groups=5        --mode=pubsub --partitions=30 --duration=120s --throttle_writes=250000
2025/10/03 18:29:44 [pub info] successfully created 30 topicpartitions w/ counters + consumer_offsets
2025/10/03 18:29:45 [pub info] spawned readers
2025/10/03 18:29:45 [pub info] producers per partition [p1=4 p2=4 p3=4 p4=4 p5=4 p6=4 p7=4 p8=4 p9=4 p10=4 p11=3 p12=3 p13=3 p14=3 p15=3 p16=3 p17=3 p18=3 p19=3 p20=3 p21=3 p22=3 p23=3 p24=3 p25=3 p26=3 p27=3 p28=3 p29=3 p30=3]
[18:29:50] W: 345758/s TotalW:1728800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      319158   1595800    0        1829     200.0    133000  
  1      328878   1644400    0        1625     200.0    84400   
  2      321198   1606000    0        1304     200.0    122800  
  3      319358   1596800    0        1244     200.0    132000  
  4      326598   1633000    0        1345     200.0    95800   
  TotalR: 1615191/s QueueDepth(min=84400, max=133000)

[18:29:55] W: 250040/s TotalW:2979000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      267320   2932400    0        7386     200.0    46600   
  1      263440   2961600    0        7680     200.0    17400   
  2      266080   2936400    0        6745     200.0    42600   
  3      264800   2920800    0        6534     200.0    58200   
  4      265360   2959800    0        7108     200.0    19200   
  TotalR: 1326998/s QueueDepth(min=17400, max=58200)

[18:30:00] W: 250321/s TotalW:4230600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      256521   4215000    0        14861    200.0    15600   
  1      251081   4217000    0        15622    200.0    13600   
  2      254561   4209200    0        14081    200.0    21400   
  3      253201   4186800    0        13727    200.0    43800   
  4      252681   4223200    0        14737    200.0    7400    
  TotalR: 1268047/s QueueDepth(min=7400, max=43800)

[18:30:05] W: 249840/s TotalW:5479800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252200   5476000    0        21163    200.0    3800    
  1      251760   5475800    0        22015    200.0    4000    
  2      252560   5472000    0        20187    200.0    7800    
  3      254440   5459000    0        19782    200.0    20800   
  4      250440   5475400    0        21032    200.0    4400    
  TotalR: 1261402/s QueueDepth(min=3800, max=20800)

[18:30:10] W: 250238/s TotalW:6731000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250078   6726400    0        27344    200.0    4600    
  1      250038   6726000    0        28372    200.0    5000    
  2      250838   6726200    0        26316    200.0    4800    
  3      251398   6716000    0        25685    200.0    15000   
  4      250078   6725800    0        27323    200.0    5200    
  TotalR: 1252432/s QueueDepth(min=4600, max=15000)

[18:30:15] W: 249241/s TotalW:7977200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248001   7966400    0        34076    200.0    10800   
  1      247121   7961600    0        35368    200.0    15600   
  2      247561   7964000    0        32911    200.0    13200   
  3      248001   7956000    0        32154    200.0    21200   
  4      248161   7966600    0        34157    200.0    10600   
  TotalR: 1238847/s QueueDepth(min=10600, max=21200)

[18:30:20] W: 250679/s TotalW:9230600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251159   9222200    0        38129    200.0    8400    
  1      251799   9220600    0        39617    200.0    10000   
  2      248119   9204600    0        36816    200.0    26000   
  3      250079   9206400    0        36125    200.0    24200   
  4      251079   9222000    0        38278    200.0    8600    
  TotalR: 1252235/s QueueDepth(min=8400, max=26000)

[18:30:25] W: 249718/s TotalW:10479200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250118   10472800   0        44042    200.0    6400    
  1      250198   10471600   0        45819    200.0    7600    
  2      251518   10462200   0        42618    200.0    17000   
  3      250558   10459200   0        41754    200.0    20000   
  4      248318   10463600   0        44369    200.0    15600   
  TotalR: 1250709/s QueueDepth(min=6400, max=20000)

[18:30:30] W: 250163/s TotalW:11730000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249403   11719800   0        51403    200.0    10200   
  1      250443   11723800   0        53631    200.0    6200    
  2      252043   11722400   0        49886    200.0    7600    
  3      252083   11719600   0        48730    200.0    10400   
  4      252323   11725200   0        51773    200.0    4800    
  TotalR: 1256296/s QueueDepth(min=4800, max=10400)

[18:30:35] W: 249680/s TotalW:12978400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250320   12971400   0        58528    200.0    7000    
  1      249240   12970000   0        61011    200.0    8400    
  2      249560   12970200   0        57176    200.0    8200    
  3      249400   12966600   0        55808    200.0    11800   
  4      249320   12971800   0        59234    200.0    6600    
  TotalR: 1247838/s QueueDepth(min=6600, max=11800)

[18:30:40] W: 250069/s TotalW:14228800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250589   14224400   0        64874    200.0    4400    
  1      250709   14223600   0        67669    200.0    5200    
  2      250789   14224200   0        63348    200.0    4600    
  3      251589   14224600   0        61892    200.0    4200    
  4      250549   14224600   0        65723    200.0    4200    
  TotalR: 1254223/s QueueDepth(min=4200, max=5200)

[18:30:45] W: 250331/s TotalW:15480400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250171   15475200   0        71835    200.0    5200    
  1      250531   15476200   0        74983    200.0    4200    
  2      249851   15473400   0        70219    200.0    7000    
  3      250131   15475200   0        68617    200.0    5200    
  4      250331   15476200   0        72820    200.0    4200    
  TotalR: 1251016/s QueueDepth(min=4200, max=7000)

[18:30:50] W: 249559/s TotalW:16728200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      247039   16710400   0        77280    200.0    17800   
  1      247959   16716000   0        80259    200.0    12200   
  2      245759   16702200   0        75401    200.0    26000   
  3      244639   16698400   0        73606    200.0    29800   
  4      248199   16717200   0        78205    200.0    11000   
  TotalR: 1233596/s QueueDepth(min=11000, max=29800)

[18:30:55] W: 250426/s TotalW:17980400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252946   17975200   0        83202    200.0    5200    
  1      251906   17975600   0        86571    200.0    4800    
  2      253626   17970400   0        81291    200.0    10000   
  3      253786   17967400   0        79589    200.0    13000   
  4      251466   17974600   0        84670    200.0    5800    
  TotalR: 1263731/s QueueDepth(min=4800, max=13000)

[18:31:00] W: 250095/s TotalW:19230800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250415   19227200   0        89642    200.0    3600    
  1      250375   19227400   0        93291    200.0    3400    
  2      251215   19226400   0        87605    200.0    4400    
  3      251895   19226800   0        85829    200.0    4000    
  4      250375   19226400   0        91365    200.0    4400    
  TotalR: 1254275/s QueueDepth(min=3400, max=4400)

[18:31:05] W: 249560/s TotalW:20478600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248360   20469000   0        96858    200.0    9600    
  1      247600   20465400   0        100197   200.0    13200   
  2      248120   20467000   0        94610    200.0    11600   
  3      248760   20470600   0        92741    200.0    8000    
  4      249200   20472400   0        98303    200.0    6200    
  TotalR: 1242040/s QueueDepth(min=6200, max=13200)

[18:31:10] W: 250312/s TotalW:21730200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251152   21724800   0        104457   200.0    5400    
  1      252032   21725600   0        107785   200.0    4600    
  2      251512   21724600   0        101918   200.0    5600    
  3      250832   21724800   0        99876    200.0    5400    
  4      250672   21725800   0        106007   200.0    4400    
  TotalR: 1256198/s QueueDepth(min=4400, max=5600)

[18:31:15] W: 250088/s TotalW:22980600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250368   22976600   0        111260   200.0    4000    
  1      250248   22976800   0        114715   200.0    3800    
  2      250328   22976200   0        108518   200.0    4400    
  3      250128   22975400   0        106380   200.0    5200    
  4      249888   22975200   0        112845   200.0    5400    
  TotalR: 1250961/s QueueDepth(min=3800, max=5400)

[18:31:20] W: 249440/s TotalW:24227800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249400   24223600   0        118134   200.0    4200    
  1      249320   24223400   0        121894   200.0    4400    
  2      249680   24224600   0        115423   200.0    3200    
  3      249640   24223600   0        113211   200.0    4200    
  4      249800   24224200   0        119941   200.0    3600    
  TotalR: 1247839/s QueueDepth(min=3200, max=4400)

[18:31:25] W: 249757/s TotalW:25476600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      246197   25454600   0        122726   200.0    22000   
  1      247197   25459400   0        126893   200.0    17200   
  2      245557   25452400   0        120300   200.0    24200   
  3      243757   25442400   0        117955   200.0    34200   
  4      248797   25468200   0        124866   200.0    8400    
  TotalR: 1231504/s QueueDepth(min=8400, max=34200)

[18:31:30] W: 250804/s TotalW:26730600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252844   26718800   0        128578   200.0    11800   
  1      253004   26724400   0        133291   200.0    6200    
  2      254484   26724800   0        126126   200.0    5800    
  3      255204   26718400   0        123685   200.0    12200   
  4      250884   26722600   0        131286   200.0    8000    
  TotalR: 1266418/s QueueDepth(min=5800, max=12200)

[18:31:35] W: 250119/s TotalW:27981200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251319   27975400   0        134833   200.0    5800    
  1      250159   27975200   0        139786   200.0    6000    
  2      250279   27976200   0        132293   200.0    5000    
  3      251079   27973800   0        129615   200.0    7400    
  4      250439   27974800   0        137691   200.0    6400    
  TotalR: 1253275/s QueueDepth(min=5000, max=7400)

2025/10/03 18:31:40 [consumer g2 r21 p22] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g3 r16 p17] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g2 r4 p5] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g2 r16 p17] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g4 r17 p18] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g0 r9 p10] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g3 r0 p1] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g2 r0 p1] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g3 r7 p8] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g0 r17 p18] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g0 r2 p3] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g4 r5 p6] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g0 r4 p5] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g4 r6 p7] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g4 r13 p14] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g2 r7 p8] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g1 r0 p1] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g1 r11 p12] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g0 r16 p17] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g1 r10 p11] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g4 r2 p3] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g1 r18 p19] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g0 r10 p11] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g2 r14 p15] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g3 r3 p4] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g0 r13 p14] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g2 r13 p14] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g0 r26 p27] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g2 r22 p23] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g4 r12 p13] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g1 r15 p16] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g3 r19 p20] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g4 r9 p10] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g1 r16 p17] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g4 r10 p11] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g1 r14 p15] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g4 r19 p20] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g4 r18 p19] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g0 r22 p23] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g3 r10 p11] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g2 r17 p18] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g3 r29 p30] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g4 r3 p4] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g0 r20 p21] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g1 r8 p9] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g4 r14 p15] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g3 r12 p13] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g3 r6 p7] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g4 r20 p21] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g2 r26 p27] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g0 r21 p22] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g1 r2 p3] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g3 r21 p22] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:31:40 [consumer g2 r6 p7] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g2 r18 p19] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:31:40 [consumer g1 r21 p22] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g4 r26 p27] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:31:40 [consumer g1 r12 p13] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:31:40 [consumer g0 r14 p15] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/03 18:31:40 [consumer g4 r16 p17] claim err: context deadline exceeded (group=g4 batch=200)

=== Summary ===
Total Writes: 29007000
Total Reads: 144976000
Total Updates: 0
Write Errors: 2400
Read Errors: 0
Avg Write Throughput: 241725.00 rows/sec
Avg Read Throughput: 1208133.33 rows/sec

Write Latencies (INSERT only):
  P50: 10.412031ms
  P95: 46.235647ms
  P99: 129.761279ms

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
Number of partitions: 30
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   28994800
  Read Errors:       18
  Updates Completed: 28994800
  Claim Errors:      12
  Empty Claims:      140668
  Avg Poll Size:     200.00
  Avg Throughput:    241623.33 reads/sec

  Read Latency (claim txn):
    P50: 10.280959ms
    P95: 19.906559ms
    P99: 25.149439ms

  End-to-End Latency (created_at → consumed):
    P50: 23.068671ms
    P95: 246.677503ms
    P99: 1.162870783s

-- Consumer Group 1 --
  Reads Completed:   28995600
  Read Errors:       14
  Updates Completed: 28995600
  Claim Errors:      11
  Empty Claims:      145594
  Avg Poll Size:     200.00
  Avg Throughput:    241630.00 reads/sec

  Read Latency (claim txn):
    P50: 10.117119ms
    P95: 19.513343ms
    P99: 24.559615ms

  End-to-End Latency (created_at → consumed):
    P50: 22.609919ms
    P95: 215.744511ms
    P99: 814.743551ms

-- Consumer Group 2 --
  Reads Completed:   28993600
  Read Errors:       16
  Updates Completed: 28993600
  Claim Errors:      12
  Empty Claims:      137860
  Avg Poll Size:     200.00
  Avg Throughput:    241613.33 reads/sec

  Read Latency (claim txn):
    P50: 10.469375ms
    P95: 20.234239ms
    P99: 25.690111ms

  End-to-End Latency (created_at → consumed):
    P50: 23.724031ms
    P95: 345.243647ms
    P99: 1.003487231s

-- Consumer Group 3 --
  Reads Completed:   28996200
  Read Errors:       16
  Updates Completed: 28996200
  Claim Errors:      10
  Empty Claims:      135250
  Avg Poll Size:     200.00
  Avg Throughput:    241635.00 reads/sec

  Read Latency (claim txn):
    P50: 10.567679ms
    P95: 20.479999ms
    P99: 25.788415ms

  End-to-End Latency (created_at → consumed):
    P50: 23.920639ms
    P95: 486.014975ms
    P99: 1.495269375s

-- Consumer Group 4 --
  Reads Completed:   28995800
  Read Errors:       12
  Updates Completed: 28995800
  Claim Errors:      15
  Empty Claims:      143528
  Avg Poll Size:     200.00
  Avg Throughput:    241631.67 reads/sec

  Read Latency (claim txn):
    P50: 10.166271ms
    P95: 19.644415ms
    P99: 24.788991ms

  End-to-End Latency (created_at → consumed):
    P50: 22.790143ms
    P95: 223.870975ms
    P99: 696.254463ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 144976000
  Total Read Errors:     76
  Total Updates:         144976000
  Total Claim Errors:    60
  Total Empty Claims:    702900
  Avg Poll Size:         200.00
  Avg Throughput:        1208133.33 reads/sec

  Read Latency (claim txn):
    P50: 10.313727ms
    P95: 19.972095ms
    P99: 25.231359ms

  End-to-End Latency (created_at → consumed):
    P50: 23.199743ms
    P95: 262.799359ms
    P99: 1.088421887s

2025/10/03 18:31:40 pubsub benchmark complete
```
# Run 3
```bash
ubuntu@ip-172-31-46-123:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres  --report=5s    --write-batch=200   --read-batch=200     --writers=100   --readers=30   --consumer-groups=5        --mode=pubsub --partitions=30 --duration=120s --throttle_writes=250000
2025/10/03 18:32:05 [pub info] successfully created 30 topicpartitions w/ counters + consumer_offsets
2025/10/03 18:32:06 [pub info] spawned readers
2025/10/03 18:32:06 [pub info] producers per partition [p1=4 p2=4 p3=4 p4=4 p5=4 p6=4 p7=4 p8=4 p9=4 p10=4 p11=3 p12=3 p13=3 p14=3 p15=3 p16=3 p17=3 p18=3 p19=3 p20=3 p21=3 p22=3 p23=3 p24=3 p25=3 p26=3 p27=3 p28=3 p29=3 p30=3]
[18:32:11] W: 345953/s TotalW:1729800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      330313   1651600    0        1430     200.0    78200   
  1      327673   1638400    0        681      200.0    91400   
  2      331633   1658200    0        1040     200.0    71600   
  3      324113   1620600    0        815      200.0    109200  
  4      327673   1638400    0        1036     200.0    91400   
  TotalR: 1641404/s QueueDepth(min=71600, max=109200)

[18:32:16] W: 249961/s TotalW:2979600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      263521   2969200    0        7225     200.0    10400   
  1      266201   2969400    0        5414     200.0    10200   
  2      260081   2958600    0        6261     200.0    21000   
  3      262001   2930600    0        5640     200.0    49000   
  4      265681   2966800    0        6421     200.0    12800   
  TotalR: 1317484/s QueueDepth(min=10200, max=49000)

[18:32:21] W: 250402/s TotalW:4231600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      251562   4227000    0        14531    200.0    4600    
  1      251482   4226800    0        11724    200.0    4800    
  2      253722   4227200    0        12926    200.0    4400    
  3      258442   4222800    0        11799    200.0    8800    
  4      251322   4223400    0        13346    200.0    8200    
  TotalR: 1266528/s QueueDepth(min=4400, max=8800)

[18:32:26] W: 249923/s TotalW:5481200
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250123   5477600    0        21026    200.0    3600    
  1      249963   5476600    0        17696    200.0    4600    
  2      249923   5476800    0        18972    200.0    4400    
  3      250643   5476000    0        17788    200.0    5200    
  4      250643   5476600    0        19673    200.0    4600    
  TotalR: 1251293/s QueueDepth(min=3600, max=5200)

[18:32:31] W: 249907/s TotalW:6730800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249547   6725400    0        28599    200.0    5400    
  1      248988   6721600    0        24520    200.0    9200    
  2      249667   6725200    0        26083    200.0    5600    
  3      248508   6718600    0        24436    200.0    12200   
  4      248108   6717200    0        26740    200.0    13600   
  TotalR: 1244818/s QueueDepth(min=5400, max=13600)

[18:32:36] W: 249812/s TotalW:7979800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248412   7967400    0        36025    200.0    12400   
  1      249172   7967400    0        31240    200.0    12400   
  2      249172   7971000    0        33231    200.0    8800    
  3      250013   7968600    0        31057    200.0    11200   
  4      249612   7965200    0        33662    200.0    14600   
  TotalR: 1246382/s QueueDepth(min=8800, max=14600)

[18:32:41] W: 250240/s TotalW:9231000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249560   9215200    0        39355    200.0    15800   
  1      250520   9220000    0        34072    200.0    11000   
  2      249640   9219200    0        36405    200.0    11800   
  3      247840   9207800    0        34114    200.0    23200   
  4      250800   9219200    0        36940    200.0    11800   
  TotalR: 1248358/s QueueDepth(min=11000, max=23200)

[18:32:46] W: 250000/s TotalW:10481000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252240   10476400   0        45507    200.0    4600    
  1      250960   10474800   0        39890    200.0    6200    
  2      251400   10476200   0        42272    200.0    4800    
  3      253400   10474800   0        39703    200.0    6200    
  4      251440   10476400   0        42976    200.0    4600    
  TotalR: 1259440/s QueueDepth(min=4600, max=6200)

[18:32:51] W: 249919/s TotalW:11730600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      248879   11720800   0        52068    200.0    9800    
  1      248359   11716600   0        46025    200.0    14000   
  2      249679   11724600   0        48773    200.0    6000    
  3      248079   11715200   0        45854    200.0    15400   
  4      248759   11720200   0        49552    200.0    10400   
  TotalR: 1243756/s QueueDepth(min=6000, max=15400)

[18:32:56] W: 249876/s TotalW:12980000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250756   12974600   0        59410    200.0    5400    
  1      251356   12973400   0        52748    200.0    6600    
  2      250236   12975800   0        55572    200.0    4200    
  3      251916   12974800   0        52422    200.0    5200    
  4      250356   12972000   0        56566    200.0    8000    
  TotalR: 1254622/s QueueDepth(min=4200, max=8000)

[18:33:01] W: 250125/s TotalW:14230600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250245   14225800   0        66486    200.0    4800    
  1      250365   14225200   0        59287    200.0    5400    
  2      250125   14226400   0        62144    200.0    4200    
  3      250325   14226400   0        58931    200.0    4200    
  4      250805   14226000   0        63348    200.0    4600    
  TotalR: 1251866/s QueueDepth(min=4200, max=5400)

[18:33:06] W: 249439/s TotalW:15477800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249559   15473600   0        73391    200.0    4200    
  1      248999   15470200   0        65467    200.0    7600    
  2      249119   15472000   0        68589    200.0    5800    
  3      249319   15473000   0        65011    200.0    4800    
  4      248999   15471000   0        70178    200.0    6800    
  TotalR: 1245996/s QueueDepth(min=4200, max=7600)

[18:33:11] W: 250641/s TotalW:16731000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      247041   16708800   0        78878    200.0    22200   
  1      249161   16716000   0        70151    200.0    15000   
  2      248401   16714000   0        73697    200.0    17000   
  3      246001   16703000   0        69498    200.0    28000   
  4      249601   16719000   0        75297    200.0    12000   
  TotalR: 1240203/s QueueDepth(min=12000, max=28000)

[18:33:16] W: 250199/s TotalW:17982000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      253519   17976400   0        85864    200.0    5600    
  1      252159   17976800   0        76351    200.0    5200    
  2      252479   17976400   0        80143    200.0    5600    
  3      254679   17976400   0        75539    200.0    5600    
  4      251439   17976200   0        82044    200.0    5800    
  TotalR: 1264276/s QueueDepth(min=5200, max=5800)

[18:33:21] W: 249601/s TotalW:19230000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249721   19225000   0        92152    200.0    5000    
  1      248041   19217000   0        81794    200.0    13000   
  2      249121   19222000   0        86067    200.0    8000    
  3      249401   19223400   0        80956    200.0    6600    
  4      249841   19225400   0        87957    200.0    4600    
  TotalR: 1246126/s QueueDepth(min=4600, max=13000)

[18:33:26] W: 249998/s TotalW:20480000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249678   20473400   0        99788    200.0    6600    
  1      251078   20472400   0        88420    200.0    7600    
  2      249838   20471200   0        93070    200.0    8800    
  3      249518   20471000   0        87719    200.0    9000    
  4      249198   20471400   0        95395    200.0    8600    
  TotalR: 1249308/s QueueDepth(min=6600, max=9000)

[18:33:31] W: 250202/s TotalW:21731000
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249722   21722000   0        106605   200.0    9000    
  1      250562   21725200   0        94838    200.0    5800    
  2      250922   21725800   0        99474    200.0    5200    
  3      250762   21724800   0        93985    200.0    6200    
  4      250242   21722600   0        102060   200.0    8400    
  TotalR: 1252209/s QueueDepth(min=5200, max=9000)

[18:33:36] W: 249879/s TotalW:22980400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250599   22975000   0        112559   200.0    5400    
  1      249799   22974200   0        100337   200.0    6200    
  2      249559   22973600   0        104938   200.0    6800    
  3      249599   22972800   0        99293    200.0    7600    
  4      250519   22975200   0        107946   200.0    5200    
  TotalR: 1250076/s QueueDepth(min=5200, max=7600)

[18:33:41] W: 249641/s TotalW:24228600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      247601   24213000   0        117862   200.0    15600   
  1      247361   24211000   0        104964   200.0    17600   
  2      248481   24216000   0        109787   200.0    12600   
  3      247001   24207800   0        104018   200.0    20800   
  4      247481   24212600   0        112898   200.0    16000   
  TotalR: 1237924/s QueueDepth(min=12600, max=20800)

[18:33:46] W: 250399/s TotalW:25480600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      252559   25475800   0        124414   200.0    4800    
  1      252159   25471800   0        110699   200.0    8800    
  2      251679   25474400   0        115697   200.0    6200    
  3      252119   25468400   0        109390   200.0    12200   
  4      251999   25472600   0        118909   200.0    8000    
  TotalR: 1260515/s QueueDepth(min=4800, max=12200)

[18:33:51] W: 250201/s TotalW:26731600
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      250201   26726800   0        131033   200.0    4800    
  1      251001   26726800   0        116558   200.0    4800    
  2      250281   26725800   0        121670   200.0    5800    
  3      251121   26724000   0        115168   200.0    7600    
  4      251001   26727600   0        125101   200.0    4000    
  TotalR: 1253607/s QueueDepth(min=4000, max=7600)

[18:33:56] W: 250021/s TotalW:27981800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      249821   27976000   0        138255   200.0    5800    
  1      249861   27976200   0        123289   200.0    5600    
  2      249461   27973200   0        128444   200.0    8600    
  3      250501   27976600   0        121687   200.0    5200    
  4      249501   27975200   0        131978   200.0    6600    
  TotalR: 1249143/s QueueDepth(min=5200, max=8600)

2025/10/03 18:34:00 [consumer g3 r14 p15] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:34:00 [consumer g4 r16 p17] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:34:00 [consumer g3 r16 p17] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:34:00 [consumer g1 r26 p27] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:34:00 [consumer g4 r28 p29] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:34:00 [consumer g3 r23 p24] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/03 18:34:00 [consumer g2 r1 p2] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/03 18:34:00 [consumer g1 r23 p24] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/03 18:34:00 [consumer g4 r4 p5] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/03 18:34:00 [consumer g3 r7 p8] scan err: context deadline exceeded

=== Summary ===
Total Writes: 28996600
Total Reads: 144931400
Total Updates: 0
Write Errors: 4400
Read Errors: 0
Avg Write Throughput: 241638.33 rows/sec
Avg Read Throughput: 1207761.67 rows/sec

Write Latencies (INSERT only):
  P50: 10.428415ms
  P95: 44.826623ms
  P99: 125.763583ms

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
Number of partitions: 30
Duration: 2m0s (120.0s)

-- Consumer Group 0 --
  Reads Completed:   28985200
  Read Errors:       1
  Updates Completed: 28985200
  Claim Errors:      0
  Empty Claims:      144622
  Avg Poll Size:     200.00
  Avg Throughput:    241543.33 reads/sec

  Read Latency (claim txn):
    P50: 10.280959ms
    P95: 18.726911ms
    P99: 23.560191ms

  End-to-End Latency (created_at → consumed):
    P50: 22.331391ms
    P95: 208.797695ms
    P99: 542.638079ms

-- Consumer Group 1 --
  Reads Completed:   28987800
  Read Errors:       6
  Updates Completed: 28987800
  Claim Errors:      2
  Empty Claims:      129175
  Avg Poll Size:     200.00
  Avg Throughput:    241565.00 reads/sec

  Read Latency (claim txn):
    P50: 10.649599ms
    P95: 19.021823ms
    P99: 24.199167ms

  End-to-End Latency (created_at → consumed):
    P50: 23.085055ms
    P95: 230.948863ms
    P99: 558.366719ms

-- Consumer Group 2 --
  Reads Completed:   28988800
  Read Errors:       6
  Updates Completed: 28988800
  Claim Errors:      1
  Empty Claims:      134700
  Avg Poll Size:     200.00
  Avg Throughput:    241573.33 reads/sec

  Read Latency (claim txn):
    P50: 10.665983ms
    P95: 19.038207ms
    P99: 23.904255ms

  End-to-End Latency (created_at → consumed):
    P50: 22.839295ms
    P95: 214.564863ms
    P99: 549.453823ms

-- Consumer Group 3 --
  Reads Completed:   28981600
  Read Errors:       4
  Updates Completed: 28981600
  Claim Errors:      3
  Empty Claims:      127270
  Avg Poll Size:     200.00
  Avg Throughput:    241513.33 reads/sec

  Read Latency (claim txn):
    P50: 10.780671ms
    P95: 19.169279ms
    P99: 24.510463ms

  End-to-End Latency (created_at → consumed):
    P50: 24.117247ms
    P95: 319.029247ms
    P99: 841.482239ms

-- Consumer Group 4 --
  Reads Completed:   28988000
  Read Errors:       4
  Updates Completed: 28988000
  Claim Errors:      3
  Empty Claims:      138335
  Avg Poll Size:     200.00
  Avg Throughput:    241566.67 reads/sec

  Read Latency (claim txn):
    P50: 10.444799ms
    P95: 18.710527ms
    P99: 23.592959ms

  End-to-End Latency (created_at → consumed):
    P50: 22.773759ms
    P95: 227.934207ms
    P99: 578.289663ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 144931400
  Total Read Errors:     21
  Total Updates:         144931400
  Total Claim Errors:    9
  Total Empty Claims:    674102
  Avg Poll Size:         200.00
  Avg Throughput:        1207761.67 reads/sec

  Read Latency (claim txn):
    P50: 10.567679ms
    P95: 18.939903ms
    P99: 23.969791ms

  End-to-End Latency (created_at → consumed):
    P50: 23.003135ms
    P95: 233.308159ms
    P99: 620.232703ms

2025/10/03 18:34:00 pubsub benchmark complete
ubuntu@ip-172-31-46-123:/tmp/postgres-queue-benchma
```