package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
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
	produtos := []Produto{
		{Nome: "Básica branca", Descricao: "Camiseta básica branca", Preco: 39.90, Quantidade: 11},
		{Nome: "Básica preta", Descricao: "Camiseta básica preta", Preco: 37.90, Quantidade: 8},
		{Nome: "Básica Azul", Descricao: "Camiseta básica azul", Preco: 34.90, Quantidade: 6},
	}

	templatePage.ExecuteTemplate(w, "Index", produtos)
}
