package pubsub

import (
	"fmt"
	"sync"
	"time"
)

func (br *PubSubBenchmarkRun) Reporter(wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(br.config.ReportInterval)
	defer ticker.Stop()

	lastWrites := int64(0)
	lastReads := make([]int64, br.config.NumConsumerGroups)
	lastTime := time.Now()

	for {
		select {
		case <-br.Ctx.Done():
			return
		case <-ticker.C:
			now := time.Now()
			secs := now.Sub(lastTime).Seconds()

			// global writes
			writes := br.Metrics.AggregateWritesCompleted.Load()
			writeRate := float64(writes-lastWrites) / secs
			lastWrites = writes

			// per-group reads
			snaps := br.PubSubMetrics.ShortSnapshot()
			totalReads := int64(0)
			fmt.Printf("[%s] W: %.0f/s TotalW:%d\n", now.Format("15:04:05"), writeRate, writes)
			fmt.Printf("  %-6s %-8s %-10s %-8s %-8s %-8s %-8s\n",
				"Group", "R/s", "TotR", "Err", "Empty", "AvgPoll", "QueueDepth")
			queueDepths := make([]int64, 0, len(snaps))
			for _, gs := range snaps {
				deltaR := gs.ReadsCompleted - lastReads[gs.GroupID]
				readRate := float64(deltaR) / secs
				lastReads[gs.GroupID] = gs.ReadsCompleted
				totalReads += deltaR
				groupDepth := writes - gs.ReadsCompleted
				queueDepths = append(queueDepths, groupDepth)

				fmt.Printf("  %-6d %-8.0f %-10d %-8d %-8d %-8.1f %-8d\n",
					gs.GroupID,
					readRate,
					gs.ReadsCompleted,
					gs.ReadErrors,
					gs.EmptyClaims,
					gs.AvgPollSize,
					groupDepth,
				)
			}

			// compute min/max depth across groups
			minDepth, maxDepth := int64(1<<62), int64(-1<<62)
			for _, d := range queueDepths {
				if d < minDepth {
					minDepth = d
				}
				if d > maxDepth {
					maxDepth = d
				}
			}

			fmt.Printf("  TotalR: %.0f/s QueueDepth(min=%d, max=%d)\n\n",
				float64(totalReads)/secs,
				minDepth,
				maxDepth,
			)
			lastTime = now
		}
	}
}
