package domain

type Customer struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	DateOfBirth string `json:"date_of_birth" db:"dob"`
	Status      bool   `json:"status"`
}

type CustomersRepository interface {
	FindAll() ([]Customer, error)
	FindCustomerByID(string) (*Customer, error)
}