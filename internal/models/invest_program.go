package models

type InvestProgram struct {
	ID            int32   `json:"id"`
	CompanyID     int32   `json:"company_id"`
	Name          string  `json:"name"`
	Description   string  `json:"description"`
	Duration      int32   `json:"duration"` // in months
	InitialAmount float64 `json:"initial_amount"`
	CreatedAt     string  `json:"created_at"`
	UpdatedAt     string  `json:"updated_at"`
}
