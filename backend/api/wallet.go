package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"gorm.io/gorm"
)

// nanti ku setup payment gatewaynya
func AddBalance(c *gin.Context, db *gorm.DB) {
	var fail_message = "fail to add balance"
	var req models.AddBalanceRequest
	var wallet models.Wallet
	user_id := c.Param("id")
	if user_id == "" {
		responses.BadRequest(c, fail_message)
		return
	}
	if db.First(&wallet, user_id).Error != nil {
		responses.BadRequest(c, fail_message)
		return
	}
	err := c.ShouldBindJSON(&req)
	if err != nil {
		panic(err)
	}
	adding_err := wallet.AddBalance(req.Balance)
	if adding_err != nil {
		panic(adding_err)
	}
	responses.GoodRequest(c, "balance added")
	return
}

// bisa ga ya pake cron buat nge cut per menit?
func CutBalance(c *gin.Context) {
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
