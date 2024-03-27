package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AnkitNayan83/StocksApi/internal/database"
	"github.com/google/uuid"
)

func (apiCfg apiConfig) handlerBuyStock(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		StockId uuid.UUID `json:"stockId"`
		Quantity int `json:"quantity"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		RespondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	if params.StockId == uuid.Nil {
		RespondWithError(w,400,"stock id not found")
		return
	}

	if params.Quantity == 0 {
		RespondWithError(w,400,"quantity cannot be zero")
		return
	}

	stockBuy,err := apiCfg.DB.CreateStockBuy(r.Context(),database.CreateStockBuyParams{
		ID: uuid.New(),
		Userid: user.ID,
		Stockid: params.StockId,
		Quantity: int32(params.Quantity),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		RespondWithError(w,500,fmt.Sprintf("Failed to create stock: %v",err))
		return
	}

	_,err = apiCfg.DB.UpdateStockQuantity(r.Context(),database.UpdateStockQuantityParams{
		Quantity: int32(params.Quantity),
		ID: params.StockId,
	})

	if err != nil {
		RespondWithError(w,500,fmt.Sprintf("Failed to update stock quantity: %v",err))
		return
	}

	RespondWithJson(w,201,databaseStockBuyToStockBuy(stockBuy))

}