package main

import (
	"github.com/eron97/testes-go.git/api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Rotas
	r.GET("/tasks", handlers.ListTasks)
	r.POST("/tasks", handlers.AddTask)
	r.PUT("/tasks/:id/complete", handlers.CompleteTask)

	// Inicializa o servidor
	r.Run(":8080")
}
