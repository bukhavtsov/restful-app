package apis

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/pkg/jwt"
	"github.com/bukhavtsov/restful-app/pkg/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type userDAO interface {
	Get(login, password string) (*models.User, error)
	Create(user *models.User) (int64, error)
	GetById(id int64) (*models.User, error)
	Update(user *models.User, refreshToken string) (*models.User, error)
}

type userAPI struct {
	dao userDAO
}

func ServeUserResource(r *mux.Router, dao userDAO) {
	r.HandleFunc("/signin", userAPI{dao}.singIn).Methods("GET")
	r.HandleFunc("/signup", userAPI{dao}.signUp).Methods("POST")
}

func (api userAPI) singIn(writer http.ResponseWriter, request *http.Request) {
	login := request.URL.Query().Get("login")
	password := request.URL.Query().Get("password")
	user, err := api.dao.Get(login, password)
	if err != nil {
		log.Println("user hasn't been found")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	refresh, err := jwt.GenerateRefresh(*user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = api.dao.Update(user, refresh)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	access, err := jwt.GenerateAccess(*user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(writer, &http.Cookie{Name: "access_token", Value: access})
	http.SetCookie(writer, &http.Cookie{Name: "refresh_token", Value: refresh})
	writer.WriteHeader(http.StatusOK)
}

func (api userAPI) signUp(writer http.ResponseWriter, request *http.Request) {
	var user models.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		log.Printf("failed reading JSON: %\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	_, err = api.dao.Get(user.Login, user.Password)
	if err == nil {
		log.Println("user has been found")
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	user.RefreshToken, err = jwt.GenerateRefresh(user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	userId, err := api.dao.Create(&user)
	if err != nil {
		log.Println("user hasn't been created")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	token, err := jwt.GenerateAccess(user)
	if err != nil {
		log.Println(err)
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	http.SetCookie(writer, &http.Cookie{
		Name:  "access_token",
		Value: token,
	})
	writer.Header().Set("Location", fmt.Sprintf("/users/%d", userId))
	writer.WriteHeader(http.StatusCreated)
}
