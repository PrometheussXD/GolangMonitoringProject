package main

import (
	"Project/internal/api"
	_ "Project/internal/dataBase"
	db "Project/internal/dataBase"
	"Project/internal/monitor"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	status, responseTime, isUp :=
		monitor.CheckURL("https://google.com")
	fmt.Println(status)
	fmt.Println(responseTime)
	fmt.Println(isUp)
	db.Connect()
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
