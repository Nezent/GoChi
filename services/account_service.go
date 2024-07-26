package services

import (
	"time"

	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/errs"
)

type AccountService interface {
	CreateAccount(dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError)
}

type DefaultAccountService struct {
	repository domain.AccountRepository
}

func (s DefaultAccountService) CreateAccount(req dto.NewAccountRequest) (*dto.NewAccountResponse, *errs.AppError){
	account := domain.CustomerAccount{
		AccountID: "",
		CustomerID: req.CustomerID,
		OpeningDate: time.Now().Format(time.RFC3339),
		AccountType: req.AccountType,
		Amount: req.Amount,
		Status: "1",
	}
	NewAccount, err := s.repository.SaveAccount(account)
	if err != nil {
		return nil, errs.NewUnexpectedError("Unexpected Error")
	}

	response := NewAccount.ToAccountResponseDTO()
	
	return &response, nil

}

func NewAccountService(repository domain.AccountRepository) DefaultAccountService {
	return DefaultAccountService{repository: repository}
}