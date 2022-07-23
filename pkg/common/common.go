package common

import (
	"net/http"
	"strings"

	"github.com/mgbcaio/desafio-go-stone/pkg/auth"
)

func ValidateToken(w http.ResponseWriter, r *http.Request) (err error) {
	err = auth.ExtractAndValidateToken(r)
	if err != nil {
		if strings.Contains(err.Error(), auth.BadRequestErr) {
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
	}
	return
}
