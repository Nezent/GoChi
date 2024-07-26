package domain

import (
	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/errs"
)

type Transaction struct {
	TransactionID   string  `db:"transaction_id"`
	AccountID       string  `db:"account_id"`
	TransactionType string  `db:"transaction_type"`
	Amount          float64 `db:"amount"`
	TransactionDate string  `db:"transaction_date"`
}

func (t Transaction) ToTransactionResponseDTO() dto.NewTransactionReponse {
	return dto.NewTransactionReponse{TransactionID: t.TransactionID,Amount: t.Amount}
}

type TransactionRepository interface {
	MakeTransaction(Transaction) (*Transaction, *errs.AppError)
}