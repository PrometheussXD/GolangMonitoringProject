package monitor

import (
	"time"

	"Project/internal/models"
	"Project/internal/repository"
)

func RunCheck(target models.Target) error {

	statusCode, responseTime, isUp :=
		CheckURL(target.URL)

	result := models.CheckResult{
		TargetID:     target.ID,
		StatusCode:   statusCode,
		ResponseTime: responseTime,
		IsUp:         isUp,
		CheckedAt:    time.Now(),
	}

	return repository.CreateCheckResult(result)
}
