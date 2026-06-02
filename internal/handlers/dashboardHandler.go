package handlers

import (
	"Project/internal/models"
	"Project/internal/repository"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func Dashboard(c *gin.Context) {
	search := c.Query("search")
	targets, err := repository.GetAllTargets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var dashboardTargets []models.DashboardTarget
	upCount := 0
	downCount := 0

	for _, target := range targets {
		if search != "" {
			if !strings.Contains(strings.ToLower(target.Name), strings.ToLower(search)) {
				continue
			}
		}

		lastResult, _ := repository.GetLatestResult(target.ID)

		uptime := int(math.Round(repository.GetUptimePercentage(target.ID)))

		if lastResult.IsUp {
			upCount++
		} else {
			downCount++
		}

		dashboardTargets = append(
			dashboardTargets,
			models.DashboardTarget{
				ID:           target.ID,
				Name:         target.Name,
				URL:          target.URL,
				IsUp:         lastResult.IsUp,
				ResponseTime: lastResult.ResponseTime,
				LastCheck:    lastResult.CheckedAt,
				Uptime:       uptime,
			},
		)
	}

	sort.Slice(dashboardTargets, func(i, j int) bool {
		if dashboardTargets[i].IsUp == dashboardTargets[j].IsUp {
			return dashboardTargets[i].ID < dashboardTargets[j].ID
		}
		return !dashboardTargets[i].IsUp
	})

	totalCount := len(dashboardTargets)

	var totalUptime float64
	if totalCount > 0 {
		totalUptime = (float64(upCount) / float64(totalCount)) * 100
	}

	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	const pageSize = 10
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(dashboardTargets) {
		start = len(dashboardTargets)
	}
	if end > len(dashboardTargets) {
		end = len(dashboardTargets)
	}

	paginatedTargets := dashboardTargets[start:end]

	data := models.DashboardData{
		Targets:   paginatedTargets,
		Count:     totalCount,
		UpCount:   upCount,
		DownCount: downCount,
		Uptime:    totalUptime,
	}

	c.HTML(http.StatusOK, "dashboard.html", data)
}
