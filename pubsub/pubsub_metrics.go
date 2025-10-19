// pubsubmetrics.go
package pubsub

import (
	"fmt"
	"sync/atomic"
	"time"

	"github.com/HdrHistogram/hdrhistogram-go"
)

// per-group counters + latency hists (aggregated across that group's subscribers)
type GroupMetrics struct {
	// totals for this consumer group
	ReadsCompleted   atomic.Int64
	ReadErrors       atomic.Int64
	UpdatesCompleted atomic.Int64 // e.g., acks/commits/aux writes tied to reads
	ClaimErrors      atomic.Int64
	EmptyClaims      atomic.Int64 // optional: claimed nothing (no work). This tells you there are too many consumers
	PolledRecords    atomic.Int64
	Polls            atomic.Int64
	// one hist per subscriber in the group (single-writer each; merge at report time)
	ReaderReadHists []*hdrhistogram.Histogram // "read/claim" latency (ns)
	ReaderE2EHists  []*hdrhistogram.Histogram // end-to-end (now - created_at) (ns)
}

// pubsubmetrics tracks per-consumer-group metrics.
// keep using common.BaseMetrics for global aggregates; this is the per-group view.
type PubSubMetrics struct {
	Groups        []*GroupMetrics // index == group id (0..numGroups-1)
	NumPartitions int
}

// ctor: numGroups = distinct consumer groups; readersPerGroup = subscribers per group
func NewPubSubMetrics(numGroups, readersPerGroup int, numPartitions int) *PubSubMetrics {
	pm := &PubSubMetrics{
		Groups:        make([]*GroupMetrics, numGroups),
		NumPartitions: numPartitions,
	}
	for g := 0; g < numGroups; g++ {
		gm := &GroupMetrics{
			ReaderReadHists: make([]*hdrhistogram.Histogram, readersPerGroup),
			ReaderE2EHists:  make([]*hdrhistogram.Histogram, readersPerGroup),
		}
		for i := 0; i < readersPerGroup; i++ {
			gm.ReaderReadHists[i] = hdrhistogram.New(int64(1), int64((10 * time.Minute).Nanoseconds()), 3)
			gm.ReaderE2EHists[i] = hdrhistogram.New(int64(1), int64((10 * time.Minute).Nanoseconds()), 3)
		}
		pm.Groups[g] = gm
	}
	return pm
}

// helpers ---------------------------------------------------------------------

func (pm *PubSubMetrics) mustGroup(id int) *GroupMetrics {
	if id < 0 || id >= len(pm.Groups) {
		panic(fmt.Sprintf("group id out of range: %d (len=%d)", id, len(pm.Groups)))
	}
	return pm.Groups[id]
}

func mustReader(idx int, sliceLen int) int {
	if idx < 0 || idx >= sliceLen {
		panic(fmt.Sprintf("reader idx out of range: %d (len=%d)", idx, sliceLen))
	}
	return idx
}

// mutation API (call from hot path) -------------------------------------------

func (pm *PubSubMetrics) IncReadOK(groupID int) {
	pm.mustGroup(groupID).ReadsCompleted.Add(1)
}

func (pm *PubSubMetrics) IncReadErr(groupID int) {
	pm.mustGroup(groupID).ReadErrors.Add(1)
}

func (pm *PubSubMetrics) IncUpdateOK(groupID int) {
	pm.mustGroup(groupID).UpdatesCompleted.Add(1)
}

func (pm *PubSubMetrics) RecordReadLatencyNS(groupID, readerIdx int, ns int64) {
	gm := pm.mustGroup(groupID)
	_ = gm.ReaderReadHists[mustReader(readerIdx, len(gm.ReaderReadHists))].RecordValue(ns)
}

func (pm *PubSubMetrics) RecordE2ELatencyNS(groupID, readerIdx int, ns int64) {
	gm := pm.mustGroup(groupID)
	_ = gm.ReaderE2EHists[mustReader(readerIdx, len(gm.ReaderE2EHists))].RecordValue(ns)
}

