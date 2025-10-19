package common

import (
	"context"
	"database/sql"
	"golang.org/x/time/rate"
	"time"
)

type BaselineConfig struct {
	Writers        int
	Readers        int
	Duration       time.Duration
	PayloadSize    int
	ReportInterval time.Duration
}

type BenchmarkRun struct {
	Db           *sql.DB
	Metrics      *BaseMetrics
	Ctx          context.Context
	WriteLimiter *rate.Limiter // global limiter
}

func NewBenchmarkRun(
	db *sql.DB,
	metrics *BaseMetrics,
	ctx context.Context,
	writeLimiter *rate.Limiter,
) *BenchmarkRun {
	return &BenchmarkRun{
		Db:           db,
		Metrics:      metrics,
		Ctx:          ctx,
		WriteLimiter: writeLimiter,
	}
}
