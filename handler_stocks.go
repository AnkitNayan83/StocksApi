package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerCreateStocks(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		CompanyName string `json:"companyName"`
		ValuePerStocks float64 `json:"valuePerStocks"`
		Quantity int `json:"quantity"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.CompanyName == "" {
		RespondWithError(w,400,"Company Name is required")
		return
	}

	if params.Quantity == 0 {
		RespondWithError(w,400,"Quantity of stock cannot be zero")
		return
	}

	if params.ValuePerStocks == 0.0 {
		RespondWithError(w,400,"Price of stock cannot be zero")
		return
	}

	stock,err := apiCfg.DB.CreateStocks(r.Context(),database.CreateStocksParams{
		ID: uuid.New(),
		Companyname: params.CompanyName,
		Valueperstock: params.ValuePerStocks,
		Quantity: int32(params.ValuePerStocks),
		Ownerid: user.ID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		RespondWithError(w,500,fmt.Sprintf("Failed to create stock: %v",err))
		return
	}

	RespondWithJson(w,201,databaseStockToStock(stock))
}


func (apiCfg apiConfig) handlerGetStocks(w http.ResponseWriter, r *http.Request) {

	stocks, err := apiCfg.DB.GetAllStocks(r.Context())

	if err != nil {
		RespondWithError(w,500,fmt.Sprintf("Failed to fetch stocks: %v",err))
		return
	}

	RespondWithJson(w,200,databaseStocksToStocks(stocks))
}