package common

import (
	"net/http"
	"strings"

	"github.com/mgbcaio/desafio-go-stone/pkg/auth"
)

// ValidateToken validates the token extracted from the cookies. returns an error if something goes wrong.
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

// ExtractClaimsFromToken extracts the claims from the token extracted from the cookies and returns it.
func ExtractClaimsFromToken(w http.ResponseWriter, r *http.Request) *auth.Claims {
	token, err := auth.ExtractToken(r)
	if err != nil {
		if err.Error() == auth.BadRequestErr {
			w.WriteHeader(http.StatusBadRequest)
			return nil
		}
		if err.Error() == auth.UnauthorizedErr {
			w.WriteHeader(http.StatusUnauthorized)
			return nil
		}
	}

	_, claims, err := auth.ParseToken(token)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}

	return claims
}
