package main

import (
	"Project/internal/api"
	_ "Project/internal/dataBase"
	db "Project/internal/dataBase"
	"Project/internal/models"
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
	target := models.Target{
		ID:       1,
		Name:     "google",
		URL:      "https://google.com",
		Interval: 30,
	}

	err := monitor.RunCheck(target)

	if err != nil {
		panic(err)
	}
	r := gin.Default()
	api.SetupRoutes(r)
	r.Run(":8080")
}
