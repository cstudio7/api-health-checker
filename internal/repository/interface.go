package repository

import "api-health-checker/internal/models"

type Repository interface {
	Save(result models.HealthResult)
	GetAll() []models.HealthResult
}
