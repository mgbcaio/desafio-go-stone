package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mgbcaio/desafio-go-stone/pkg/common"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
	"github.com/mgbcaio/desafio-go-stone/pkg/models"
)

var (
	transfers []models.Transfer
)

// GetAllTransfers returns a list of all the transfers from the authenticated user.
func GetAllTransfers(w http.ResponseWriter, r *http.Request) {
	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	claims := common.ExtractClaimsFromToken(w, r)

	for _, transfer := range mocks.Transfers {
		if transfer.AccountOriginId == claims.AccountID {
			transfers = append(transfers, transfer)
		}

	}

	w.Header().Add("Content-Type", "application/json")

	if len(transfers) == 0 {
		w.Write([]byte(fmt.Sprintf("No transfers found for the user %s", claims.Cpf)))
		return
	}

	json.NewEncoder(w).Encode(transfers)
}

// MakeTransfer performs a transfer between two users. The origin account its taken form the token and the destination account and balance from the http body.
func MakeTransfer(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")

	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var transfer models.Transfer
	json.Unmarshal(body, &transfer)

	claims := common.ExtractClaimsFromToken(w, r)

	for _, account := range mocks.Accounts {
		if claims.AccountID == account.Id {
			if transfer.Amount > account.Balance {
				json.NewEncoder(w).Encode("This account does not have enough balance to make the transfer.")
				return
			}
			account.Balance -= transfer.Amount
			json.NewEncoder(w).Encode(account)
		}
		if transfer.AccountDestinationId == account.Id {
			account.Balance += transfer.Amount
			json.NewEncoder(w).Encode(account)
		}
	}

}
