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

	findaAllProducts, err := db.Query("select * from produtos order by 1 desc")

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

	insertProduct, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		log.Fatal("Erro ao cadastrar o novo produto.")
	}

	insertProduct.Exec(name, description, price, quantity)

	defer db.Close()
}

func DeleteProductById(productId string) {
	db := config.DbConnect()

	deleteProduct, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		log.Fatal("Erro ao excluir o produto.")
	}

	deleteProduct.Exec(productId)

	defer db.Close()
}

func GetProductById(productId int) Produto {
	db := config.DbConnect()

	getProduct, err := db.Query("select * from produtos where id = $1", productId)

	if err != nil {
		log.Fatal("Erro ao buscar as informações do produto.", err)
	}

	product := Produto{}

	for getProduct.Next() {
		var id int
		var name string
		var description string
		var price float64
		var quantity int

		err = getProduct.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			log.Fatal(err)
		}

		product.Id = id
		product.Nome = name
		product.Descricao = description
		product.Preco = price
		product.Quantidade = quantity
	}

	defer db.Close()

	return product
}

func UpdateProduct(id string, name string, description string, price float64, quantity int) {
	db := config.DbConnect()

	updateProduct, err := db.Prepare("update produtos set nome = $2, descricao = $3, preco = $4, quantidade = $5 where id = $1")

	if err != nil {
		log.Fatal("Erro ao cadastrar o novo produto.")
	}

	updateProduct.Exec(id, name, description, price, quantity)

	defer db.Close()
}
