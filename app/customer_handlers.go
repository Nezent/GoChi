package app

import (
	"encoding/json"
	"net/http"

	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
)

type CustomerHandler struct {
	service services.CustomerService
}

func (ch *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers,_ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (ch *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r,"id")
	customer,err := ch.service.GetCustomerByID(id)
	if(err != nil) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(err.Code)
		json.NewEncoder(w).Encode(err.AsMessage())
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(customer)
	}
}