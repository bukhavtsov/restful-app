package apis

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/daos"
	"github.com/bukhavtsov/restful-app/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type customerDAO interface {
	Create(customer *models.Customer) (int64, error)
	Read(id int64) (*models.Customer, error)
	ReadAll() ([]*models.Customer, error)
	Update(customer *models.Customer) (*models.Customer, error)
	Delete(id int64) error
}

func ServeCustomerResource(r *mux.Router) {
	r.HandleFunc("/customers", getCustomers).Methods("GET")
	r.HandleFunc("/customers/{id}", getCustomer).Methods("GET")
	r.HandleFunc("/customers", createCustomer).Methods("POST")
	r.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	r.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
}

func getCustomers(writer http.ResponseWriter, request *http.Request) {
	var dao customerDAO
	dao = daos.CustomerDAO{}
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
	var dao customerDAO
	dao = daos.CustomerDAO{}
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
	respCustomer := models.Customer{
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
	var dao customerDAO
	dao = daos.CustomerDAO{}
	customer := new(models.Customer)
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
	var dao customerDAO
	dao = daos.CustomerDAO{}
	customer := new(models.Customer)
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
	var dao customerDAO
	dao = daos.CustomerDAO{}
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
