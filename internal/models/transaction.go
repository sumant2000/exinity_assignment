package models

type Transaction struct {
	ID          string  `json:"id"`
	AccountID   string  `json:"account_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	GatewayName string  `json:"gateway_name"`
}
