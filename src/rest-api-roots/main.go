package main

import (
	"fmt"

	"github.com/felipealvesprestes/rest-api-roots/models"
	"github.com/felipealvesprestes/rest-api-roots/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Pessoa 1", Historia: "O melhor procrastinador ja conhecido!"},
		{Nome: "Pessoa 2", Historia: "O melhor procrastinador ja conhecido!"},
	}

	fmt.Println("Iniciando o servidor.")
	routes.HandleRequest()
}
