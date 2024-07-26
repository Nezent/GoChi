package app

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Nezent/GoChi/dto"
	"github.com/Nezent/GoChi/services"
	"github.com/go-chi/chi/v5"
)

type TransactionHandler struct {
	service services.TransactionService
}

func (h TransactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	account_id := chi.URLParam(r,"account_id")
	var request dto.NewTransactionRequest

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Fatal(err)
	}
	request.AccountID = account_id
	transaction, appError := h.service.Transfer(request)

	if appError != nil {
		log.Fatal(appError)
	} else {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(transaction)
	}	
}