package main

import (
	"github.com/raihan-faza/lahhgrow/backend/models"
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
		panic(err)
	}
	db.AutoMigrate(&models.Account{}, &models.Video{}, &models.Course{})
}
