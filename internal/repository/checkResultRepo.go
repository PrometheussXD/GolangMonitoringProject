package repository

import (
	"Project/internal/dataBase"
	"Project/internal/models"
)

func CreateCheckResult(result models.CheckResult) error {
	return db.DB.Create(&result).Error
}
