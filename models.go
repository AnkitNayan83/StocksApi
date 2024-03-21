package main

import (
	"time"

	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	Email string `json:"email"`
	FirstName string `json:"firstName"`
	LastName *string `json:"lastName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Token string `json:"token"`
}

func databaseUserToUser(dbUser database.User, token string) User {
	var lastName *string
	if dbUser.Lastname.Valid {
		lastName = &dbUser.Lastname.String
	}
	return User{
		ID: dbUser.ID,
		Email: dbUser.Email,
		FirstName: dbUser.Firstname,
		LastName: lastName,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Token: token,
	}
}