package app

import (
	"log"
	"net/http"

	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
)

func Start(){
	router := chi.NewRouter()
	ch := CustomerHandler{service: services.NewCustomerService(domain.NewCustomerRepositoryDB())}
	router.Get("/customer", ch.GetCustomers)
	router.Get("/customer/{id:[0-9]+}",ch.GetCustomerByID)
	log.Fatal(http.ListenAndServe(":3000", router))
}