package pubsub

import (
	"fmt"
	"github.com/HdrHistogram/hdrhistogram-go"
	"time"
)

func (pm *PubSubMetrics) PrintSummary(dur time.Duration) {
	fmt.Printf("\n=== PubSub Summary ===\n")
	secs := dur.Seconds()

	fmt.Printf("Delivery Semantics: at-least-once (strict ordering)\n")
	fmt.Printf("Number of partitions: %d\n", pm.NumPartitions)
	fmt.Printf("Duration: %s (%.1fs)\n", dur, secs)

	snapshots := pm.Snapshot()

	// accumulators
	var totalReads, totalReadErr, totalUpdates, totalClaimErr, totalEmpty int64
	var totalPolls, totalRecs int64
	allReadHists := []*hdrhistogram.Histogram{}
	allE2EHists := []*hdrhistogram.Histogram{}

	for _, gs := range snapshots {
		fmt.Printf("\n-- Consumer Group %d --\n", gs.GroupID)
		fmt.Printf("  Reads Completed:   %d\n", gs.ReadsCompleted)
		fmt.Printf("  Read Errors:       %d\n", gs.ReadErrors)
		fmt.Printf("  Updates Completed: %d\n", gs.UpdatesCompleted)
		fmt.Printf("  Claim Errors:      %d\n", gs.ClaimErrors)
		fmt.Printf("  Empty Claims:      %d\n", gs.EmptyClaims)
		fmt.Printf("  Avg Poll Size:     %.2f\n", gs.AvgPollSize)

		if secs > 0 {
			fmt.Printf("  Avg Throughput:    %.2f reads/sec\n", float64(gs.ReadsCompleted)/secs)
		}

		fmt.Printf("\n  Read Latency (claim txn):\n")
		fmt.Printf("    P50: %v\n", time.Duration(gs.ReadP50))
		fmt.Printf("    P95: %v\n", time.Duration(gs.ReadP95))
		fmt.Printf("    P99: %v\n", time.Duration(gs.ReadP99))

		fmt.Printf("\n  End-to-End Latency (created_at → consumed):\n")
		fmt.Printf("    P50: %v\n", time.Duration(gs.E2EP50))
		fmt.Printf("    P95: %v\n", time.Duration(gs.E2EP95))
		fmt.Printf("    P99: %v\n", time.Duration(gs.E2EP99))

		// accumulate totals
		totalReads += gs.ReadsCompleted
		totalReadErr += gs.ReadErrors
		totalUpdates += gs.UpdatesCompleted
		totalClaimErr += gs.ClaimErrors
		totalEmpty += gs.EmptyClaims

		// grab the original hist slices for merging
		gm := pm.mustGroup(gs.GroupID)
		allReadHists = append(allReadHists, gm.ReaderReadHists...)
		allE2EHists = append(allE2EHists, gm.ReaderE2EHists...)

		totalPolls += gm.Polls.Load()
		totalRecs += gm.PolledRecords.Load()
	}

	// final merge
	readAgg := mergeHists(allReadHists)
	e2eAgg := mergeHists(allE2EHists)

	avgPollSize := 0.0
	if totalPolls > 0 {
		avgPollSize = float64(totalRecs) / float64(totalPolls)
	}

	fmt.Printf("\n=== Aggregate Across All Groups ===\n")
	fmt.Printf("  Total Reads Completed: %d\n", totalReads)
	fmt.Printf("  Total Read Errors:     %d\n", totalReadErr)
	fmt.Printf("  Total Updates:         %d\n", totalUpdates)
	fmt.Printf("  Total Claim Errors:    %d\n", totalClaimErr)
	fmt.Printf("  Total Empty Claims:    %d\n", totalEmpty)
	fmt.Printf("  Avg Poll Size:         %.2f\n", avgPollSize)

	if secs > 0 {
		fmt.Printf("  Avg Throughput:        %.2f reads/sec\n", float64(totalReads)/secs)
	}

	fmt.Printf("\n  Read Latency (claim txn):\n")
	fmt.Printf("    P50: %v\n", time.Duration(readAgg.ValueAtQuantile(50)))
	fmt.Printf("    P95: %v\n", time.Duration(readAgg.ValueAtQuantile(95)))
	fmt.Printf("    P99: %v\n", time.Duration(readAgg.ValueAtQuantile(99)))

	fmt.Printf("\n  End-to-End Latency (created_at → consumed):\n")
	fmt.Printf("    P50: %v\n", time.Duration(e2eAgg.ValueAtQuantile(50)))
	fmt.Printf("    P95: %v\n", time.Duration(e2eAgg.ValueAtQuantile(95)))
	fmt.Printf("    P99: %v\n", time.Duration(e2eAgg.ValueAtQuantile(99)))
	fmt.Println()
}
