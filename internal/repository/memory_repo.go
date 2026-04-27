package repository

import (
	"sync"

	"api-health-checker/internal/models"
)

type MemoryRepo struct {
	mu      sync.RWMutex
	results []models.HealthResult
}

func NewMemoryRepo() *MemoryRepo {
	return &MemoryRepo{
		results: make([]models.HealthResult, 0),
	}
}

func (r *MemoryRepo) Save(result models.HealthResult) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.results = append(r.results, result)
}

func (r *MemoryRepo) GetAll() []models.HealthResult {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return append([]models.HealthResult(nil), r.results...)
}
