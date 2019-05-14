package main

import (
	"github.com/bukhavtsov/restful-app/pkg/apis"
	"github.com/bukhavtsov/restful-app/pkg/data"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func loggingMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		h.ServeHTTP(w, r)
	})
}

func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		h.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	apis.ServeCustomerResource(r, data.NewCustomerDAO())
	apis.ServeDeveloperResource(r, data.NewDeveloperDAO())
	apis.ServeUserResource(r, data.NewUserDAO())
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", wrap(r)))
}
