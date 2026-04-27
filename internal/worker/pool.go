package worker

import (
	"context"

	"api-health-checker/internal/checker"
	"api-health-checker/internal/models"
	"api-health-checker/internal/repository"
)

type Job struct {
	URL string
}

type Pool struct {
	workers int
	jobs    chan Job
	results chan models.HealthResult
	checker *checker.HTTPChecker
	repo    repository.Repository
}

func NewPool(workers int, checker *checker.HTTPChecker, repo repository.Repository) *Pool {
	return &Pool{
		workers: workers,
		jobs:    make(chan Job),
		results: make(chan models.HealthResult),
		checker: checker,
		repo:    repo,
	}
}

func (p *Pool) Start(ctx context.Context) {
	for i := 0; i < p.workers; i++ {
		go p.worker(ctx)
	}

	go p.collector()
}

func (p *Pool) worker(ctx context.Context) {
	for {
		select {
		case job := <-p.jobs:
			result := p.checker.Check(ctx, job.URL)
			p.results <- result
		case <-ctx.Done():
			return
		}
	}
}

func (p *Pool) collector() {
	for r := range p.results {
		p.repo.Save(r)
	}
}

func (p *Pool) Submit(job Job) {
	p.jobs <- job
}
