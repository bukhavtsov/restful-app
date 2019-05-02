package apis

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/auth"
	"github.com/bukhavtsov/restful-app/daos"
	"github.com/bukhavtsov/restful-app/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type userDAO interface {
	Get(login, password string) (*models.User, error)
	Create(user *models.User) (int64, error)
}

func ServeUserResource(r *mux.Router) {
	r.HandleFunc("/signIn", singIn).Methods("GET")
	r.HandleFunc("/signUp", SignUp).Methods("POST")
}

func singIn(writer http.ResponseWriter, request *http.Request) {
	var dao userDAO = daos.UserDAO{}
	login := request.URL.Query().Get("login")
	password := request.URL.Query().Get("password")
	user, err := dao.Get(login, password)
	if err != nil {
		log.Println("user hasn't been found")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	token, err := auth.GenerateUserToken(*user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	writer.WriteHeader(http.StatusOK)
}

func SignUp(writer http.ResponseWriter, request *http.Request) {
	var dao userDAO = daos.UserDAO{}
	user := new(models.User)
	err := json.NewDecoder(request.Body).Decode(&user)
	user, err = dao.Get(user.Login, user.Password)
	if err == nil && user != nil {
		log.Println("user has been found")
		writer.WriteHeader(http.StatusNotFound)
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
	token, err := auth.GenerateUserToken(*user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:  "token",
		Value: token,
	})
	writer.Header().Set("Location", fmt.Sprintf("/users/%d", userId))
	writer.WriteHeader(http.StatusCreated)
}
