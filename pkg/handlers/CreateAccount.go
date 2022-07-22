package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
	"github.com/mgbcaio/desafio-go-stone/pkg/models"
)

func CreateAccount(w http.ResponseWriter, r *http.Request) {
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
