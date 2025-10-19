package common

import (
	"fmt"
	"github.com/HdrHistogram/hdrhistogram-go"
	"log"
	"time"
)

func mergeHists(hists []*hdrhistogram.Histogram) *hdrhistogram.Histogram {
	if len(hists) == 0 {
		return nil
	}
	// 1ns to 1m range, 3 sig figs
	out := hdrhistogram.New(1, int64(time.Minute), 3)

	var totalDropped int64
	for _, h := range hists {
		if h == nil {
			continue
		}
		dropped := out.Merge(h)
		totalDropped += dropped
	}
	if totalDropped > 0 {
		log.Printf("[merge] dropped %d values outside histogram range", totalDropped)
	}
	return out
}
func (br *BenchmarkRun) PrintSummary(dur time.Duration) {
	fmt.Printf("\n=== Summary ===\n")
	fmt.Printf("Total Writes: %d\n", br.Metrics.AggregateWritesCompleted.Load())
	fmt.Printf("Total Reads: %d\n", br.Metrics.AggregateReadsCompleted.Load())
	fmt.Printf("Total Updates: %d\n", br.Metrics.AggregateUpdatesCompleted.Load())
	fmt.Printf("Write Errors: %d\n", br.Metrics.AggregateWriteErrors.Load())
	fmt.Printf("Read Errors: %d\n", br.Metrics.AggregateReadErrors.Load())

	secs := dur.Seconds()
	if secs > 0 {
		fmt.Printf("Avg Write Throughput: %.2f rows/sec\n", float64(br.Metrics.AggregateWritesCompleted.Load())/secs)
		fmt.Printf("Avg Read Throughput: %.2f rows/sec\n", float64(br.Metrics.AggregateReadsCompleted.Load())/secs)
	}

	writeHist := mergeHists(br.Metrics.WriterHists)
	readHist := mergeHists(br.Metrics.ReaderReadHists)
	e2eHist := mergeHists(br.Metrics.ReaderE2EHists)

	if writeHist != nil {
		fmt.Printf("\nWrite Latencies (INSERT only):\n")
		fmt.Printf("  P50: %v\n", time.Duration(writeHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(writeHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(writeHist.ValueAtQuantile(99)))
	}
	if readHist != nil {
		fmt.Printf("\nRead Latencies (txn: SELECT+DELETE+INSERT in queue; txn: UPDATE+SELECT range in pub-sub kafka semantics):\n")
		fmt.Printf("  P50: %v\n", time.Duration(readHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(readHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(readHist.ValueAtQuantile(99)))
	}
	if e2eHist != nil {
		fmt.Printf("\nEnd-to-End Latencies (created_at → consumed):\n")
		fmt.Printf("  P50: %v\n", time.Duration(e2eHist.ValueAtQuantile(50)))
		fmt.Printf("  P95: %v\n", time.Duration(e2eHist.ValueAtQuantile(95)))
		fmt.Printf("  P99: %v\n", time.Duration(e2eHist.ValueAtQuantile(99)))
	}
	fmt.Println()
}
