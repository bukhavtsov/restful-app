package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/bukhavtsov/restful-app/models"
)

const (
	secretKey = "eXamp1eK3y"
	iss       = "restful-app"
)

type payload struct {
	Iss string      `json:"iss"`
	Sub models.User `json:"sub"`
	Exp int         `json:"exp"`
}

func isAuthorized(token string) {

}
func GenerateUserToken(user models.User) (tokenString string, err error) {
	payloadJSON, err := json.Marshal(payload{iss, user, 100})
	if err != nil {
		return "", err
	}

	header := base64.StdEncoding.EncodeToString([]byte(`{ "alg": "HS256", "typ": "JWT"}`))
	payload := base64.StdEncoding.EncodeToString(payloadJSON)
	signature := base64.StdEncoding.EncodeToString([]byte(generateSignature(header, payload)))
	return fmt.Sprintf("%s.%s.%s", header, payload, signature), nil
}

func generateSignature(header, payload string) string {
	h := sha256.New()
	h.Write([]byte(secretKey))
	result := fmt.Sprintf("%s.%s.%s", header, payload, hex.EncodeToString(h.Sum(nil)))
	return base64.StdEncoding.EncodeToString([]byte(hex.EncodeToString(hmac.New(sha256.New, []byte(result)).Sum(nil))))
}