// snapshot types --------------------------------------------------------------

type GroupSnapshot struct {
	GroupID int

	ReadsCompleted   int64
	ReadErrors       int64
	UpdatesCompleted int64
	ClaimErrors      int64
	EmptyClaims      int64
	AvgPollSize      float64

	// read/claim latency (ns)
	ReadP50 int64
	ReadP95 int64
	ReadP99 int64

	// end-to-end latency (ns)
	E2EP50 int64
	E2EP95 int64
	E2EP99 int64
}

// note: histograms are not concurrency-safe. take snapshots when writers are quiescent,
// or accept that percentiles are best-effort. you can also gate this behind your reporter tick.
func (pm *PubSubMetrics) Snapshot() []GroupSnapshot {
	out := make([]GroupSnapshot, 0, len(pm.Groups))
	for gid, gm := range pm.Groups {
		readAgg := mergeHists(gm.ReaderReadHists)
		e2eAgg := mergeHists(gm.ReaderE2EHists)
		polls := gm.Polls.Load()
		recs := gm.PolledRecords.Load()
		avgPollSize := 0.0
		if polls > 0 {
			avgPollSize = float64(recs) / float64(polls)
		}

		out = append(out, GroupSnapshot{
			GroupID: gid,

			ReadsCompleted:   gm.ReadsCompleted.Load(),
			ReadErrors:       gm.ReadErrors.Load(),
			UpdatesCompleted: gm.UpdatesCompleted.Load(),
			ClaimErrors:      gm.ClaimErrors.Load(),
			EmptyClaims:      gm.EmptyClaims.Load(),

			ReadP50: readAgg.ValueAtQuantile(50),
			ReadP95: readAgg.ValueAtQuantile(95),
			ReadP99: readAgg.ValueAtQuantile(99),

			AvgPollSize: avgPollSize,

			E2EP50: e2eAgg.ValueAtQuantile(50),
			E2EP95: e2eAgg.ValueAtQuantile(95),
			E2EP99: e2eAgg.ValueAtQuantile(99),
		})
	}
	return out
}

// utilities -------------------------------------------------------------------

func mergeHists(hs []*hdrhistogram.Histogram) *hdrhistogram.Histogram {
	if len(hs) == 0 {
		// degenerate: return an empty 1..1 hist
		return hdrhistogram.New(1, 1, 3)
	}
	agg := hdrhistogram.New(hs[0].LowestTrackableValue(), hs[0].HighestTrackableValue(), int(hs[0].SignificantFigures()))
	for _, h := range hs {
		_ = agg.Merge(h)
	}
	return agg
}

// short snapshot: safe for live reporters (no histogram access)
type GroupShortSnapshot struct {
	GroupID        int
	ReadsCompleted int64
	ReadErrors     int64
	Updates        int64
	ClaimErrors    int64
	EmptyClaims    int64
	AvgPollSize    float64
}

func (pm *PubSubMetrics) ShortSnapshot() []GroupShortSnapshot {
	out := make([]GroupShortSnapshot, 0, len(pm.Groups))
	for gid, gm := range pm.Groups {
		polls := gm.Polls.Load()
		recs := gm.PolledRecords.Load()
		avgPollSize := 0.0
		if polls > 0 {
			avgPollSize = float64(recs) / float64(polls)
		}
		out = append(out, GroupShortSnapshot{
			GroupID:        gid,
			ReadsCompleted: gm.ReadsCompleted.Load(),
			ReadErrors:     gm.ReadErrors.Load(),
			Updates:        gm.UpdatesCompleted.Load(),
			ClaimErrors:    gm.ClaimErrors.Load(),
			EmptyClaims:    gm.EmptyClaims.Load(),
			AvgPollSize:    avgPollSize,
		})
	}
	return out
}
