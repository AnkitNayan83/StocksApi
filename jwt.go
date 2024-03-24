package main

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


func createToken(userId string) (string, error) {
	secret := os.Getenv("JWT_KEY")
	if secret == "" {
		return "", errors.New("no JWT Key found")
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