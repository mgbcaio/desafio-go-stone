package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mgbcaio/desafio-go-stone/pkg/common"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
)

func GetAllAccounts(w http.ResponseWriter, r *http.Request) {
	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//Add verificatoin to check if the Account its not empty
	json.NewEncoder(w).Encode(mocks.Accounts)
}

func GetAccountBalance(w http.ResponseWriter, r *http.Request) {
	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	vars := mux.Vars(r)
	id, _ := strconv.ParseInt(vars["id"], 10, 64)

	for _, account := range mocks.Accounts {
		if account.Id == id {
			w.Header().Add("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)

			json.NewEncoder(w).Encode(account.Balance)
			break
		}
	}
}
