# Run 1
```bash
ubuntu@ip-172-31-89-39:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s    --write-batch 100   --read-batch 200     --writers=10   --readers=4   --consumer-groups 5     --mode=pubsub --partitions 4 --duration=300s --throttle_writes=5000
2025/10/05 20:52:48 [pub info] successfully created 4 topicpartitions w/ counters + consumer_offsets
2025/10/05 20:52:49 [pub info] spawned readers
2025/10/05 20:52:49 [pub info] producers per partition [p1=3 p2=3 p3=2 p4=2]
[20:52:54] W: 6980/s TotalW:34900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      6960     34800      0        5287     103.9    100     
  1      6960     34800      0        5033     104.8    100     
  2      6960     34800      0        5518     103.9    100     
  3      6960     34800      0        4841     104.8    100     
  4      6980     34900      0        5543     103.6    0       
  TotalR: 34820/s QueueDepth(min=0, max=100)

[20:52:59] W: 5000/s TotalW:59900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     59800      0        10041    102.2    100     
  1      5020     59900      0        9531     102.7    0       
  2      5000     59800      0        10517    102.2    100     
  3      5000     59800      0        9193     102.7    100     
  4      4980     59800      0        10558    102.0    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:53:04] W: 5000/s TotalW:84900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     84800      0        13843    101.6    100     
  1      4980     84800      0        13229    101.9    100     
  2      5000     84800      0        14431    101.6    100     
  3      5000     84800      0        12735    101.9    100     
  4      5000     84800      0        14482    101.4    100     
  TotalR: 24980/s QueueDepth(min=100, max=100)

[20:53:09] W: 5000/s TotalW:109900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     109800     0        18027    101.2    100     
  1      5000     109800     0        17254    101.5    100     
  2      5000     109800     0        18790    101.2    100     
  3      5000     109800     0        16585    101.5    100     
  4      5020     109900     0        18833    101.1    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:53:14] W: 5000/s TotalW:134900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     134900     0        22852    101.0    0       
  1      5000     134800     0        21810    101.2    100     
  2      5000     134800     0        23861    101.0    100     
  3      5000     134800     0        20976    101.2    100     
  4      5000     134900     0        23897    100.9    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:53:19] W: 5000/s TotalW:159900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     159900     0        27025    100.9    0       
  1      5020     159900     0        25828    101.0    0       
  2      5000     159800     0        28181    100.8    100     
  3      5000     159800     0        24795    101.0    100     
  4      5000     159900     0        28224    100.8    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:53:24] W: 5000/s TotalW:184900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     184900     0        30802    100.8    0       
  1      4980     184800     0        29495    100.9    100     
  2      5000     184800     0        32049    100.7    100     
  3      5020     184900     0        28285    100.9    0       
  4      5000     184900     0        32095    100.7    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:53:29] W: 5000/s TotalW:209900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     209900     0        35548    100.7    0       
  1      5020     209900     0        33990    100.8    0       
  2      5020     209900     0        37039    100.6    0       
  3      5000     209900     0        32623    100.8    0       
  4      5000     209900     0        37085    100.6    0       
  TotalR: 25040/s QueueDepth(min=0, max=0)

[20:53:34] W: 5000/s TotalW:234900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     234800     0        40174    100.6    100     
  1      5000     234900     0        38381    100.7    0       
  2      4980     234800     0        41876    100.6    100     
  3      4980     234800     0        36856    100.7    100     
  4      5000     234900     0        41945    100.5    0       
  TotalR: 24940/s QueueDepth(min=0, max=100)

[20:53:39] W: 5000/s TotalW:259900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     259900     0        43626    100.5    0       
  1      4980     259800     0        41738    100.7    100     
  2      5000     259800     0        45423    100.5    100     
  3      5000     259800     0        40060    100.6    100     
  4      5000     259900     0        45509    100.5    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:53:44] W: 5000/s TotalW:284900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     284800     0        47919    100.5    100     
  1      5000     284800     0        45839    100.6    100     
  2      5000     284800     0        49888    100.5    100     
  3      5000     284800     0        44001    100.6    100     
  4      5000     284900     0        49988    100.4    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:53:49] W: 5000/s TotalW:309900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     309800     0        52687    100.5    100     
  1      5020     309900     0        50347    100.6    0       
  2      5000     309800     0        54915    100.4    100     
  3      5000     309800     0        48360    100.5    100     
  4      5000     309900     0        55013    100.4    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:53:54] W: 5000/s TotalW:334900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     334800     0        55731    100.5    100     
  1      4980     334800     0        53285    100.6    100     
  2      5000     334800     0        58068    100.5    100     
  3      5000     334800     0        51172    100.5    100     
  4      5000     334900     0        58179    100.4    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:53:59] W: 5000/s TotalW:359900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     359900     0        59580    100.4    0       
  1      5020     359900     0        57006    100.5    0       
  2      5000     359800     0        62039    100.4    100     
  3      5000     359800     0        54732    100.5    100     
  4      5000     359900     0        62147    100.4    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:54:04] W: 5000/s TotalW:384900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     384900     0        64359    100.4    0       
  1      5000     384900     0        61524    100.5    0       
  2      5020     384900     0        67062    100.4    0       
  3      5000     384800     0        59090    100.5    100     
  4      5000     384900     0        67183    100.4    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:54:09] W: 5000/s TotalW:409900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     409800     0        68813    100.4    100     
  1      4980     409800     0        65749    100.5    100     
  2      4980     409800     0        71685    100.4    100     
  3      5000     409800     0        63155    100.4    100     
  4      5000     409900     0        71829    100.4    0       
  TotalR: 24940/s QueueDepth(min=0, max=100)

[20:54:14] W: 5000/s TotalW:434900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     434900     0        72602    100.4    0       
  1      5000     434800     0        69413    100.4    100     
  2      5000     434800     0        75586    100.3    100     
  3      5000     434800     0        66667    100.4    100     
  4      5000     434900     0        75720    100.3    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:54:19] W: 5000/s TotalW:459900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     459900     0        76993    100.3    0       
  1      5020     459900     0        73610    100.4    0       
  2      5020     459900     0        80176    100.3    0       
  3      5020     459900     0        70696    100.4    0       
  4      5000     459900     0        80312    100.3    0       
  TotalR: 25060/s QueueDepth(min=0, max=0)

[20:54:24] W: 5000/s TotalW:484900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     484900     0        81743    100.3    0       
  1      5000     484900     0        78100    100.4    0       
  2      4980     484800     0        85163    100.3    100     
  3      5000     484900     0        75042    100.4    0       
  4      5000     484900     0        85292    100.3    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:54:29] W: 5000/s TotalW:509900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     509900     0        85282    100.3    0       
  1      5000     509900     0        81549    100.4    0       
  2      5000     509800     0        88811    100.3    100     
  3      4980     509800     0        78340    100.4    100     
  4      5000     509900     0        88941    100.3    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:54:34] W: 5000/s TotalW:534900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     534900     0        89413    100.3    0       
  1      5000     534900     0        85530    100.4    0       
  2      5020     534900     0        93053    100.3    0       
  3      5000     534800     0        82140    100.3    100     
  4      5000     534900     0        93202    100.3    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:54:39] W: 5000/s TotalW:559900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     559800     0        94132    100.3    100     
  1      4980     559800     0        89980    100.3    100     
  2      4980     559800     0        98016    100.3    100     
  3      5000     559800     0        86437    100.3    100     
  4      4980     559800     0        98166    100.3    100     
  TotalR: 24920/s QueueDepth(min=100, max=100)

[20:54:44] W: 5000/s TotalW:584900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     584800     0        98170    100.3    100     
  1      5000     584800     0        93859    100.3    100     
  2      5000     584800     0        102156   100.3    100     
  3      5000     584800     0        90137    100.3    100     
  4      5020     584900     0        102314   100.3    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:54:49] W: 5000/s TotalW:609900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     609900     0        101979   100.3    0       
  1      5020     609900     0        97533    100.3    0       
  2      5000     609800     0        106047   100.2    100     
  3      5000     609800     0        93650    100.3    100     
  4      5000     609900     0        106217   100.2    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:54:54] W: 5000/s TotalW:634900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     634800     0        106496   100.3    100     
  1      4980     634800     0        101826   100.3    100     
  2      5000     634800     0        110767   100.2    100     
  3      5000     634800     0        97786    100.3    100     
  4      4980     634800     0        110972   100.2    100     
  TotalR: 24940/s QueueDepth(min=100, max=100)

[20:54:59] W: 5000/s TotalW:659900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     659800     0        109759   100.2    100     
  1      5000     659800     0        104909   100.3    100     
  2      5000     659800     0        114194   100.2    100     
  3      5020     659900     0        100779   100.3    0       
  4      5020     659900     0        114403   100.2    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:55:04] W: 5000/s TotalW:684900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     684800     0        113553   100.2    100     
  1      5020     684900     0        108597   100.3    0       
  2      5000     684800     0        118102   100.2    100     
  3      4980     684800     0        104314   100.3    100     
  4      5000     684900     0        118325   100.2    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:55:09] W: 5000/s TotalW:709900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     709800     0        117652   100.2    100     
  1      5000     709900     0        112541   100.3    0       
  2      5000     709800     0        122356   100.2    100     
  3      5000     709800     0        108094   100.3    100     
  4      5000     709900     0        122588   100.2    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:55:14] W: 5000/s TotalW:734900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     734800     0        122456   100.2    100     
  1      5000     734900     0        117078   100.3    0       
  2      5000     734800     0        127398   100.2    100     
  3      5000     734800     0        112477   100.2    100     
  4      5000     734900     0        127648   100.2    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:55:19] W: 5000/s TotalW:759900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     759800     0        126655   100.2    100     
  1      4980     759800     0        121108   100.3    100     
  2      5000     759800     0        131781   100.2    100     
  3      5000     759800     0        116352   100.2    100     
  4      5000     759900     0        132017   100.2    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:55:24] W: 5000/s TotalW:784900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     784800     0        130452   100.2    100     
  1      5000     784800     0        124793   100.2    100     
  2      5000     784800     0        135676   100.2    100     
  3      5020     784900     0        119861   100.2    0       
  4      4980     784800     0        135929   100.2    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:55:29] W: 5000/s TotalW:809900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     809800     0        134659   100.2    100     
  1      5020     809900     0        128780   100.2    0       
  2      5020     809900     0        140088   100.2    0       
  3      5000     809900     0        123718   100.2    0       
  4      5020     809900     0        140348   100.2    0       
  TotalR: 25060/s QueueDepth(min=0, max=100)

[20:55:34] W: 5000/s TotalW:834900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     834800     0        139169   100.2    100     
  1      4980     834800     0        133075   100.2    100     
  2      4980     834800     0        144802   100.2    100     
  3      4980     834800     0        127850   100.2    100     
  4      4980     834800     0        145073   100.2    100     
  TotalR: 24920/s QueueDepth(min=100, max=100)

[20:55:39] W: 4520/s TotalW:857500
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4380     856700     0        142097   100.2    800     
  1      4400     856800     0        135918   100.2    700     
  2      4380     856700     0        147821   100.2    800     
  3      4380     856700     0        130566   100.2    800     
  4      4380     856700     0        148091   100.2    800     
  TotalR: 21919/s QueueDepth(min=700, max=800)

[20:55:44] W: 5480/s TotalW:884900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5640     884900     0        146244   100.3    0       
  1      5620     884900     0        139901   100.4    0       
  2      5640     884900     0        152139   100.3    0       
  3      5640     884900     0        134366   100.4    0       
  4      5640     884900     0        152410   100.3    0       
  TotalR: 28182/s QueueDepth(min=0, max=0)

[20:55:49] W: 5000/s TotalW:909900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     909900     0        151024   100.3    0       
  1      4980     909800     0        144417   100.4    100     
  2      5000     909900     0        157152   100.3    0       
  3      5000     909900     0        138706   100.3    0       
  4      5000     909900     0        157435   100.3    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:55:54] W: 5000/s TotalW:934900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     934900     0        155094   100.3    0       
  1      5000     934800     0        148338   100.4    100     
  2      5000     934900     0        161360   100.3    0       
  3      4980     934800     0        142465   100.3    100     
  4      4980     934800     0        161646   100.3    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[20:55:59] W: 5000/s TotalW:959900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     959900     0        158897   100.3    0       
  1      5000     959800     0        152032   100.3    100     
  2      5000     959900     0        165250   100.3    0       
  3      5020     959900     0        145982   100.3    0       
  4      5020     959900     0        165535   100.3    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:56:04] W: 5000/s TotalW:984900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     984900     0        163137   100.3    0       
  1      5020     984900     0        156058   100.3    0       
  2      5000     984900     0        169699   100.3    0       
  3      5000     984900     0        149865   100.3    0       
  4      5000     984900     0        169996   100.3    0       
  TotalR: 25020/s QueueDepth(min=0, max=0)

[20:56:09] W: 5000/s TotalW:1009900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1009900    0        167643   100.3    0       
  1      4980     1009800    0        160351   100.3    100     
  2      5000     1009900    0        174429   100.3    0       
  3      5000     1009900    0        154001   100.3    0       
  4      5000     1009900    0        174732   100.3    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:56:14] W: 5000/s TotalW:1034900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1034900    0        171429   100.3    0       
  1      5000     1034800    0        164032   100.3    100     
  2      5000     1034900    0        178330   100.3    0       
  3      5000     1034900    0        157498   100.3    0       
  4      4980     1034800    0        178641   100.3    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:56:19] W: 5000/s TotalW:1059900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1059800    0        175765   100.3    100     
  1      5000     1059800    0        168161   100.3    100     
  2      5000     1059900    0        182857   100.3    0       
  3      5000     1059900    0        161494   100.3    0       
  4      5000     1059800    0        183181   100.3    100     
  TotalR: 24979/s QueueDepth(min=0, max=100)

[20:56:24] W: 5000/s TotalW:1084900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1084800    0        180518   100.3    100     
  1      5020     1084900    0        172649   100.3    0       
  2      5000     1084900    0        187866   100.2    0       
  3      5000     1084900    0        165820   100.3    0       
  4      5020     1084900    0        188187   100.2    0       
  TotalR: 25041/s QueueDepth(min=0, max=100)

[20:56:29] W: 5000/s TotalW:1109900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1109900    0        184443   100.3    0       
  1      4980     1109800    0        176450   100.3    100     
  2      5000     1109900    0        191930   100.2    0       
  3      4980     1109800    0        169454   100.3    100     
  4      5000     1109900    0        192262   100.2    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:56:34] W: 5000/s TotalW:1134900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1134800    0        187818   100.3    100     
  1      5000     1134800    0        179705   100.3    100     
  2      4980     1134800    0        195405   100.3    100     
  3      5000     1134800    0        172566   100.3    100     
  4      5000     1134900    0        195743   100.3    0       
  TotalR: 24960/s QueueDepth(min=0, max=100)

[20:56:39] W: 5000/s TotalW:1159900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1159900    0        192399   100.3    0       
  1      5000     1159800    0        184067   100.3    100     
  2      5000     1159800    0        200216   100.3    100     
  3      5020     1159900    0        176764   100.3    0       
  4      5000     1159900    0        200545   100.3    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:56:44] W: 5000/s TotalW:1184900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1184800    0        196700   100.3    100     
  1      5000     1184800    0        188176   100.3    100     
  2      5000     1184800    0        204700   100.2    100     
  3      4980     1184800    0        180709   100.3    100     
  4      5000     1184900    0        205033   100.3    0       
  TotalR: 24960/s QueueDepth(min=0, max=100)

[20:56:49] W: 4560/s TotalW:1207700
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4460     1207100    0        199496   100.3    600     
  1      4420     1206900    0        190896   100.3    800     
  2      4460     1207100    0        207579   100.3    600     
  3      4440     1207000    0        183300   100.3    700     
  4      4460     1207200    0        207910   100.3    500     
  TotalR: 22238/s QueueDepth(min=500, max=800)

[20:56:54] W: 5441/s TotalW:1234900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5541     1234800    0        203478   100.4    100     
  1      5581     1234800    0        194683   100.4    100     
  2      5561     1234900    0        211762   100.3    0       
  3      5561     1234800    0        186951   100.4    100     
  4      5541     1234900    0        212090   100.3    0       
  TotalR: 27783/s QueueDepth(min=0, max=100)

[20:56:59] W: 5000/s TotalW:1259900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1259900    0        208245   100.4    0       
  1      5000     1259800    0        199183   100.4    100     
  2      5000     1259900    0        216779   100.3    0       
  3      5000     1259800    0        191301   100.4    100     
  4      4980     1259800    0        217088   100.3    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:57:04] W: 5000/s TotalW:1284900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1284900    0        212077   100.3    0       
  1      5020     1284900    0        202893   100.4    0       
  2      5000     1284900    0        220691   100.3    0       
  3      5020     1284900    0        194836   100.4    0       
  4      5000     1284800    0        221024   100.3    100     
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:57:09] W: 5000/s TotalW:1309900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1309900    0        216093   100.3    0       
  1      5000     1309900    0        206764   100.4    0       
  2      5000     1309900    0        224857   100.3    0       
  3      5000     1309900    0        198552   100.4    0       
  4      5000     1309800    0        225217   100.3    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:57:14] W: 5000/s TotalW:1334900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1334900    0        220887   100.3    0       
  1      5000     1334900    0        211282   100.4    0       
  2      5000     1334900    0        229877   100.3    0       
  3      5000     1334900    0        202918   100.4    0       
  4      5000     1334800    0        230249   100.3    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:57:19] W: 5000/s TotalW:1359900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1359800    0        225119   100.3    100     
  1      4980     1359800    0        215314   100.4    100     
  2      4980     1359800    0        234267   100.3    100     
  3      4980     1359800    0        206799   100.4    100     
  4      5000     1359800    0        234655   100.3    100     
  TotalR: 24920/s QueueDepth(min=100, max=100)

[20:57:24] W: 5000/s TotalW:1384900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1384900    0        228887   100.3    0       
  1      5020     1384900    0        218972   100.4    0       
  2      5020     1384900    0        238136   100.3    0       
  3      5020     1384900    0        210287   100.3    0       
  4      5000     1384800    0        238526   100.3    100     
  TotalR: 25080/s QueueDepth(min=0, max=100)

[20:57:29] W: 5000/s TotalW:1409900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1409800    0        233477   100.3    100     
  1      4980     1409800    0        223328   100.4    100     
  2      5000     1409900    0        242927   100.3    0       
  3      4980     1409800    0        214483   100.3    100     
  4      5000     1409800    0        243335   100.3    100     
  TotalR: 24940/s QueueDepth(min=0, max=100)

[20:57:34] W: 4980/s TotalW:1434800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4960     1434600    0        237821   100.3    200     
  1      4960     1434600    0        227493   100.4    200     
  2      4940     1434600    0        247462   100.3    200     
  3      4960     1434600    0        218458   100.3    200     
  4      4960     1434600    0        247875   100.3    200     
  TotalR: 24778/s QueueDepth(min=200, max=200)

[20:57:39] W: 5020/s TotalW:1459900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5040     1459800    0        240669   100.3    100     
  1      5060     1459900    0        230247   100.4    0       
  2      5040     1459800    0        250383   100.3    100     
  3      5060     1459900    0        221093   100.3    0       
  4      5040     1459800    0        250801   100.3    100     
  TotalR: 25242/s QueueDepth(min=0, max=100)

[20:57:44] W: 5000/s TotalW:1484900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1484900    0        244741   100.3    0       
  1      5000     1484900    0        234142   100.4    0       
  2      5000     1484800    0        254598   100.3    100     
  3      4980     1484800    0        224830   100.3    100     
  4      5020     1484900    0        255007   100.3    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

2025/10/05 20:57:48 [consumer g2 r3 p4] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/05 20:57:48 [consumer g4 r2 p3] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/05 20:57:48 [consumer g0 r3 p4] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/05 20:57:48 [consumer g1 r3 p4] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/05 20:57:48 [consumer g3 r1 p2] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/05 20:57:48 [consumer g1 r1 p2] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/05 20:57:48 [consumer g3 r3 p4] claim err: context deadline exceeded (group=g3 batch=200)

=== Summary ===
Total Writes: 1504600
Total Reads: 7523000
Total Updates: 0
Write Errors: 0
Read Errors: 0
Avg Write Throughput: 5015.33 rows/sec
Avg Read Throughput: 25076.67 rows/sec

Write Latencies (INSERT only):
  P50: 5.042175ms
  P95: 6.873087ms
  P99: 115.408895ms

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
Duration: 5m0s (300.0s)

-- Consumer Group 0 --
  Reads Completed:   1504600
  Read Errors:       0
  Updates Completed: 1504600
  Claim Errors:      1
  Empty Claims:      248525
  Avg Poll Size:     100.31
  Avg Throughput:    5015.33 reads/sec

  Read Latency (claim txn):
    P50: 3.913727ms
    P95: 4.943871ms
    P99: 53.018623ms

  End-to-End Latency (created_at → consumed):
    P50: 8.978431ms
    P95: 12.476415ms
    P99: 134.610943ms

-- Consumer Group 1 --
  Reads Completed:   1504600
  Read Errors:       0
  Updates Completed: 1504600
  Claim Errors:      2
  Empty Claims:      237716
  Avg Poll Size:     100.35
  Avg Throughput:    5015.33 reads/sec

  Read Latency (claim txn):
    P50: 3.946495ms
    P95: 4.943871ms
    P99: 52.756479ms

  End-to-End Latency (created_at → consumed):
    P50: 9.068543ms
    P95: 12.345343ms
    P99: 137.101311ms

-- Consumer Group 2 --
  Reads Completed:   1504600
  Read Errors:       0
  Updates Completed: 1504600
  Claim Errors:      1
  Empty Claims:      258581
  Avg Poll Size:     100.29
  Avg Throughput:    5015.33 reads/sec

  Read Latency (claim txn):
    P50: 3.827711ms
    P95: 4.898815ms
    P99: 52.756479ms

  End-to-End Latency (created_at → consumed):
    P50: 8.765439ms
    P95: 12.378111ms
    P99: 128.909311ms

-- Consumer Group 3 --
  Reads Completed:   1504600
  Read Errors:       0
  Updates Completed: 1504600
  Claim Errors:      2
  Empty Claims:      228292
  Avg Poll Size:     100.33
  Avg Throughput:    5015.33 reads/sec

  Read Latency (claim txn):
    P50: 4.184063ms
    P95: 5.111807ms
    P99: 52.035583ms

  End-to-End Latency (created_at → consumed):
    P50: 9.363455ms
    P95: 12.795903ms
    P99: 133.824511ms

-- Consumer Group 4 --
  Reads Completed:   1504600
  Read Errors:       0
  Updates Completed: 1504600
  Claim Errors:      1
  Empty Claims:      259012
  Avg Poll Size:     100.28
  Avg Throughput:    5015.33 reads/sec

  Read Latency (claim txn):
    P50: 3.657727ms
    P95: 4.833279ms
    P99: 53.149695ms

  End-to-End Latency (created_at → consumed):
    P50: 8.544255ms
    P95: 12.156927ms
    P99: 137.625599ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 7523000
  Total Read Errors:     0
  Total Updates:         7523000
  Total Claim Errors:    7
  Total Empty Claims:    1232126
  Avg Poll Size:         100.31
  Avg Throughput:        25076.67 reads/sec

  Read Latency (claim txn):
    P50: 3.921919ms
    P95: 4.964351ms
    P99: 52.789247ms

  End-to-End Latency (created_at → consumed):
    P50: 8.962047ms
    P95: 12.468223ms
    P99: 134.152191ms

2025/10/05 20:57:48 pubsub benchmark complete
```


