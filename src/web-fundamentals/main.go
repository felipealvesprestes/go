package main

import (
	"net/http"

	"github.com/felipealvesprestes/web-fundamentals/routes"
)

func main() {
	routes.CarregaRoutes()
	http.ListenAndServe(":8000", nil)
}
