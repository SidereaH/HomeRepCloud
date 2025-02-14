package database

import (
	"HomeRepCloud/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}

	log.Println("Connected to Database!")
}

func Migrate() {
	Instance.AutoMigrate(&models.Image{})

	log.Println("Database Migration Completed!")
}
