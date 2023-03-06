package main

import (
	"net/http"

	"brunosantana.me/go-alura/routes"
)

//para rodar esse projeto: go run main.go
//depois precisa acessar: http://localhost:8000/

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
