package main

import (
	"context"
	"os/signal"
	"syscall"
	"time"

	"api-health-checker/internal/api"
	"api-health-checker/internal/checker"
	"api-health-checker/internal/repository"
	"api-health-checker/internal/scheduler"
	"api-health-checker/internal/worker"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	repo := repository.NewMemoryRepo()

	httpChecker := checker.NewHTTPChecker(5 * time.Second)

	pool := worker.NewPool(10, httpChecker, repo)
	pool.Start(ctx)

	urls := []string{
		"https://google.com",
		"https://github.com",
	}

	sched := scheduler.New(10*time.Second, urls, pool)
	sched.Start(ctx)

	server := api.NewServer(repo)
	server.Start()
}
