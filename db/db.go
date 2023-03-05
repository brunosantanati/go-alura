package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=mysecretpassword host=172.17.0.2 sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}
