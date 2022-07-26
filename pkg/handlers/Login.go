package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mgbcaio/desafio-go-stone/pkg/auth"
	"github.com/mgbcaio/desafio-go-stone/pkg/mocks"
	log "github.com/sirupsen/logrus"
)

// Login performs the authentication of the given user
func Login(w http.ResponseWriter, r *http.Request) {
	var creds auth.Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		log.Errorf("Error occurred: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for _, user := range mocks.Users {
		if user.Cpf == creds.Cpf {
			expectedSecret := user.Secret
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
	log.Errorf("User not authenticated!")
	w.WriteHeader(http.StatusUnauthorized)

}