# Run 2

```bash
ubuntu@ip-172-31-89-39:/tmp/postgres-queue-benchmarks$ ./pg_msg_bench   --host=$HOST   --port=5432   --db=benchmark   --user=postgres   --password=postgres --report=5s    --write-batch 100   --read-batch 200     --writers=10   --readers=4   --consumer-groups 5     --mode=pubsub --partitions 4 --duration=300s --throttle_writes=5000
2025/10/05 20:58:24 [pub info] successfully created 4 topicpartitions w/ counters + consumer_offsets
2025/10/05 20:58:25 [pub info] spawned readers
2025/10/05 20:58:25 [pub info] producers per partition [p1=3 p2=3 p3=2 p4=2]
[20:58:30] W: 6980/s TotalW:34900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      6960     34800      0        4643     103.3    100     
  1      6960     34800      0        4690     103.6    100     
  2      6960     34800      0        4956     102.7    100     
  3      6960     34800      0        4596     103.6    100     
  4      6980     34900      0        4796     103.9    0       
  TotalR: 34820/s QueueDepth(min=0, max=100)

[20:58:35] W: 5000/s TotalW:59900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     59800      0        8620     101.9    100     
  1      5020     59900      0        8720     102.0    0       
  2      5020     59900      0        9196     101.5    0       
  3      5000     59800      0        8540     102.0    100     
  4      5000     59900      0        8900     102.2    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[20:58:40] W: 5000/s TotalW:84900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     84800      0        13225    101.3    100     
  1      4980     84800      0        13495    101.4    100     
  2      5000     84900      0        14288    101.1    0       
  3      5000     84800      0        13247    101.4    100     
  4      4980     84800      0        13750    101.6    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[20:58:45] W: 5000/s TotalW:109900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     109800     0        17372    101.0    100     
  1      5000     109800     0        17719    101.1    100     
  2      5000     109900     0        18782    100.8    0       
  3      5000     109800     0        17404    101.1    100     
  4      5000     109800     0        18048    101.2    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:58:50] W: 5000/s TotalW:134900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     134800     0        21156    100.8    100     
  1      5000     134800     0        21498    100.9    100     
  2      5000     134900     0        22734    100.7    0       
  3      5000     134800     0        21104    100.9    100     
  4      5000     134800     0        21929    101.0    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:58:55] W: 5000/s TotalW:159900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     159800     0        25599    100.7    100     
  1      5000     159800     0        26082    100.8    100     
  2      5000     159900     0        27615    100.6    0       
  3      5000     159800     0        25611    100.8    100     
  4      5000     159800     0        26594    100.8    100     
  TotalR: 24999/s QueueDepth(min=0, max=100)

[20:59:00] W: 5000/s TotalW:184900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     184900     0        30106    100.6    0       
  1      5000     184800     0        30738    100.7    100     
  2      5000     184900     0        32581    100.5    0       
  3      5000     184800     0        30202    100.7    100     
  4      5020     184900     0        31314    100.7    0       
  TotalR: 25041/s QueueDepth(min=0, max=100)

[20:59:05] W: 5000/s TotalW:209900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     209900     0        33888    100.5    0       
  1      5000     209800     0        34544    100.6    100     
  2      5000     209900     0        36569    100.4    0       
  3      5000     209800     0        33909    100.6    100     
  4      4980     209800     0        35205    100.6    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:59:10] W: 5000/s TotalW:234900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     234800     0        37967    100.5    100     
  1      5000     234800     0        38688    100.5    100     
  2      5000     234900     0        40966    100.4    0       
  3      5000     234800     0        37996    100.5    100     
  4      5000     234800     0        39428    100.6    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:59:15] W: 5000/s TotalW:259900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     259900     0        42578    100.4    0       
  1      5000     259800     0        43445    100.5    100     
  2      5000     259900     0        46055    100.3    0       
  3      5000     259800     0        42706    100.5    100     
  4      5000     259800     0        44265    100.5    100     
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:59:20] W: 5000/s TotalW:284900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     284800     0        46665    100.4    100     
  1      5000     284800     0        47584    100.4    100     
  2      5000     284900     0        50437    100.3    0       
  3      5000     284800     0        46767    100.4    100     
  4      5000     284800     0        48488    100.5    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[20:59:25] W: 5000/s TotalW:309900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     309800     0        50458    100.4    100     
  1      5020     309900     0        51399    100.4    0       
  2      4980     309800     0        54426    100.3    100     
  3      5000     309800     0        50450    100.4    100     
  4      5000     309800     0        52366    100.4    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:59:30] W: 5000/s TotalW:334900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     334800     0        54726    100.3    100     
  1      4980     334800     0        55808    100.4    100     
  2      5020     334900     0        59117    100.3    0       
  3      5000     334800     0        54800    100.4    100     
  4      5020     334900     0        56821    100.4    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:59:35] W: 5000/s TotalW:359900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     359900     0        59095    100.3    0       
  1      5000     359800     0        60278    100.3    100     
  2      4980     359800     0        63895    100.3    100     
  3      5020     359900     0        59186    100.3    0       
  4      5000     359900     0        61378    100.4    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[20:59:40] W: 5000/s TotalW:384900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     384800     0        62849    100.3    100     
  1      5020     384900     0        64042    100.3    0       
  2      5000     384800     0        67848    100.2    100     
  3      4980     384800     0        62865    100.3    100     
  4      4980     384800     0        65229    100.3    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[20:59:45] W: 5000/s TotalW:409900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     409800     0        65327    100.7    100     
  1      4980     409800     0        66591    100.7    100     
  2      5000     409800     0        70544    100.6    100     
  3      5000     409800     0        65359    100.7    100     
  4      5020     409900     0        67804    100.8    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[20:59:50] W: 5000/s TotalW:434900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     434800     0        69909    100.6    100     
  1      5000     434800     0        71323    100.6    100     
  2      5000     434800     0        75609    100.6    100     
  3      5000     434800     0        70026    100.7    100     
  4      4980     434800     0        72628    100.7    100     
  TotalR: 24980/s QueueDepth(min=100, max=100)

[20:59:55] W: 5000/s TotalW:459900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     459800     0        74529    100.6    100     
  1      5020     459900     0        76076    100.6    0       
  2      5020     459900     0        80698    100.5    0       
  3      5000     459800     0        74691    100.7    100     
  4      5020     459900     0        77448    100.7    0       
  TotalR: 25060/s QueueDepth(min=0, max=100)

[21:00:00] W: 5000/s TotalW:484900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     484900     0        79081    100.6    0       
  1      4980     484800     0        80768    100.6    100     
  2      4980     484800     0        85723    100.5    100     
  3      5000     484800     0        79316    100.6    100     
  4      4980     484800     0        82224    100.6    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[21:00:05] W: 5000/s TotalW:509900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     509900     0        82819    100.5    0       
  1      5020     509900     0        84526    100.6    0       
  2      5020     509900     0        89663    100.5    0       
  3      5000     509800     0        82986    100.6    100     
  4      5020     509900     0        86067    100.6    0       
  TotalR: 25060/s QueueDepth(min=0, max=100)

[21:00:10] W: 5000/s TotalW:534900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     534800     0        86878    100.5    100     
  1      5000     534900     0        88615    100.5    0       
  2      4980     534800     0        94003    100.5    100     
  3      5000     534800     0        87002    100.6    100     
  4      5000     534900     0        90258    100.6    0       
  TotalR: 24960/s QueueDepth(min=0, max=100)

[21:00:15] W: 5000/s TotalW:559900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     559800     0        91492    100.5    100     
  1      4980     559800     0        93394    100.5    100     
  2      5000     559800     0        99101    100.4    100     
  3      5000     559800     0        91695    100.5    100     
  4      4980     559800     0        95091    100.6    100     
  TotalR: 24960/s QueueDepth(min=100, max=100)

[21:00:20] W: 5000/s TotalW:584900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     584800     0        95565    100.5    100     
  1      5000     584800     0        97531    100.5    100     
  2      5000     584800     0        103487   100.4    100     
  3      5000     584800     0        95755    100.5    100     
  4      5000     584800     0        99322    100.5    100     
  TotalR: 25000/s QueueDepth(min=100, max=100)

[21:00:25] W: 5000/s TotalW:609900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     609800     0        99328    100.4    100     
  1      5000     609800     0        101300   100.5    100     
  2      5000     609800     0        107467   100.4    100     
  3      5000     609800     0        99443    100.5    100     
  4      5020     609900     0        103187   100.5    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:00:30] W: 4979/s TotalW:634800
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4999     634800     0        103847   100.4    0       
  1      4999     634800     0        105949   100.4    0       
  2      4999     634800     0        112439   100.4    0       
  3      4999     634800     0        104025   100.5    0       
  4      4979     634800     0        107909   100.5    0       
  TotalR: 24977/s QueueDepth(min=0, max=0)

[21:00:35] W: 5021/s TotalW:659900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5001     659800     0        107531   100.4    100     
  1      5001     659800     0        109667   100.4    100     
  2      5021     659900     0        116403   100.4    0       
  3      5001     659800     0        107675   100.5    100     
  4      5021     659900     0        111710   100.5    0       
  TotalR: 25043/s QueueDepth(min=0, max=100)

[21:00:40] W: 5000/s TotalW:684900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     684800     0        111227   100.4    100     
  1      5000     684800     0        113370   100.4    100     
  2      4980     684800     0        120305   100.4    100     
  3      5000     684800     0        111278   100.4    100     
  4      4980     684800     0        115498   100.5    100     
  TotalR: 24959/s QueueDepth(min=100, max=100)

[21:00:45] W: 5000/s TotalW:709900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     709800     0        115354   100.4    100     
  1      5000     709800     0        117576   100.4    100     
  2      5000     709800     0        124762   100.4    100     
  3      5000     709800     0        115398   100.4    100     
  4      5000     709800     0        119788   100.4    100     
  TotalR: 25001/s QueueDepth(min=100, max=100)

[21:00:50] W: 5000/s TotalW:734900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     734900     0        118257   100.6    0       
  1      5020     734900     0        120577   100.7    0       
  2      5000     734800     0        127962   100.6    100     
  3      5000     734800     0        118342   100.7    100     
  4      5020     734900     0        122831   100.6    0       
  TotalR: 25060/s QueueDepth(min=0, max=100)

[21:00:55] W: 5000/s TotalW:759900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     759900     0        122214   100.6    0       
  1      5000     759900     0        124598   100.6    0       
  2      5000     759800     0        132175   100.5    100     
  3      5020     759900     0        122280   100.7    0       
  4      5000     759900     0        126924   100.6    0       
  TotalR: 25019/s QueueDepth(min=0, max=100)

[21:01:00] W: 5000/s TotalW:784900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     784900     0        126009   100.6    0       
  1      5000     784900     0        128403   100.6    0       
  2      5000     784800     0        136181   100.5    100     
  3      5000     784900     0        126006   100.6    0       
  4      5000     784900     0        130817   100.6    0       
  TotalR: 25001/s QueueDepth(min=0, max=100)

[21:01:05] W: 5000/s TotalW:809900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     809900     0        130593   100.6    0       
  1      4980     809800     0        133120   100.6    100     
  2      5020     809900     0        141228   100.5    0       
  3      5000     809900     0        130666   100.6    0       
  4      5000     809900     0        135615   100.6    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[21:01:10] W: 5000/s TotalW:834900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     834800     0        134950   100.6    100     
  1      5000     834800     0        137580   100.6    100     
  2      5000     834900     0        145992   100.5    0       
  3      4980     834800     0        135067   100.6    100     
  4      5000     834900     0        140164   100.6    0       
  TotalR: 24960/s QueueDepth(min=0, max=100)

[21:01:15] W: 5000/s TotalW:859900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     859800     0        138728   100.5    100     
  1      5020     859900     0        141368   100.6    0       
  2      4980     859800     0        149946   100.5    100     
  3      5000     859800     0        138763   100.6    100     
  4      5000     859900     0        144021   100.5    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[21:01:20] W: 5000/s TotalW:884900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     884800     0        142991   100.5    100     
  1      5000     884900     0        145721   100.5    0       
  2      5020     884900     0        154575   100.5    0       
  3      5000     884800     0        143044   100.6    100     
  4      5000     884900     0        148443   100.5    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:01:25] W: 5000/s TotalW:909900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     909800     0        147571   100.5    100     
  1      5000     909900     0        150464   100.5    0       
  2      5000     909900     0        159645   100.5    0       
  3      5020     909900     0        147700   100.6    0       
  4      5000     909900     0        153251   100.5    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:01:30] W: 5000/s TotalW:934900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     934800     0        151428   100.5    100     
  1      5000     934900     0        154363   100.5    0       
  2      4980     934800     0        163755   100.4    100     
  3      4980     934800     0        151515   100.5    100     
  4      5000     934900     0        157240   100.5    0       
  TotalR: 24960/s QueueDepth(min=0, max=100)

[21:01:35] W: 4900/s TotalW:959400
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4860     959100     0        154900   100.5    300     
  1      4820     959000     0        157865   100.5    400     
  2      4840     959000     0        167425   100.4    400     
  3      4860     959100     0        154934   100.5    300     
  4      4800     958900     0        160816   100.5    500     
  TotalR: 24180/s QueueDepth(min=300, max=500)

[21:01:40] W: 5100/s TotalW:984900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5140     984800     0        158401   100.5    100     
  1      5160     984800     0        161469   100.6    100     
  2      5180     984900     0        171268   100.5    0       
  3      5140     984800     0        158479   100.6    100     
  4      5180     984800     0        164472   100.5    100     
  TotalR: 25801/s QueueDepth(min=0, max=100)

[21:01:45] W: 5000/s TotalW:1009900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1009800    0        162530   100.5    100     
  1      5000     1009800    0        165651   100.5    100     
  2      4980     1009800    0        175691   100.5    100     
  3      5000     1009800    0        162584   100.6    100     
  4      5000     1009800    0        168744   100.5    100     
  TotalR: 24980/s QueueDepth(min=100, max=100)

[21:01:50] W: 5000/s TotalW:1034900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1034800    0        166317   100.5    100     
  1      5000     1034800    0        169444   100.5    100     
  2      5020     1034900    0        179678   100.4    0       
  3      5000     1034800    0        166287   100.5    100     
  4      5000     1034800    0        172638   100.5    100     
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:01:55] W: 5000/s TotalW:1059900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1059900    0        169553   100.6    0       
  1      5020     1059900    0        172774   100.6    0       
  2      5000     1059900    0        183221   100.6    0       
  3      5020     1059900    0        169540   100.7    0       
  4      5020     1059900    0        176011   100.6    0       
  TotalR: 25080/s QueueDepth(min=0, max=0)

[21:02:00] W: 5000/s TotalW:1084900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1084900    0        174138   100.6    0       
  1      5000     1084900    0        177486   100.6    0       
  2      5000     1084900    0        188270   100.5    0       
  3      5000     1084900    0        174179   100.7    0       
  4      4980     1084800    0        180799   100.6    100     
  TotalR: 24979/s QueueDepth(min=0, max=100)

[21:02:05] W: 5000/s TotalW:1109900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1109800    0        177908   100.6    100     
  1      5000     1109900    0        181271   100.6    0       
  2      5000     1109900    0        192245   100.5    0       
  3      4980     1109800    0        177849   100.6    100     
  4      5020     1109900    0        184644   100.6    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[21:02:10] W: 5000/s TotalW:1134900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1134900    0        181913   100.6    0       
  1      4980     1134800    0        185333   100.6    100     
  2      5000     1134900    0        196535   100.5    0       
  3      5020     1134900    0        181828   100.6    0       
  4      5000     1134900    0        188803   100.6    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:02:15] W: 5000/s TotalW:1159900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1159800    0        186515   100.6    100     
  1      5020     1159900    0        190089   100.6    0       
  2      5000     1159900    0        201610   100.5    0       
  3      5000     1159900    0        186507   100.6    0       
  4      4980     1159800    0        193616   100.6    100     
  TotalR: 24980/s QueueDepth(min=0, max=100)

[21:02:20] W: 5000/s TotalW:1184900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1184800    0        190649   100.6    100     
  1      4980     1184800    0        194300   100.6    100     
  2      5000     1184900    0        206081   100.5    0       
  3      5000     1184900    0        190646   100.6    0       
  4      5020     1184900    0        197901   100.6    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[21:02:25] W: 5000/s TotalW:1209900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1209900    0        194407   100.5    0       
  1      5000     1209800    0        198049   100.6    100     
  2      4980     1209800    0        210028   100.5    100     
  3      4980     1209800    0        194331   100.6    100     
  4      5000     1209900    0        201760   100.5    0       
  TotalR: 24980/s QueueDepth(min=0, max=100)

[21:02:30] W: 5000/s TotalW:1234900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1234900    0        198882   100.5    0       
  1      5000     1234800    0        202647   100.5    100     
  2      5020     1234900    0        214931   100.5    0       
  3      5020     1234900    0        198853   100.6    0       
  4      5000     1234900    0        206451   100.5    0       
  TotalR: 25040/s QueueDepth(min=0, max=100)

[21:02:35] W: 5000/s TotalW:1259900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1259800    0        203384   100.5    100     
  1      5020     1259900    0        207264   100.5    0       
  2      5000     1259900    0        219856   100.5    0       
  3      5000     1259900    0        203380   100.6    0       
  4      5000     1259900    0        211145   100.5    0       
  TotalR: 25000/s QueueDepth(min=0, max=100)

[21:02:40] W: 5000/s TotalW:1284900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1284800    0        206428   100.5    100     
  1      4980     1284800    0        210305   100.5    100     
  2      4980     1284800    0        223041   100.5    100     
  3      4980     1284800    0        206349   100.6    100     
  4      5000     1284900    0        214242   100.5    0       
  TotalR: 24940/s QueueDepth(min=0, max=100)

[21:02:45] W: 5000/s TotalW:1309900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1309900    0        210356   100.5    0       
  1      5000     1309800    0        214298   100.5    100     
  2      5000     1309800    0        227260   100.5    100     
  3      5000     1309800    0        210242   100.5    100     
  4      5000     1309900    0        218305   100.5    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:02:50] W: 5000/s TotalW:1334900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1334900    0        214863   100.5    0       
  1      5000     1334800    0        218914   100.5    100     
  2      5020     1334900    0        232188   100.4    0       
  3      5000     1334800    0        214790   100.5    100     
  4      5000     1334900    0        223003   100.5    0       
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:02:55] W: 5000/s TotalW:1359900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5000     1359900    0        217721   100.6    0       
  1      5000     1359800    0        221814   100.6    100     
  2      5000     1359900    0        235252   100.5    0       
  3      5020     1359900    0        217624   100.6    0       
  4      4980     1359800    0        225960   100.6    100     
  TotalR: 25000/s QueueDepth(min=0, max=100)

[21:03:00] W: 5000/s TotalW:1384900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1384800    0        221485   100.6    100     
  1      5000     1384800    0        225587   100.6    100     
  2      4980     1384800    0        239219   100.5    100     
  3      5000     1384900    0        221320   100.6    0       
  4      5000     1384800    0        229806   100.6    100     
  TotalR: 24960/s QueueDepth(min=0, max=100)

[21:03:05] W: 5000/s TotalW:1409900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1409900    0        226058   100.6    0       
  1      5000     1409800    0        230310   100.6    100     
  2      5000     1409800    0        244255   100.5    100     
  3      5000     1409900    0        225957   100.6    0       
  4      5000     1409800    0        234589   100.6    100     
  TotalR: 25020/s QueueDepth(min=0, max=100)

[21:03:10] W: 5000/s TotalW:1434900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1434800    0        230431   100.5    100     
  1      5000     1434800    0        234790   100.6    100     
  2      5000     1434800    0        249041   100.5    100     
  3      4980     1434800    0        230363   100.6    100     
  4      5000     1434800    0        239149   100.6    100     
  TotalR: 24960/s QueueDepth(min=100, max=100)

[21:03:15] W: 5000/s TotalW:1459900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      5020     1459900    0        234200   100.5    0       
  1      5020     1459900    0        238578   100.6    0       
  2      5020     1459900    0        253001   100.5    0       
  3      5020     1459900    0        234059   100.6    0       
  4      5000     1459800    0        243011   100.5    100     
  TotalR: 25080/s QueueDepth(min=0, max=100)

[21:03:20] W: 5000/s TotalW:1484900
  Group  R/s      TotR       Err      Empty    AvgPoll  QueueDepth
  0      4980     1484800    0        238416   100.5    100     
  1      4980     1484800    0        242876   100.5    100     
  2      4980     1484800    0        257565   100.5    100     
  3      5000     1484900    0        238296   100.6    0       
  4      5000     1484800    0        247395   100.5    100     
  TotalR: 24940/s QueueDepth(min=0, max=100)

2025/10/05 21:03:24 [consumer g2 r2 p3] claim err: context deadline exceeded (group=g2 batch=200)
2025/10/05 21:03:24 [consumer g0 r0 p1] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/05 21:03:24 [consumer g1 r0 p1] claim err: context deadline exceeded (group=g1 batch=200)
2025/10/05 21:03:24 [consumer g3 r1 p2] claim err: context deadline exceeded (group=g3 batch=200)
2025/10/05 21:03:24 [consumer g4 r1 p2] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/05 21:03:24 [consumer g0 r1 p2] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/05 21:03:24 [consumer g4 r2 p3] claim err: context deadline exceeded (group=g4 batch=200)
2025/10/05 21:03:24 [consumer g0 r2 p3] claim err: context deadline exceeded (group=g0 batch=200)
2025/10/05 21:03:24 [consumer g1 r1 p2] claim err: context deadline exceeded (group=g1 batch=200)

=== Summary ===
Total Writes: 1504200
Total Reads: 7521000
Total Updates: 0
Write Errors: 0
Read Errors: 0
Avg Write Throughput: 5014.00 rows/sec
Avg Read Throughput: 25070.00 rows/sec

Write Latencies (INSERT only):
  P50: 4.972543ms
  P95: 6.668287ms
  P99: 191.496191ms

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
Duration: 5m0s (300.0s)

-- Consumer Group 0 --
  Reads Completed:   1504200
  Read Errors:       0
  Updates Completed: 1504200
  Claim Errors:      3
  Empty Claims:      241975
  Avg Poll Size:     100.52
  Avg Throughput:    5014.00 reads/sec

  Read Latency (claim txn):
    P50: 3.885055ms
    P95: 4.784127ms
    P99: 59.604991ms

  End-to-End Latency (created_at → consumed):
    P50: 8.888319ms
    P95: 12.165119ms
    P99: 233.439231ms

-- Consumer Group 1 --
  Reads Completed:   1504200
  Read Errors:       0
  Updates Completed: 1504200
  Claim Errors:      2
  Empty Claims:      246536
  Avg Poll Size:     100.53
  Avg Throughput:    5014.00 reads/sec

  Read Latency (claim txn):
    P50: 3.891199ms
    P95: 4.886527ms
    P99: 60.555263ms

  End-to-End Latency (created_at → consumed):
    P50: 8.912895ms
    P95: 12.263423ms
    P99: 240.254975ms

-- Consumer Group 2 --
  Reads Completed:   1504200
  Read Errors:       0
  Updates Completed: 1504200
  Claim Errors:      1
  Empty Claims:      261485
  Avg Poll Size:     100.48
  Avg Throughput:    5014.00 reads/sec

  Read Latency (claim txn):
    P50: 3.676159ms
    P95: 4.644863ms
    P99: 57.671679ms

  End-to-End Latency (created_at → consumed):
    P50: 8.593407ms
    P95: 11.829247ms
    P99: 231.866367ms

-- Consumer Group 3 --
  Reads Completed:   1504200
  Read Errors:       0
  Updates Completed: 1504200
  Claim Errors:      1
  Empty Claims:      241920
  Avg Poll Size:     100.57
  Avg Throughput:    5014.00 reads/sec

  Read Latency (claim txn):
    P50: 4.007935ms
    P95: 5.013503ms
    P99: 59.572223ms

  End-to-End Latency (created_at → consumed):
    P50: 9.035775ms
    P95: 12.533759ms
    P99: 235.929599ms

-- Consumer Group 4 --
  Reads Completed:   1504200
  Read Errors:       0
  Updates Completed: 1504200
  Claim Errors:      2
  Empty Claims:      251129
  Avg Poll Size:     100.53
  Avg Throughput:    5014.00 reads/sec

  Read Latency (claim txn):
    P50: 3.758079ms
    P95: 4.689919ms
    P99: 61.079551ms

  End-to-End Latency (created_at → consumed):
    P50: 8.724479ms
    P95: 11.952127ms
    P99: 238.157823ms

=== Aggregate Across All Groups ===
  Total Reads Completed: 7521000
  Total Read Errors:     0
  Total Updates:         7521000
  Total Claim Errors:    9
  Total Empty Claims:    1243045
  Avg Poll Size:         100.53
  Avg Throughput:        25070.00 reads/sec

  Read Latency (claim txn):
    P50: 3.848191ms
    P95: 4.853759ms
    P99: 60.588031ms

  End-to-End Latency (created_at → consumed):
    P50: 8.822783ms
    P95: 12.230655ms
    P99: 236.584959ms

2025/10/05 21:03:24 pubsub benchmark complete
```
```