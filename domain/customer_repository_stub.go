package domain

type CustomersRepositoryStub struct {
	customers []Customer
}

func (s CustomersRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomersRepositoryStub {
	customers := []Customer{
		{Id: "1001", Name: "Anon", City: "Bangladesh", DateOfBirth: "29-07-2002", Status: true},
		{Id: "1002", Name: "Munir", City: "Bangladesh", DateOfBirth: "10-02-2000", Status: false},
		{Id: "1003", Name: "Abdullah", City: "Bangladesh", DateOfBirth: "31-09-2006", Status: true},
	}
	return CustomersRepositoryStub{customers: customers}
}
