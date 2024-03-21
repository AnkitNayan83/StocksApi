package main

import (
	"net/http"
)

func handlerStatus(w http.ResponseWriter, r *http.Request) {
	type msg struct {
		Msg string `json:"message"`
	}
	RespondWithJson(w, 200, msg{Msg: "Server is up and running"})
}
