package apis

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/pkg/jwt"
	"github.com/bukhavtsov/restful-app/pkg/models"
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

type customerAPI struct {
	dao customerDAO
}

func ServeCustomerResource(r *mux.Router, dao customerDAO) {
	r.Handle("/customers", jwt.VerifyPermission(customerAPI{dao}.getCustomers)).Methods("GET")
	r.Handle("/customers/{id}", jwt.VerifyPermission(customerAPI{dao}.getCustomer)).Methods("GET")
	r.Handle("/customers", jwt.VerifyPermission(customerAPI{dao}.createCustomer)).Methods("POST")
	r.Handle("/customers/{id}", jwt.VerifyPermission(customerAPI{dao}.updateCustomer)).Methods("PUT")
	r.Handle("/customers/{id}", jwt.VerifyPermission(customerAPI{dao}.deleteCustomer)).Methods("DELETE")
}

func (api customerAPI) getCustomers(writer http.ResponseWriter, request *http.Request) {
	customers, err := api.dao.ReadAll()
	if err != nil {
		log.Println("customers haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(customers)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api customerAPI) getCustomer(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	customer, err := api.dao.Read(id)
	if err != nil {
		log.Println("customer hasn't been read")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(customer)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api customerAPI) createCustomer(writer http.ResponseWriter, request *http.Request) {
	customer := new(models.Customer)
	err := json.NewDecoder(request.Body).Decode(&customer)
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	customerId, err := api.dao.Create(customer)
	if err != nil {
		log.Println("customer hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/customers/%d", customerId))
	writer.WriteHeader(http.StatusCreated)
}

func (api customerAPI) updateCustomer(writer http.ResponseWriter, request *http.Request) {
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
	updatedCustomer, err := api.dao.Update(customer)
	if err != nil {
		log.Println("customer hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(updatedCustomer)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (api customerAPI) deleteCustomer(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = api.dao.Delete(id)
	if err != nil {
		log.Println("customer hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
