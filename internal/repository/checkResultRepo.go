package repository

import (
	"Project/internal/dataBase"
	"Project/internal/models"
)

func CreateCheckResult(result models.CheckResult) error {
	return db.DB.Create(&result).Error
}

func GetLatestResult(
	targetID uint,
) (models.CheckResult, error) {

	var result models.CheckResult

	err := db.DB.
		Where("target_id = ?", targetID).
		Order("checked_at desc").
		First(&result).
		Error

	return result, err
}
func GetResultsByTargetID(
	targetID uint,
) ([]models.CheckResult, error) {

	var results []models.CheckResult

	err := db.DB.
		Where(
			"target_id = ?",
			targetID,
		).
		Order(
			"checked_at desc",
		).
		Limit(20).
		Find(&results).
		Error

	return results, err
}
func GetUptimePercentage(
	targetID uint,
) float64 {

	var total int64

	db.DB.
		Model(
			&models.CheckResult{},
		).
		Where(
			"target_id = ?",
			targetID,
		).
		Count(&total)

	if total == 0 {

		return 0
	}

	var up int64

	db.DB.
		Model(
			&models.CheckResult{},
		).
		Where(
			"target_id = ? AND is_up = true",
			targetID,
		).
		Count(&up)

	return float64(up) /
		float64(total) * 100
}
