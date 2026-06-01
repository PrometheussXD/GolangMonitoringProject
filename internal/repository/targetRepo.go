package repository

import (
	"Project/internal/dataBase"
	"Project/internal/models"
)

func CreateTarget(target models.Target) error {
	return db.DB.Create(&target).Error
}

func GetTargets() ([]models.Target, error) {

	var targets []models.Target

	err := db.DB.Find(&targets).Error

	return targets, err
}
func DeleteTarget(id uint) error {
	return db.DB.Delete(&models.Target{}, id).Error
}
