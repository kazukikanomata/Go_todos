package controllers

import (
	"go_todo/config"
	"net/http"
)

func StartMainServer() error {
	// topというhandlerの処理を行う
	http.HandleFunc("/", top)
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
