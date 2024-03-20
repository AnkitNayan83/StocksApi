package userhandlers

import (
	"net/http"

	jsonresponse "github.com/AnkitNayan83/StocksApi/jsonResponse"
)

func HandlerStatus(w http.ResponseWriter, r *http.Request) {
	type msg struct {
		Msg string `json:"message"`
	}
	
	
	jsonresponse.RespondWithJson(w,200,msg{Msg: "USER Route"})
}