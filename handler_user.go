package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (apiCfg apiConfig) handlerRegisterUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email string `json:"email"`
		FirstName string `json:"firstName"`
		LastName string `json:"lastName"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Email == "" {
		RespondWithError(w, 400, fmt.Sprintf("Email is required: %v", err))
		return
	}

	if params.Password == ""{
		RespondWithError(w, 400, fmt.Sprintf("Password is required: %v", err))
		return
	}

	if len(params.Password) < 6 {
		RespondWithError(w, 400, "Password is too small")
		return
	}

	if params.FirstName == "" {
		RespondWithError(w, 400, fmt.Sprintf("FirstName is required: %v", err))
		return
	}

	lastName := sql.NullString{}

	if params.LastName != "" {
		lastName.String = params.LastName
		lastName.Valid = true
	}

	byte_password := []byte(params.Password)

	hashedPassword, err := bcrypt.GenerateFromPassword(byte_password,bcrypt.DefaultCost)

	log.Print(string(hashedPassword))

	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("Failed to hash password: %v", err))
		return
	}

    

	user,err := apiCfg.DB.CreateUser(r.Context(),database.CreateUserParams{
		ID: uuid.New(),
		Email: params.Email,
		Firstname: params.FirstName,
		Lastname: lastName,
		Hashedpassword: string(hashedPassword),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("Failed to create user: %v", err))
		return
	}

    userToken, err := createToken(user.ID.String())

	if err != nil {
		RespondWithError(w, 500, fmt.Sprintf("Failed to create token: %v", err))
		return
	}

	RespondWithJson(w,200,databaseUserToUser(user,userToken))

}


func (apiCfg apiConfig) handlerLogin(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email string `json:"email"`
		Password string `json:"password"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.Email == "" {
		RespondWithError(w, 400, fmt.Sprintf("Email is required: %v", err))
		return
	}

	if params.Password == ""{
		RespondWithError(w, 400, fmt.Sprintf("Password is required: %v", err))
		return
	}

	user,err:=apiCfg.DB.GetUserWithEmail(r.Context(),params.Email)

	if err!=nil {
		RespondWithError(w,404,fmt.Sprintf("User not found: %v",err))
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Hashedpassword),[]byte(params.Password))

	if err != nil {
		RespondWithError(w,401,"Wrong email or password")
		return
	}

	userToken, err := createToken(user.ID.String())

	if err != nil {
		RespondWithError(w,401,fmt.Sprintf("Failed to generate token: %v",err))
		return
	}

	RespondWithJson(w,200,databaseUserToUser(user,userToken))
}
