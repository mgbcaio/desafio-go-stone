package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mgbcaio/desafio-go-stone/pkg/common"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
	"github.com/mgbcaio/desafio-go-stone/pkg/models"
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

func CreateAccount(w http.ResponseWriter, r *http.Request) {
	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var account models.Account
	json.Unmarshal(body, &account)

	account.Id = rand.Int63n(100)
	account.CreatedAt = time.Now()
	account.Balance = 0

	mocks.Accounts = append(mocks.Accounts, account)
	fmt.Print(mocks.Accounts)

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Account Created")
}