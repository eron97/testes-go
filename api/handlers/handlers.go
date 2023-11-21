package handlers

import (
	"net/http"
	"strconv"

	"github.com/eron97/testes-go.git/api/models"
	"github.com/gin-gonic/gin"
)

// ListTasks retorna a lista de tarefas.
func ListTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"tarefas": models.TaskList})
}

/* ao fazer um GET em ListTask é devolvido:

[
    {
        "id": 1,
        "title": "Fazer compras",
        "done": false
    },
    {
        "id": 2,
        "title": "Estudar Go",
        "done": true
    }
]

Ou seja, de fato uma array/lista de structs do tipo TaskList

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

*/

// AddTask adiciona uma nova tarefa à lista.
func AddTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTask.ID = len(models.TaskList) + 1
	models.TaskList = append(models.TaskList, newTask)

	c.JSON(http.StatusCreated, newTask)
}

// CompleteTask marca uma tarefa como concluída.
func CompleteTask(c *gin.Context) {
	taskID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	for i := range models.TaskList {
		if models.TaskList[i].ID == taskID {
			models.TaskList[i].Done = true
			c.JSON(http.StatusOK, models.TaskList[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}
