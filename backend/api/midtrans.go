package api

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/raihan-faza/lahhgrow/backend/models"
	"github.com/raihan-faza/lahhgrow/backend/responses"
	"github.com/raihan-faza/lahhgrow/backend/utils"
	"gorm.io/gorm"
)

var s = utils.Snapc

func create_transaction() snap.ResponseWithMap {
	res, err := snap.CreateTransactionWithMap(example.SnapParamWithMap())
	if err != nil {
		return nil
	}
	return res
}

func GenerateSnapReq(cust_detail *midtrans.CustomerDetails, transaction_detail midtrans.TransactionDetails) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: transaction_detail,
		CustomerDetail:     cust_detail,
		EnabledPayments:    snap.AllSnapPaymentType,
	}
	return snapReq
}

func CreateTransaction(c *gin.Context, gross_amount int64, db *gorm.DB) *snap.Response {
	var account models.Account
	user_token, err := DecodeJwt(GetJwtFromHeader(c))
	if err != nil {
		panic(err)
	}
	if db.First(&account, user_token["user_id"]).Error != nil {
		responses.BadRequest(c, "user not found")
	}
	transaction_detail := midtrans.TransactionDetails{
		OrderID:  "LAHHGROW-WALLET-TOPUP-" + example.Random(),
		GrossAmt: gross_amount,
	}
	user_addr := midtrans.CustomerAddress{
		FName:       account.FirstName,
		LName:       account.LastName,
		Phone:       account.Phone,
		City:        account.City,
		Postcode:    account.Postcode,
		CountryCode: account.CountryCode,
	}
	customer_detail := midtrans.CustomerDetails{
		FName:    account.FirstName,
		LName:    account.LastName,
		Email:    account.Email,
		Phone:    account.Phone,
		BillAddr: &user_addr,
		ShipAddr: &user_addr,
	}
	resp, resp_err := s.CreateTransaction(GenerateSnapReq(&customer_detail, transaction_detail))
	if resp_err != nil {
		panic(resp_err)
	}
	return resp
}
