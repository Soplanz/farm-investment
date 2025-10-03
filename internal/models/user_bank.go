package models

type UserBank struct {
	BankID        int32  `json:"bank_id"`
	UserID        int32  `json:"user_id"`
	BankName      string `json:"bank_name"`
	AccountNumber string `json:"account_number"`
}
