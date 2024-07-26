package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
)

type AccountHandler struct {
	service services.AccountService
}

func (h AccountHandler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	customer_id := chi.URLParam(r,"customer_id")
	var request dto.NewAccountRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	request.CustomerID = customer_id
	account, appError := h.service.CreateAccount(request)

	if appError != nil {
		log.Fatal(appError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(account)
	}


}