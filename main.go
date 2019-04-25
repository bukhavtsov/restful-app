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
		log.Println(err)
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
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(respDeveloper)
	if err != nil {
		fmt.Println(err)
		writer.WriteHeader(http.StatusBadRequest)
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
	writer.Header().Set("Content-Type", "application/json")
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
		log.Printf("failed reading JSON: %s", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	developer.Id = id
	updatedDeveloper, err := dao.Update(developer)
	if err != nil {
		log.Println("developer hasn't been updated")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(updatedDeveloper)
	if err != nil {
		fmt.Println(err)
	}
	writer.WriteHeader(http.StatusNoContent)
}
func getDevelopers(writer http.ResponseWriter, request *http.Request) {
	var dao interfaces.DeveloperDAO
	dao = implementations.DeveloperDAOImpl{}
	developers, err := dao.ReadAll()
	if err != nil {
		log.Println("developers haven't been read")
		writer.WriteHeader(http.StatusNoContent)
		return
	}
	writer.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(writer).Encode(developers)
	if err != nil {
		fmt.Println(err)
	}
	writer.WriteHeader(http.StatusOK)
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
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNoContent)
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
	http.Handle("/", r)
	r.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":8080", r))
}
