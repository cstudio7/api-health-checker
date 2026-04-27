package checker

import (
	"context"
	"net/http"
	"time"

	"api-health-checker/internal/models"
)

type HTTPChecker struct {
	client *http.Client
}

func NewHTTPChecker(timeout time.Duration) *HTTPChecker {
	return &HTTPChecker{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *HTTPChecker) Check(ctx context.Context, url string) models.HealthResult {
	start := time.Now()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return models.HealthResult{
			URL:       url,
			Error:     err.Error(),
			Timestamp: time.Now(),
		}
	}

	resp, err := c.client.Do(req)

	result := models.HealthResult{
		URL:       url,
		Latency:   time.Since(start),
		Timestamp: time.Now(),
	}

	if err != nil {
		result.Error = err.Error()
		return result
	}

	defer resp.Body.Close()
	result.StatusCode = resp.StatusCode

	return result
}
