package models

import (
	"log"

	"github.com/felipealvesprestes/web-fundamentals/config"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosProdutos() []Produto {
	db := config.ConectaBanco()

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

	defer db.Close()

	return produtos
}
