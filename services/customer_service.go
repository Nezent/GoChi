package services

import (
	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/errs"
)

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer,*errs.AppError)
	GetCustomerByID(string) (*domain.Customer,*errs.AppError)
}

type DefaultCustomerService struct {
	repository domain.CustomersRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer,*errs.AppError) {
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer,*errs.AppError) {
	return s.repository.FindCustomerByID(id)
}

func NewCustomerService(repository domain.CustomersRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}