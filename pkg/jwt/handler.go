package jwt

import (
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/pkg/data"
	"github.com/bukhavtsov/restful-app/pkg/models"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	secretKeyAccess  = "eXamp1eK3yACceS$"
	secretKeyRefresh = "r3Fr3S4eXamp1eK3y"
	iss              = "restful-app"
	refreshTokenName = "refresh_token"
	accessTokenName  = "access_token"
)

type jti struct {
	Id   int64  `json:"Id"`
	Role string `json:"Role"`
}

func parse(tokenString, secretKey string) (*jwt.Token, error) {
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

func VerifyPermission(endPoint func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		accessCookie, err := r.Cookie(accessTokenName)
		if err == nil && isVerifiedAccess(accessCookie.Value) {
			endPoint(w, r)
			return
		}
		refreshCookie, err := r.Cookie(refreshTokenName)
		if err == nil && isVerifiedRefresh(refreshCookie.Value) {
			user, err := getUser(refreshCookie.Value, secretKeyRefresh)
			if err != nil {
				log.Println(err)
				return
			}
			updatedAccess, err := getUpdatedAccess(*user)
			if err != nil {
				fmt.Println("accessCookie token :", err)
				return
			}
			http.SetCookie(w, &http.Cookie{Name: accessTokenName, Value: updatedAccess})
			endPoint(w, r)
			return
		}
		w.WriteHeader(http.StatusGone)
	})
}

func isVerifiedAccess(access string) bool {
	if !isValidTime(access, secretKeyAccess) {
		log.Println("access token time is over")
		return false
	}
	token, err := parse(access, secretKeyAccess)
	if err != nil {
		log.Println(err)
		return false
	}
	jti, err := GetJTI(token)
	if err != nil {
		log.Println(err)
		return false
	}
	user, err := data.NewUserDAO().GetById(jti.Id)
	if err != nil {
		log.Println("user hasn't been found:", err)
		return false
	}
	if user.Role != jti.Role {
		log.Println(err)
		return false
	}
	return true
}

func getUpdatedAccess(user models.User) (access string, err error) {
	access, err = GenerateAccess(user)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func isVerifiedRefresh(refresh string) bool {
	if !isValidTime(refresh, secretKeyRefresh) {
		log.Println("refresh token time is over")
		return false
	}
	_, err := getUser(refresh, secretKeyRefresh)
	if err != nil {
		return false
	}
	return true
}

func getUser(tokenString, secretKeyAccess string) (*models.User, error) {
	token, err := parse(tokenString, secretKeyAccess)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	jti, err := GetJTI(token)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	user, err := data.NewUserDAO().GetById(jti.Id)
	if err != nil {
		log.Println("user hasn't been found:", err)
		return nil, err
	}
	if user.Role != jti.Role {
		log.Println("user has different role")
		return nil, err
	}
	return user, err
}

func isValidTime(tokenString, secretKey string) bool {
	token, err := parse(tokenString, secretKey)
	if err != nil {
		log.Println(err)
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		expString := fmt.Sprintf("%f", claims["exp"])
		exp, err := strconv.ParseFloat(expString, 64)
		if err != nil {
			fmt.Println(err)
			return false
		}
		now := float64(time.Now().Unix())
		if exp > now {
			return true
		}
	}
	return false
}

func GetJTI(token *jwt.Token) (*jti, error) {
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		jtiJson := fmt.Sprintf("%v", claims["jti"])
		var jti jti
		err := json.Unmarshal([]byte(jtiJson), &jti)
		if err != nil {
			return nil, err
		}
		return &jti, nil
	}
	return nil, fmt.Errorf("user hasn't been found")
}

func GenerateAccess(user models.User) (tokenString string, err error) {
	jti, err := json.Marshal(&jti{user.Id, user.Role})
	if err != nil {
		return "", err
	}
	claims := jwt.StandardClaims{
		Issuer:    iss,
		Id:        string(jti),
		ExpiresAt: time.Now().Add(time.Hour).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(secretKeyAccess))
	if err != nil {
		return "", err
	}
	return token, nil
}

func GenerateRefresh(user models.User) (tokenString string, err error) {
	jti, err := json.Marshal(&jti{user.Id, user.Role})
	if err != nil {
		return "", err
	}
	claims := jwt.StandardClaims{
		Issuer:    iss,
		Id:        string(jti),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := rawToken.SignedString([]byte(secretKeyRefresh))
	if err != nil {
		return "", err
	}
	return token, nil
}
