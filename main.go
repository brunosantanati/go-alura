package main

import (
	"html/template"
	"net/http"

	"brunosantana.me/go-alura/models"
)

//para rodar esse projeto: go run main.go
//depois precisa acessar: http://localhost:8000/

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	tmpl.ExecuteTemplate(w, "Index", todosOsProdutos)
}
