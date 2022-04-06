package controllers

import (
	"html/template"
	"net/http"

	"github.com/felipealvesprestes/web-fundamentals/models"
)

var templatePage = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.BuscaTodosProdutos()
	templatePage.ExecuteTemplate(w, "Index", produtos)
}
