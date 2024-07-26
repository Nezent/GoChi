package dto

type NewTransactionRequest struct {
	AccountID       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}