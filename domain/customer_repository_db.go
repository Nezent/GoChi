package domain

import (
	"database/sql"
	"log"

	"github.com/Nezent/GoChi/errs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type CustomerRepositoryDB struct {
	client *sqlx.DB
}

func (d CustomerRepositoryDB) FindAll() ([]Customer,*errs.AppError) {
	findCustomer := "SELECT * FROM customers"
	customers := make([]Customer,0)
	err := d.client.Select(&customers,findCustomer)
	if err != nil {
		log.Fatal(err)
		return nil, errs.NewNotFoundError("Customers Not Found")
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

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	connStr := "postgres://postgres:anon@localhost/go_bank?sslmode=disable"
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return CustomerRepositoryDB{client: client}
}