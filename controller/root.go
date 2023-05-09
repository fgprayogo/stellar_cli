package controller

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
)

func Init(){
	r := mux.NewRouter()
	// Account Controller Route
	r.HandleFunc("/accounts/{accountID}/balance", GetAccountBalances)
	log.Println("Service started at http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}