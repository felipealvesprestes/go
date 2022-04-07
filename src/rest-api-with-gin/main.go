package main

import (
	"rest-api-with-gin/models"
	"rest-api-with-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{Id: 1, Nome: "Felipe Prestes", CPF: "034.425.678-45", RG: "9845362789"},
		{Id: 1, Nome: "Gabriela Krauzer", CPF: "567.234.456-12", RG: "46578988746"},
		{Id: 1, Nome: "Mariana Prestes", CPF: "234.345.456-43", RG: "09446788934"},
	}
	routes.HandleRequests()
}
