package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbConnect() *sql.DB {
	stringConexao := "user=postgres password=postgres dbname=loja host=localhost sslmode=disable"
	db, err := sql.Open("postgres", stringConexao)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
