package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mgbcaio/desafio-go-stone/pkg/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/accounts", handlers.GetAllAccounts).Methods(http.MethodGet)

	log.Println("API is running!")

	err := http.ListenAndServe(":9090", router)
	if err != nil {
		log.Fatalf("ListenAndServe error: %s", err)
	}

}
