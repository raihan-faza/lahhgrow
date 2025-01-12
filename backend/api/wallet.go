package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"gorm.io/gorm"
)

func AddBalance(c *gin.Context, db *gorm.DB) {
	var account models.Account
	user_token, err := DecodeJwt(GetJwtFromHeader(c))
	if err != nil {
		panic(err)
	}
	if db.First(&account, user_token["user_id"]).Error != nil {
		responses.BadRequest(c, "user not found")
	}
	return
}

func StartCutBalance(c *gin.Context, db *gorm.DB) {
	var cron models.WalletCronManager
	user_token, err := DecodeJwt(GetJwtFromHeader(c))
	if err != nil {
		panic(err)
	}
	if db.Preload("Wallet", "account_id = ?", user_token["user_id"]).Find(&cron).Error != nil {
		responses.BadRequest(c, "wallet not found")
		return
	}
	cron.StartCutBalance()
	return
}

func StopCutBalance(c *gin.Context, db *gorm.DB) {
	var cron models.WalletCronManager
	user_token, err := DecodeJwt(GetJwtFromHeader(c))
	if err != nil {
		panic(err)
	}
	if db.Preload("Wallet", "account_id = ?", user_token["user_id"]).Find(&cron).Error != nil {
		responses.BadRequest(c, "wallet not found")
		return
	}
	cron.StopCutBalance()
	return
}

func GetBalance(c *gin.Context, db *gorm.DB) {
	var wallet models.Wallet
	var fail_message = "fail to get wallet"
	wallet_id := c.Param("id")
	if db.First(&wallet, wallet_id).Error != nil {
		responses.BadRequest(c, fail_message)
	}
	c.JSON(http.StatusOK, wallet)
	return
}
