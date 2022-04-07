package main

import (
	"fmt"

	"rest-api-roots/database"
	"rest-api-roots/models"
	"rest-api-roots/routes"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Pessoa 1", Historia: "O melhor procrastinador ja conhecido!"},
		{Id: 2, Nome: "Pessoa 2", Historia: "O melhor procrastinador ja conhecido!"},
	}

	database.ConectaBanco()

	fmt.Println("Iniciando o servidor.")
	routes.HandleRequest()
}
