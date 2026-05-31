package db

import (
	"log"

	"Project/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := "host=localhost user=postgres password=root dbname=goMonitor port=5432 sslmode=disable"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	err = database.AutoMigrate(&models.Target{})

	if err != nil {
		log.Fatal(err)
	}

	DB = database

	log.Println("Connected to PostgreSQL")
}
