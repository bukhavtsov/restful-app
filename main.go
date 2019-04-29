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
func wrap(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
			w.Header().Set("Pragma", "no-cache")
			h.ServeHTTP(w, r)
		})
}
func main() {
	r := mux.NewRouter()
	apis.ServeCustomerResource(r)
	apis.ServeDeveloperResource(r)
	apis.ServeUserResource(r)
	http.Handle("/", r)
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", wrap(r)))
}
