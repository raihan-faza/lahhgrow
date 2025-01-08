package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/raihan-faza/lahhgrow/backend/utils"
)

var s = utils.Snapc

func create_transaction() snap.ResponseWithMap {
	res, err := snap.CreateTransactionWithMap(example.SnapParamWithMap())
	if err != nil {
		return nil
	}
	return res
}

// rencananya ini buat generate pembayaran topup
func GenerateSnapReq(cust_detail *midtrans.CustomerDetails, gross_amt int64, items *[]midtrans.ItemDetails) *snap.Request {
	transDetail := midtrans.TransactionDetails{
		OrderID:  "lahhgrow-topup-wallet" + example.Random(),
		GrossAmt: gross_amt,
	}
	// Initiate Snap Request
	snapReq := &snap.Request{
		TransactionDetails: transDetail,
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail:  cust_detail,
		EnabledPayments: snap.AllSnapPaymentType,
		Items:           items,
	}
	return snapReq
}

// ini api buat topupnya
func CreateTransaction(c *gin.Context) {
	//tulis kode buat bikin input funtsi ^&&&&&&@) {U!)@JDI@!JD@[]}
	resp, err := s.CreateTransaction(GenerateSnapReq())
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, resp)
}
