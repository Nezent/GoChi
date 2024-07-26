package domain

import (
	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/errs"
)

type CustomerAccount struct {
	AccountID   string  `json:"account_id"`
	CustomerID  string  `json:"customer_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

func (account CustomerAccount) ToAccountResponseDTO() dto.NewAccountResponse {
	return dto.NewAccountResponse{AccountID: account.AccountID}
}

type AccountRepository interface {
	SaveAccount(CustomerAccount) (*CustomerAccount, *errs.AppError)
}