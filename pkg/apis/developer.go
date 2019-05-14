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

type developerDAO interface {
	Create(developer *models.Developer) (int64, error)
	Read(id int64) (*models.Developer, error)
	ReadAll() ([]*models.Developer, error)
	Update(developer *models.Developer) (*models.Developer, error)
	Delete(id int64) error
}

type developerAPI struct {
	dao developerDAO
}

func ServeDeveloperResource(r *mux.Router, dao developerDAO) {
	r.Handle("/developers", jwt.VerifyPermission(developerAPI{dao}.getDevelopers)).Methods("GET")
	r.Handle("/developers/{id}", jwt.VerifyPermission(developerAPI{dao}.getDeveloper)).Methods("GET")
	r.Handle("/developers", jwt.VerifyPermission(developerAPI{dao}.createDeveloper)).Methods("POST")
	r.Handle("/developers/{id}", jwt.VerifyPermission(developerAPI{dao}.updateDeveloper)).Methods("PUT")
	r.Handle("/developers/{id}", jwt.VerifyPermission(developerAPI{dao}.deleteDeveloper)).Methods("DELETE")
}

func (api developerAPI) getDevelopers(writer http.ResponseWriter, request *http.Request) {
	developers, err := api.dao.ReadAll()
	if err != nil {
		log.Println("developers haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	err = json.NewEncoder(writer).Encode(developers)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api developerAPI) getDeveloper(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developer, err := api.dao.Read(id)
	if err != nil {
		log.Println("developer hasn't been read")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	err = json.NewEncoder(writer).Encode(developer)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusOK)
}

func (api developerAPI) createDeveloper(writer http.ResponseWriter, request *http.Request) {
	developer := new(models.Developer)
	err := json.NewDecoder(request.Body).Decode(&developer)
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developerId, err := api.dao.Create(developer)
	if err != nil {
		log.Println("developer hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/developers/%d", developerId))
	writer.WriteHeader(http.StatusCreated)
}

func (api developerAPI) updateDeveloper(writer http.ResponseWriter, request *http.Request) {
	developer := new(models.Developer)
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
	updatedDeveloper, err := api.dao.Update(developer)
	if err != nil {
		log.Println("developer hasn't been updated")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(writer).Encode(updatedDeveloper)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}

func (api developerAPI) deleteDeveloper(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println(err)
	}
	err = api.dao.Delete(id)
	if err != nil {
		log.Println("developer hasn't been removed")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusNoContent)
}
