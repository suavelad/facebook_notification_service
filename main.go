package main

import (
	"net/http"

	

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/suavelad/notification_service_go/models"
	"github.com/suavelad/notification_service_go/initializer"
	"github.com/suavelad/notification_service_go/service"
)


var db *gorm.DB

func initDB() {
	database, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	database.AutoMigrate(&models.AdAccount{}, &models.FacebookAdBalanceLog{}, &models.NotificationSentLog{)
	db = database
}



func init() {
	initializer.LoadEnvVariables()

	initDB()
	defer db.Close()

}


func main() {

    // Create a new cron scheduler
    c := cron.New()

    // Define the cron schedule for running sendBalanceAlert
    _, err := c.AddFunc("0 6,23 * * *", service.sendBalanceAlert)

    if err != nil {
        fmt.Println("Error adding cron job:", err)
        return
    }

    c.Start()

    for {
        time.Sleep(1 * time.Minute)
    }
}
