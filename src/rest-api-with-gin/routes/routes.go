package routes

import (
	"rest-api-with-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/hello/:nome", controllers.Hello)
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.POST("/alunos", controllers.CriaAluno)
	r.GET("/alunos/:id", controllers.ExibeAluno)
	r.Run(":8000")
}
