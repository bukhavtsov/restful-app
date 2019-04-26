package main

import (
	"github.com/bukhavtsov/restful-app/apis"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	apis.ServeCustomerResource(r)
	apis.ServeDeveloperResource(r)
	http.Handle("/", r)
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", r))
}
