package handlers

import (
	"Project/internal/models"
	"Project/internal/repository"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {

	targets, err := repository.GetAllTargets()
	var dashboardTargets []models.DashboardTarget
	for _, target := range targets {

		lastResult, _ :=
			repository.GetLatestResult(target.ID)

		dashboardTargets =
			append(
				dashboardTargets,
				models.DashboardTarget{
					ID:           target.ID,
					Name:         target.Name,
					URL:          target.URL,
					IsUp:         lastResult.IsUp,
					ResponseTime: lastResult.ResponseTime,
					LastCheck:    lastResult.CheckedAt,
				},
			)
	}
	sort.Slice(
		dashboardTargets,
		func(i, j int) bool {

			if dashboardTargets[i].IsUp ==
				dashboardTargets[j].IsUp {

				return dashboardTargets[i].ID <
					dashboardTargets[j].ID
			}

			return !dashboardTargets[i].IsUp
		},
	)
	if err != nil {

		c.JSON(
			http.StatusInternalServerError,
			gin.H{"error": err.Error()},
		)

		return
	}
	upCount := 0
	downCount := 0

	for _, target := range dashboardTargets {

		if target.IsUp {
			upCount++
		} else {
			downCount++
		}
	}
	data := models.DashboardData{
		Targets: dashboardTargets,

		Count: len(dashboardTargets),

		UpCount: upCount,

		DownCount: downCount,
	}
	c.HTML(
		http.StatusOK,
		"dashboard.html",
		data,
	)
}
