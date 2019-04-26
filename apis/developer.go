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

type developerDAO interface {
	Create(developer *models.Developer) (int64, error)
	Read(id int64) (*models.Developer, error)
	ReadAll() ([]*models.Developer, error)
	Update(developer *models.Developer) (*models.Developer, error)
	Delete(id int64) error
}

func ServeDeveloperResource(r *mux.Router) {
	r.HandleFunc("/developers", getDevelopers).Methods("GET")
	r.HandleFunc("/developers/{id}", getDeveloper).Methods("GET")
	r.HandleFunc("/developers", createDeveloper).Methods("POST")
	r.HandleFunc("/developers/{id}", updateDeveloper).Methods("PUT")
	r.HandleFunc("/developers/{id}", deleteDeveloper).Methods("DELETE")
}

func getDevelopers(writer http.ResponseWriter, request *http.Request) {
	var dao developerDAO
	dao = daos.DeveloperDAO{}
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
	var dao developerDAO
	dao = daos.DeveloperDAO{}
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
	respDeveloper := models.Developer{
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
	var dao developerDAO
	dao = daos.DeveloperDAO{}
	developer := new(models.Developer)
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
	var dao developerDAO
	dao = daos.DeveloperDAO{}
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
	var dao developerDAO
	dao = daos.DeveloperDAO{}
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

func addCachingHeader(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	writer.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate")
	writer.Header().Set("Pragma", "no-cache")
}
