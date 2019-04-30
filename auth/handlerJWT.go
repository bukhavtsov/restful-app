package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bukhavtsov/restful-app/models"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	secretKey = "eXamp1eK3y"
	iss       = "restful-app"
)

type JwtClaims struct {
	Iss string `json:"name"`
	jwt.StandardClaims
}

func ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
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
	claims := JwtClaims{
		iss,
		jwt.StandardClaims{
			Id:        string(jsonUser),
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}
	return token, nil
}
