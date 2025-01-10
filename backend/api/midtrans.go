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

/*
rencananya ini buat generate pembayaran topup
nanti kan customer tinggal masukin nominal yang mau di topup, terus muncul page buat bayar kan
berarti user perlu ngasih kita beberapa hal:
  - jumlah uang yang mau di topup (ini lewat input form misalnya)
  - wallet id nya (ini bisa kita query pake user id yang ada di jwt token)

nah itu doang yang kita butuh di database.
tapi, ini banyak perintilannya dari midtrans.
*/
func GenerateSnapReq(cust_detail *midtrans.CustomerDetails, transaction_detail midtrans.TransactionDetails) *snap.Request {
	snapReq := &snap.Request{
		TransactionDetails: transaction_detail,
		CustomerDetail:     cust_detail,
		EnabledPayments:    snap.AllSnapPaymentType,
	}
	return snapReq
}

// ini api buat topupnya
func CreateTransaction(c *gin.Context, gross_amount int64, db *gorm.DB) *snap.Response {
	//bind dulu data yang mau diambil, sisanya kita query di backend (jumlah uang, user id)
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
