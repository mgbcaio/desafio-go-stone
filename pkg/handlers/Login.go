package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mgbcaio/desafio-go-stone/pkg/auth"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, u := range mocks.Users {
		if u.Cpf == creds.Cpf {
			expectedSecret := u.Secret
			if creds.Secret != expectedSecret {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			token, expTime, err := auth.GenerateJWT(creds.Cpf)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   token,
				Expires: expTime,
			})

			json.NewEncoder(w).Encode(token)
			return

		}

	}
	w.WriteHeader(http.StatusUnauthorized)

}
