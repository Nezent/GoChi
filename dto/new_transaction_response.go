package dto

type NewTransactionReponse struct {
	TransactionID string  `json:"transaction_id"`
	Amount        float64 `json:"amount"`
}