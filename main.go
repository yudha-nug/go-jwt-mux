package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yudha-nug/go-jwt-mux/controller/authcontroller"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/login", authcontroller.Login).Methods("POST")
	r.HandleFunc("/register", authcontroller.Register).Methods("POST")
	r.HandleFunc("/logout", authcontroller.Logout).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
