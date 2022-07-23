package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mgbcaio/desafio-go-stone/pkg/handlers"
)

func ConfigureRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/accounts", handlers.GetAllAccounts).Methods(http.MethodGet)
	router.HandleFunc("/accounts", handlers.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts/{id}/balance", handlers.GetAccountBalance).Methods(http.MethodGet)

	router.HandleFunc("/login", handlers.Login).Methods(http.MethodPost)

	return router
}

func main() {
	router := ConfigureRouter()

	log.Println("API is running!")

	log.Fatal(http.ListenAndServe(":9090", router))

}
