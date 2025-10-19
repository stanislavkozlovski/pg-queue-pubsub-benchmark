package common

import (
	"sync/atomic"

	"github.com/HdrHistogram/hdrhistogram-go"
)

type BaseMetrics struct {
	AggregateWritesCompleted  atomic.Int64
	AggregateReadsCompleted   atomic.Int64
	AggregateUpdatesCompleted atomic.Int64
	AggregateWriteErrors      atomic.Int64
	AggregateReadErrors       atomic.Int64

	// Writer captures the latency to execute the INSERT write query
	WriterHists []*hdrhistogram.Histogram
	// Read Select captures the latency to execute the read query combination
	// for queues:  SELECT+DELETE+INSERT
	ReaderReadHists []*hdrhistogram.Histogram
	// E2E captures the end-to-end latency, `read_time - write_time`
	ReaderE2EHists []*hdrhistogram.Histogram
}

func NewMetrics(writers, readers int) *BaseMetrics {
	m := &BaseMetrics{
		WriterHists:     make([]*hdrhistogram.Histogram, writers),
		ReaderReadHists: make([]*hdrhistogram.Histogram, readers),
		ReaderE2EHists:  make([]*hdrhistogram.Histogram, readers),
	}
	for i := range m.WriterHists {
		m.WriterHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	for i := range m.ReaderReadHists {
		m.ReaderReadHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	for i := range m.ReaderE2EHists {
		m.ReaderE2EHists[i] = hdrhistogram.New(1, int64(600e9), 3)
	}
	return m
}
