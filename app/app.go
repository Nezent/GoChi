package app

import (
	"net/http"

	"github.com/Nezent/GoChi/domain"
	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
)

func Start(){
	r := chi.NewRouter()
	ch := CustomerHandler{service: services.NewCustomerService(domain.NewCustomerRepositoryStub())}
	r.Get("/", ch.GetCustomers)
	http.ListenAndServe(":3000", r)
}