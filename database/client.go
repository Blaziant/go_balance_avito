package database

import (
	"log"

	"github.com/Blaziant/go_balance_avito/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	Instance.AutoMigrate(&models.User{}, &models.Favour{}, &models.ReserveAccount{}, &models.Order{})
	log.Println("Database Migration Completed!")
}
