package services

import "github.com/Nezent/GoChi/domain"

type CustomerService interface {
	GetAllCustomer() ([]domain.Customer,error)
	GetCustomerByID(string) (*domain.Customer,error)
}

type DefaultCustomerService struct {
	repository domain.CustomersRepository
}

func (s DefaultCustomerService) GetAllCustomer() ([]domain.Customer,error) {
	return s.repository.FindAll()
}

func (s DefaultCustomerService) GetCustomerByID(id string) (*domain.Customer,error) {
	return s.repository.FindCustomerByID(id)
}

func NewCustomerService(repository domain.CustomersRepository) DefaultCustomerService {
	return DefaultCustomerService{
		repository: repository,
	}
}