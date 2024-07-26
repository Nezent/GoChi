package domain

import (
	"database/sql"

	"github.com/Nezent/GoChi/errs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll(status bool) ([]Customer,*errs.AppError) {
	var findCustomer string
	var err error
	customers := make([]Customer,0)
	if !status {
		findCustomer = "SELECT * FROM customers"
		err = d.client.Select(&customers,findCustomer)
	} else {
		findCustomer = "SELECT * FROM customers WHERE status = $1"
		err = d.client.Select(&customers,findCustomer,status)
	}
	
	
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else{
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return customers,nil
}

func (d CustomerRepositoryDB) FindCustomerByID(id string) (*Customer, *errs.AppError) {
	findCustomerById := "SELECT * FROM customers WHERE id = $1"
	var customer Customer
	err := d.client.Get(&customer,findCustomerById,id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer Not Found")
		} else{
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	return &customer,nil
}

func NewCustomerRepositoryDB(dbClient *sqlx.DB) CustomerRepositoryDB {
	return CustomerRepositoryDB{client: dbClient}
}