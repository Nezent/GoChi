package app

import (
	"log"
	"net/http"

	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

func Start(){
	router := chi.NewRouter()
	dbClient := GetDBClient()
	ch := CustomerHandler{service: services.NewCustomerService(domain.NewCustomerRepositoryDB(dbClient))}
	ah := AccountHandler{service: services.NewAccountService(domain.NewAccountRepositoryDb(dbClient))}
	th := TransactionHandler{service: services.NewTransactionService(domain.NewTransactionRepositoryDb(dbClient))}
	router.Get("/customer", ch.GetCustomers)
	router.Get("/customer/{id:[0-9]+}",ch.GetCustomerByID)
	router.Post("/customer/{customer_id:[0-9]+}/account",ah.CreateAccount)
	router.Post("/transaction/{account_id:[0-9]+}",th.Transfer)
	log.Fatal(http.ListenAndServe(":3000", router))
}

func GetDBClient() *sqlx.DB {
	connStr := "postgres://postgres:anon@localhost/go_bank?sslmode=disable"
	client, err := sqlx.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return client
}