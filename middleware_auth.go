package main

import (
	"fmt"
	"net/http"

	"github.com/AnkitNayan83/StocksApi/internal/auth"
	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/google/uuid"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)


func (cfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId, err := auth.GetApiToken(r.Header)

		if err != nil {
			RespondWithError(w,401,fmt.Sprintf("api key error: %v",err))
			return
		}

		uuid_userId, err := uuid.Parse(userId);

		if err != nil {
			RespondWithError(w,401,"Cannot parse id")
			return
		}

		user,err := cfg.DB.GetUserWithId(r.Context(),uuid_userId)

		if err != nil {
			RespondWithError(w,404,"No user found with this id")
			return
		}

		handler(w,r,user)
	}
}