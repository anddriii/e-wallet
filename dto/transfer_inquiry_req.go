package dto

type TransferInquiryReq struct {
	AccountNumber string `json: "account_number"`
	Amount        string `json: "amount"`
}
