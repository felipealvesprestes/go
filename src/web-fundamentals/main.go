package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var templatePage = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectaBanco()

	selectProdutos, err := db.Query("select * from produtos")

	if err != nil {
		log.Fatal(err)
	}

	produto := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id int
		var nome string
		var descricao string
		var preco float64
		var quantidade int

		err = selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			log.Fatal(err)
		}

		produto.Nome = nome
		produto.Descricao = descricao
		produto.Preco = preco
		produto.Quantidade = quantidade

		produtos = append(produtos, produto)
	}

	templatePage.ExecuteTemplate(w, "Index", produtos)

	defer db.Close()
}

func conectaBanco() *sql.DB {
	stringConexao := "user=postgres password=postgres dbname=loja_alura host=localhost sslmode=disable"
	db, err := sql.Open("postgres", stringConexao)

	if err != nil {
		log.Fatal(err)
	}

	return db
}
