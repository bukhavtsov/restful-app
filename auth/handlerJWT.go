package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bukhavtsov/restful-app/daos"
	"github.com/bukhavtsov/restful-app/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

const (
	secretKey = "eXamp1eK3y"
	iss       = "restful-app"
)

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v\n", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func JWTMiddleware(endPoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			tokenCookie, err := r.Cookie("token")
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusGone)
			}
			token, err := ParseToken(tokenCookie.Value)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusGone)
			}
			user, err := GetUser(token)
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusGone)
			}
			dao := daos.UserDAO{}
			user, err = dao.Get(user.Login, user.Password)
			if err != nil {
				log.Println("user hasn't been found:", err)
				w.WriteHeader(http.StatusNotFound)
			}
			endPoint(w, r)
		})
}
func GetUser(token *jwt.Token) (*models.User, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userJson := fmt.Sprintf("%v", claims["jti"])
		var user models.User
		err := json.Unmarshal([]byte(userJson), &user)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	err := errors.New("user not found") // remake
	return nil, err
}

func GenerateUserToken(user models.User) (tokenString string, err error) {
	jsonUser, err := json.Marshal(user)
	if err != nil {
		return "", err
	}
	claims := jwt.StandardClaims{
		Issuer:    iss,
		Id:        string(jsonUser),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
