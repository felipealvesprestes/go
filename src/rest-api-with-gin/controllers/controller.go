package controllers

import (
	"rest-api-with-gin/models"

	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"data": "Fala " + nome + "! Mais um pateta pro Go. Welcome bro!",
	})
}
