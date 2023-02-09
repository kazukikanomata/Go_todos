package controllers

import (
	"log"
	"net/http"
	"text/template"
)

func top(w http.ResponseWriter, r *http.Request) {
	// 引数の値を解析している
	t, err := template.ParseFiles("app/views/templates/top.html")
	if err != nil {
		log.Fatalln(err)
	}
	t.Execute(w, err)
}
