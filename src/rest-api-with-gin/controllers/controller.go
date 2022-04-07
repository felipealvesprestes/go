package controllers

import "github.com/gin-gonic/gin"

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "123",
		"nome": "Felipe Prestes",
	})
}

func Hello(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"data": "Fala " + nome + "! Mais um pateta pro Go. Welcome bro!",
	})
}
