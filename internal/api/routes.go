package api

import (
	"Project/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	r.POST("/targets", handlers.CreateTarget)
	r.GET("/targets", handlers.GetTargets)
	r.DELETE("/targets/:id", handlers.DeleteTarget)
	r.GET("/targets/:id", handlers.GetTargetByID)
	r.PUT("/targets/:id", handlers.UpdateTarget)
	r.GET("/dashboard", handlers.Dashboard)
}
