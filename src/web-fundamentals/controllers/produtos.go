package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/felipealvesprestes/web-fundamentals/models"
)

var templatePage = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	templatePage.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	templatePage.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("nome")
		description := r.FormValue("descricao")

		price, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			log.Fatal("Erro na converção do preço.")
		}

		quantity, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			log.Fatal("Erro na converção da quantidade.")
		}

		models.InsertNewProduct(name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	models.DeleteProductById(productId)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	productId, _ := strconv.Atoi(r.URL.Query().Get("id"))
	product := models.GetProductById(productId)
	templatePage.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("nome")
		description := r.FormValue("descricao")

		price, err := strconv.ParseFloat(r.FormValue("preco"), 64)

		if err != nil {
			log.Fatal("Erro na converção do preço.")
		}

		quantity, err := strconv.Atoi(r.FormValue("quantidade"))

		if err != nil {
			log.Fatal("Erro na converção da quantidade.")
		}

		models.UpdateProduct(id, name, description, price, quantity)
	}

	http.Redirect(w, r, "/", 301)
}
