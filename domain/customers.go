package domain

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	DateOfBirth string `json:"dob"`
	Status      string `json:"status"`
}

type CustomersRepository interface {
	FindAll() ([]Customer, error)
}