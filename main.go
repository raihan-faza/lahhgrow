package main

import (
	"log"

	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/router"
	"github.com/raihan-faza/lahhgrow/backend/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	utils.LoadEnv()
}

func main() {
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&models.Account{}, &models.Video{}, &models.Course{})
	r := router.MainRouter(db)
	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
