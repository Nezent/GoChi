package app

import (
	"encoding/json"
	"net/http"

	"github.com/Nezent/GoChi/services"
)

type CustomerHandler struct {
	service services.CustomerService
}

func (ch *CustomerHandler) GetCustomers(w http.ResponseWriter, r *http.Request) {
	customers,_ := ch.service.GetAllCustomer()

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}