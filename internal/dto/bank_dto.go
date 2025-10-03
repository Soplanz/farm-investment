package dto

type CreateBankRequest struct {
	Name    string `json:"name" binding:"required"`
	Swift   string `json:"swift" binding:"required"`
	Country string `json:"country" binding:"required"`
}
