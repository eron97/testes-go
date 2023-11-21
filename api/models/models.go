package models

// Task representa uma tarefa na To-Do List.
type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title" binding:"required,min=3"`
	Done  bool   `json:"done"`
}

// TaskList representa a lista de tarefas.
// var TaskList []Task

/*
var TaskList = []Task{
	{ID: 1, Title: "Fazer compras", Done: true},
	{ID: 2, Title: "Estudar Go", Done: false},
}
*/

var TaskList = []Task{}
