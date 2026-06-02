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
