package services

import (
	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/errs"
)

type CustomerService interface {
	GetAllCustomer(string) ([]domain.Customer,*errs.AppError)
	GetCustomerByID(string) (*domain.Customer,*errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomersRepository
}

func (s DefaultCustomerService) GetAllCustomer(status string) ([]domain.Customer,*errs.AppError) {
	var activity bool
	if status == "active"{
		activity = true
	} else{
		activity = false
	}
	return s.repository.FindAll(activity)
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer,*errs.AppError) {
	return s.repository.FindCustomerByID(id)
}

func NewCustomerService(repository domain.CustomersRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}