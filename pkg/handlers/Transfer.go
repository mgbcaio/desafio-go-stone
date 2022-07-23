package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mgbcaio/desafio-go-stone/pkg/auth"
	"github.com/mgbcaio/desafio-go-stone/pkg/common"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
	"github.com/mgbcaio/desafio-go-stone/pkg/models"
)

var (
	transfers []models.Transfer
)

func GetAllTransfers(w http.ResponseWriter, r *http.Request) {
	err := common.ValidateToken(w, r)
	if err != nil {
		return
	}

	token, err := auth.ExtractToken(r)
	if err != nil {
		if err.Error() == auth.BadRequestErr {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if err.Error() == auth.UnauthorizedErr {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	_, claims, err := auth.ParseToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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
