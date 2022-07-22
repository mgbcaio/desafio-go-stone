package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
)

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Add verificatoin to check if the Account its not empty
	json.NewEncoder(w).Encode(mocks.Accounts)
}
