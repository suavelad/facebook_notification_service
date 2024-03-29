package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@notify-go-db:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate((&models.AdAccount{}, &models.FacebookAdBalanceLog{}, &models.NotificationSentLog{}))
	DB = db
	log.Println("🚀 🚀 🚀  Connected Successfully to the Database   🚀 🚀 🚀 ")
}
