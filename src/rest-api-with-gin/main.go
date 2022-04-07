package main

import (
	"rest-api-with-gin/database"
	"rest-api-with-gin/routes"
)

func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}
