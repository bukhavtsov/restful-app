package main

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/dao/entities"
	"github.com/bukhavtsov/restful-app/dao/implementations"
	"github.com/bukhavtsov/restful-app/dao/interfaces"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func getDevelopers(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developers, err := dao.ReadAll()
	if err != nil {
		log.Println("developers haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(developers)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func getDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developer, err := dao.Read(id)
	if err != nil {
		log.Println("developer hasn't been read")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	respDeveloper := entities.Developer{
		Id:           developer.Id,
		Name:         developer.Name,
		Age:          developer.Age,
		PrimarySkill: developer.PrimarySkill,
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(respDeveloper)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func createDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developer := new(entities.Developer)
	err := json.NewDecoder(request.Body).Decode(&developer)
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developerId, err := dao.Create(developer)
	if err != nil {
		log.Println("developer hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	addCachingHeader(writer)
	writer.Header().Set("Location", fmt.Sprintf("/developers/%d", developerId))
	writer.WriteHeader(http.StatusCreated)
}

func updateDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developer := new(entities.Developer)
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(request.Body).Decode(&developer)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developer.Id = id
	updatedDeveloper, err := dao.Update(developer)
	if err != nil {
		log.Println("developer hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(updatedDeveloper)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func deleteDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = dao.Delete(id)
	if err != nil {
		log.Println("developer hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	addCachingHeader(writer)
	writer.WriteHeader(http.StatusNoContent)
}

//--------------------------------------------------------------------------

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.CustomerDAO
	dao = implementations.CustomerDAOImpl{}
	customers, err := dao.ReadAll()
	if err != nil {
		log.Println("customers haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(customers)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func getCustomer(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.CustomerDAO
	dao = implementations.CustomerDAOImpl{}
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	customer, err := dao.Read(id)
	if err != nil {
		log.Println("customer hasn't been read")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	respCustomer := entities.Customer{
		Id:       customer.Id,
		Name:     customer.Name,
		Money:    customer.Money,
		Discount: customer.Discount,
		State:    customer.State,
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(respCustomer)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func createCustomer(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.CustomerDAO
	dao = implementations.CustomerDAOImpl{}
	customer := new(entities.Customer)
	err := json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	customerId, err := dao.Create(customer)
	if err != nil {
		log.Println("customer hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	addCachingHeader(writer)
	writer.Header().Set("Location", fmt.Sprintf("/customers/%d", customerId))
	writer.WriteHeader(http.StatusCreated)
}

func updateCustomer(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.CustomerDAO
	dao = implementations.CustomerDAOImpl{}
	customer := new(entities.Customer)
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		log.Printf("failed reading JSON: %s\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	customer.Id = id
	updatedCustomer, err := dao.Update(customer)
	if err != nil {
		log.Println("customer hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	addCachingHeader(writer)
	err = json.NewEncoder(writer).Encode(updatedCustomer)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func deleteCustomer(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.CustomerDAO
	dao = implementations.CustomerDAOImpl{}
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = dao.Delete(id)
	if err != nil {
		log.Println("customer hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	addCachingHeader(writer)
	writer.WriteHeader(http.StatusNoContent)
}

func addCachingHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/developers", getDevelopers).Methods("GET")
	r.HandleFunc("/developers/{id}", getDeveloper).Methods("GET")
	r.HandleFunc("/developers", createDeveloper).Methods("POST")
	r.HandleFunc("/developers/{id}", updateDeveloper).Methods("PUT")
	r.HandleFunc("/developers/{id}", deleteDeveloper).Methods("DELETE")

	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")

	http.Handle("/", r)
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", r))
}
