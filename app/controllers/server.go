package controllers

import (
	"go_todo/config"
	"net/http"
)

func StartMainServer() error {
	return http.ListenAndServe(":"+config.Config.Port, nil)
}
