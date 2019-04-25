package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/developers", ProductsHandler).Methods("GET")
	r.HandleFunc("/developers/{1}", ProductsHandler).Methods("GET")
	r.HandleFunc("/developers", ProductsHandler).Methods("POST")
	r.HandleFunc("/developers/{id}", ProductsHandler).Methods("PUT")
	r.HandleFunc("/developers/{id}", ProductsHandler).Methods("DELETE")
	http.Handle("/", r)
}
