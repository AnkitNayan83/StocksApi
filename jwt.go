package main

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)


func createToken(userId string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", err;
	}

	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return "", err
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"userId" : userId,
			"exp" : time.Now().Add(time.Hour * 72).Unix(),
		})
	tokenString, err  := token.SignedString([]byte(secret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}