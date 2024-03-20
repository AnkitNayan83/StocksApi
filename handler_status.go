package main

import (
	"net/http"

	jsonresponse "github.com/AnkitNayan83/StocksApi/jsonResponse"
)

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	type msg struct {
		Msg string `json:"message"`
	}
	jsonresponse.RespondWithJson(w, 200, msg{Msg: "Server is up and running"})
}
