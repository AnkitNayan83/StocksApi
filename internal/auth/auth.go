package auth

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func GetApiToken(headers http.Header) (string,error) {
	err := godotenv.Load()
	if err != nil {
		return "", err;
	}

	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return "", err
	}

	secretKey := []byte(secret)

	val := headers.Get("Authorization");
	if val == "" {
		return "", errors.New("no token found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "",errors.New("wrong token format")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("wrong token format")
	}

	token, err := jwt.Parse(vals[1], func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	 })

	 if err != nil {
		return "",errors.New("invalid token");
	 }

	 if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if userId, ok := claims["userId"].(string); ok {
			return userId, nil
		} else {
			return "", errors.New("user id not found in the token")
		}
	 } else {
		return "", errors.New("invalid token")
	 }
}