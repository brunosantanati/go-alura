package routes

import (
	"net/http"

	"brunosantana.me/go-alura/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
