package main

import (
	"Project/internal/api"
	db "Project/internal/dataBase"
	"Project/internal/monitor"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	go monitor.StartScheduler()
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
