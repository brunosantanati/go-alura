package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func conectarComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=mysecretpassword host=172.17.0.2 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

//para rodar esse projeto: go run main.go

var tmpl = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectarComBancoDeDados()
	defer db.Close()

	http.HandleFunc("/", Index)
	http.ListenAndServe(":8000", nil)
}

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul, bem bonita", Preco: 39, Quantidade: 5},
		{"Tênis", "Confortável", 89, 3},
		{"Fone", "Muito bom", 59, 2},
		{"Produto novo", "Muito legal", 1.99, 1},
	}

	tmpl.ExecuteTemplate(w, "Index", produtos)
}
