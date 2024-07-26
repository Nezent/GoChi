package services

import (
	"time"

	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/errs"
)

type TransactionService interface {
	Transfer(dto.NewTransactionRequest) (*dto.NewTransactionReponse, *errs.AppError)
}

type DefaultTransactionService struct {
	repository domain.TransactionRepository
}


func (t DefaultTransactionService) Transfer(request dto.NewTransactionRequest) (*dto.NewTransactionReponse, *errs.AppError) {
	transferModel := domain.Transaction{
		AccountID: request.AccountID,
		TransactionType: request.TransactionType,
		Amount: request.Amount,
		TransactionDate: time.Now().Format(time.RFC3339),
	}

	NewTransaction, err := t.repository.MakeTransaction(transferModel)

	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}
	response := NewTransaction.ToTransactionResponseDTO()

	return &response, nil
}

func NewTransactionService(repository domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		repository: repository,
	}
}