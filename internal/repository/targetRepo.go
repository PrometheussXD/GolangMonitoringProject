package repository

import "Project/internal/models"

var targets []models.Target
var nextID = 1

func CreateTarget(t models.Target) models.Target {
	t.Id = nextID
	nextID++

	targets = append(targets, t)
	return t
}

func GetTargets() []models.Target {
	return targets
}
