package models

type MidtransRequest struct {
	UserId   int    `json:"user_id" binding:"required"`
	Amount   int64  `json:"amount" binding:"required"`
	ItemID   string `json:"item_id" binding:"required"`
	ItemName string `json:"item_name" binding:"required"`
}

type MidtransResponse struct {
	Token       string `json:"token"`
	RedirectUrl string `json:"redirect_url"`
}
