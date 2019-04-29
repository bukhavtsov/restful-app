package apis

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/daos"
	"github.com/bukhavtsov/restful-app/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type userDAO interface {
	Get(login, password string) (bool, error)
	Create(user *models.User) (int64, error)
}

func ServeUserResource(r *mux.Router) {
	r.HandleFunc("/login/{username}/{password}", singIn).Methods("GET")
	r.HandleFunc("/registration", SignUp).Methods("POST")
}

func singIn(writer http.ResponseWriter, request *http.Request) {
	var dao userDAO = daos.UserDAO{}
	params := mux.Vars(request)
	login := params["username"]
	password := params["password"]
	if found, err := dao.Get(login, password); !found || err != nil {
		log.Println("user hasn't been found")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
}
func SignUp(writer http.ResponseWriter, request *http.Request) {
	var dao userDAO = daos.UserDAO{}
	user := new(models.User)
	err := json.NewDecoder(request.Body).Decode(&user)
	if found, err := dao.Get(user.Login, user.Password); found || err != nil {
		log.Println("user has been found")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	userId, err := dao.Create(user)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	writer.Header().Set("Location", fmt.Sprintf("/users/%d", userId))
	writer.WriteHeader(http.StatusCreated)
}
