package domain

import "github.com/Nezent/GoChi/errs"

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	DateOfBirth string `json:"date_of_birth" db:"dob"`
	Status      bool   `json:"status"`
}

type CustomersRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindCustomerByID(string) (*Customer, *errs.AppError)
}