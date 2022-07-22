package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	log.Println("API is running!")
	http.ListenAndServe(":9001", router)
}
