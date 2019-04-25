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

func getDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println("id is't integer")
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
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(respDeveloper)
}
func createDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developer := new(entities.Developer)
	err := json.NewDecoder(request.Body).Decode(&developer)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developerId, err := dao.Create(developer)
	if err != nil {
		log.Println("developer hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(developerId)
}
func updateDeveloper(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developer := new(entities.Developer)

	params := mux.Vars(request)
	id, err := strconv.ParseInt(params["id"], 0, 64)
	if err != nil {
		log.Println("id is't integer")
	}
	err = json.NewDecoder(request.Body).Decode(&developer)
	if err != nil {
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	fmt.Println(id)
	developer.Id = id
	err = dao.Update(developer)
	if err != nil {
		log.Println("developer hasn't been updated")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
}

func main() {
	r := mux.NewRouter()
	//r.HandleFunc("/developers", ProductsHandler).Methods("GET")
	r.HandleFunc("/developers/{id}", getDeveloper).Methods("GET")
	r.HandleFunc("/developers", createDeveloper).Methods("POST")
	r.HandleFunc("/developers/{id}", updateDeveloper).Methods("PUT")
	//r.HandleFunc("/developers/{id}", ProductsHandler).Methods("DELETE")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
