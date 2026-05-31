package main

import (
	"Project/internal/api"
	_ "Project/internal/dataBase"
	db "Project/internal/dataBase"

	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect()
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
