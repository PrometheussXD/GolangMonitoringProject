package monitor

import (
	"time"

	"Project/internal/models"
)

func StartTargetWorker(
	target models.Target,
) {

	ticker := time.NewTicker(
		time.Duration(target.Interval) * time.Second,
	)

	for range ticker.C {

		RunCheck(target)
	}
}
