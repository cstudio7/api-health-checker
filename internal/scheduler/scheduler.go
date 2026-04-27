package scheduler

import (
	"context"
	"time"

	"api-health-checker/internal/worker"
)

type Scheduler struct {
	interval time.Duration
	urls     []string
	pool     *worker.Pool
}

func New(interval time.Duration, urls []string, pool *worker.Pool) *Scheduler {
	return &Scheduler{
		interval: interval,
		urls:     urls,
		pool:     pool,
	}
}

func (s *Scheduler) Start(ctx context.Context) {
	ticker := time.NewTicker(s.interval)

	go func() {
		for {
			select {
			case <-ticker.C:
				for _, url := range s.urls {
					s.pool.Submit(worker.Job{URL: url})
				}
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}
