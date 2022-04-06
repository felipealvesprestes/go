package models

import (
	"log"

	"github.com/felipealvesprestes/web-fundamentals/config"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func FindAllProducts() []Produto {
	db := config.DbConnect()

	findaAllProducts, err := db.Query("select * from produtos")

	if err != nil {
		log.Fatal(err)
	}

	product := Produto{}
	products := []Produto{}

	for findaAllProducts.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err = findaAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			log.Fatal(err)
		}

		product.Id = id
		product.Nome = name
		product.Descricao = description
		product.Preco = price
		product.Quantidade = quantity

		products = append(products, product)
	}

	defer db.Close()

	return products
}

func InsertNewProduct(name string, description string, price float64, quantity int) {
	db := config.DbConnect()

	insertDatabase, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		log.Fatal("Erro ao cadastrar o novo produto.")
	}

	insertDatabase.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProductById(productId string) {
	db := config.DbConnect()

	deleteDatabase, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		log.Fatal("Erro ao excluir o produto.")
	}

	deleteDatabase.Exec(productId)

	defer db.Close()
}
