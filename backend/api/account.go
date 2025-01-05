package api

import (
	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"github.com/raihan-faza/lahhgrow/backend/utils"
	"gorm.io/gorm"
)

func CreateAccount(c *gin.Context, db *gorm.DB) {
	var account models.Account
	var fail_message = "failed to create account"
	err := c.ShouldBindJSON(&account)
	if err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	hashed_password, hash_err := utils.CreateHash(account.Password)
	if hash_err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	account.Password = hashed_password
	db_error := db.Create(&account).Error
	if db_error != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	responses.GoodRequest(c, "account created")
	return
}

func DeleteAccount(c *gin.Context, db *gorm.DB) {
	var fail_message = "failed to delete account"
	course_id := c.Param("id")
	if course_id == "" {
		responses.BadRequest(c, fail_message)
		return
	}
	db_err := db.Delete(&models.Course{}, course_id).Error
	if db_err != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	responses.GoodRequest(c, "account deleted succesfully")
	return
}
