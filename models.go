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

type Stock struct {
	ID            uuid.UUID `json:"id"`
	Companyname   string `json:"companyName"`
	Valueperstock float64 `json:"valuePerStock"`
	Quantity      int32 `json:"quantity"`
	Ownerid       uuid.UUID `json:"ownerId"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

func databaseStockToStock(dbStock database.Stock) Stock {
	return Stock{
		ID: dbStock.ID,
		Companyname: dbStock.Companyname,
		Valueperstock: dbStock.Valueperstock,
		Quantity: dbStock.Quantity,
		Ownerid: dbStock.Ownerid,
		CreatedAt: dbStock.CreatedAt,
		UpdatedAt: dbStock.UpdatedAt,
	}
}

func databaseStocksToStocks(dbStocks []database.Stock) []Stock{
	stocks := []Stock{}

	for _,stock := range dbStocks {
		stocks = append(stocks, databaseStockToStock(stock))
	}

	return stocks;
}



type Stockbuy struct {
	ID        uuid.UUID `json:"id"`
	Userid    uuid.UUID `json:"userId"`
	Stockid   uuid.UUID `json:"stockId"`
	Quantity  int32 `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func databaseStockBuyToStockBuy(dbStockBuy database.Stockbuy) Stockbuy {
	return Stockbuy{
		ID: dbStockBuy.ID,
		Userid: dbStockBuy.Userid,
		Stockid: dbStockBuy.Stockid,
		Quantity: dbStockBuy.Quantity,
		CreatedAt: dbStockBuy.CreatedAt,
		UpdatedAt: dbStockBuy.UpdatedAt,
	}
}