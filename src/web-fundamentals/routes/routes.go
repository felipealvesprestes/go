package routes

import (
	"net/http"

	"github.com/felipealvesprestes/web-fundamentals/controllers"
)

func CarregaRoutes() {
	http.HandleFunc("/", controllers.Index)
}
