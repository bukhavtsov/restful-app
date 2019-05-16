package main

import (
	"github.com/bukhavtsov/restful-app/pkg/apis"
	"github.com/bukhavtsov/restful-app/pkg/data"
	"github.com/bukhavtsov/restful-app/pkg/db"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	engine   = "postgres"
	username = "postgres"
	password = "root"
	name     = "restful_app"
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
	connection := db.GetConnection(engine, username, password, name)
	defer connection.Close()
	apis.ServeCustomerResource(r, data.NewCustomerData(connection))
	apis.ServeDeveloperResource(r, data.NewDeveloperData(connection))
	apis.ServeUserResource(r, data.NewUserData(connection))
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", wrap(r)))
}
