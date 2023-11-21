package handlers_test

// ok  - coverage: 100.0% of statements

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/eron97/testes-go.git/api/handlers"
	"github.com/eron97/testes-go.git/api/models"
	"github.com/gin-gonic/gin"
)

// teste de unidade para testar a comunicação entre rota e handler | considerado um teste de integração também
// -> A rota acionada por um request consegue obter uma comunicação com o handler e conseguir uma resposta de sucesso? SIM
// Para o teste mockamos um roteador Gin isolado na função de teste (não há necessidade nem mesmo de acionar um localhost para testar se funciona.)

// Testes de Cobertura para listTasks

func TestListTasks(t *testing.T) {
	tests := []struct {
		name      string
		listTasks []models.Task
		want      string
	}{
		{
			name:      "Teste: Sem tarefas preenchidas",
			listTasks: []models.Task{},
			want:      `{"tarefas":[]}`,
		},
		{
			name: "Teste: Com tarefas preenchidas",
			listTasks: []models.Task{
				{ID: 1, Title: "Fazer compras", Done: true},
				{ID: 2, Title: "Estudar Go", Done: false},
			},
			want: `{"tarefas":[{"id":1,"title":"Fazer compras","done":true},{"id":2,"title":"Estudar Go","done":false}]}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.GET("/tasks", handlers.ListTasks)

			models.TaskList = []models.Task{}
			newTasks := tt.listTasks
			models.TaskList = append(models.TaskList, newTasks...)

			req, err := http.NewRequest(http.MethodGet, "/tasks", nil)
			if err != nil {
				t.Fatal(err)
			}

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			corpo, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Erro interno no servidor: %d", http.StatusInternalServerError)
				return
			}

			corpoString := string(corpo)

			if corpoString != tt.want {
				t.Errorf("Erro | %s: Esperado código %s, obtido %s", tt.name, tt.want, corpoString)
			}
		})
	}
}

func TestAddTask(t *testing.T) {
	tests := []struct {
		name        string
		listTasks   []models.Task
		requestBody string
		want        string
	}{
		{
			name:        "Teste: Lista vazia",
			listTasks:   []models.Task{},
			requestBody: `{"title":"Nova Tarefa"}`,
			want:        `{"id":1,"title":"Nova Tarefa","done":false}`,
		},
		{
			name: "Teste: Lista preenchida",
			listTasks: []models.Task{
				{ID: 1, Title: "Fazer compras", Done: true},
			},
			requestBody: `{"title":"Comprar novo computador"}`,
			want:        `{"id":2,"title":"Comprar novo computador","done":false}`,
		},
		{
			name:        "Teste: Corpo de requisição inválido",
			listTasks:   []models.Task{},
			requestBody: "{}",
			want:        `{"error":"Key: 'Task.Title' Error:Field validation for 'Title' failed on the 'required' tag"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.POST("/tasks", handlers.AddTask)

			models.TaskList = []models.Task{}
			newTasks := tt.listTasks
			models.TaskList = append(models.TaskList, newTasks...)

			requestBody := []byte(tt.requestBody)

			req, err := http.NewRequest(http.MethodPost, "/tasks", bytes.NewBuffer(requestBody))
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			corpo, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Erro interno no servidor: %d", http.StatusInternalServerError)
				return
			}

			corpoString := string(corpo)

			if corpoString != tt.want {
				t.Errorf("Resposta inesperada. Esperado: %+v, Obtido: %+v", tt.want, corpoString)
				return
			}
		})
	}
}

func TestCompleteTask(t *testing.T) {
	tests := []struct {
		name      string
		listTasks []models.Task
		url       string
		want      string
	}{
		{
			name:      "Teste: Lista vazia",
			listTasks: []models.Task{},
			url:       "/tasks/2/complete",
			want:      `{"error":"Task not found"}`,
		},
		{
			name: "Teste: Lista preenchida",
			listTasks: []models.Task{
				{ID: 1, Title: "Fazer compras", Done: true},
				{ID: 2, Title: "Estudar Go", Done: false},
			},
			url:  "/tasks/2/complete",
			want: `{"id":2,"title":"Estudar Go","done":true}`,
		},
		{
			name:      "Teste: ID de tarefa inválido",
			listTasks: []models.Task{},
			url:       "/tasks/invalidID/complete",
			want:      `{"error":"Invalid task ID"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := gin.Default()
			r.PUT("/tasks/:id/complete", handlers.CompleteTask)

			models.TaskList = []models.Task{}
			newTasks := tt.listTasks
			models.TaskList = append(models.TaskList, newTasks...)

			req, err := http.NewRequest(http.MethodPut, tt.url, nil)
			if err != nil {
				t.Fatal(err)
			}

			req.Header.Set("Content-Type", "application/json")

			resp := httptest.NewRecorder()
			r.ServeHTTP(resp, req)

			corpo, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("Erro interno no servidor: %d", http.StatusInternalServerError)
				return
			}

			corpoString := string(corpo)

			if corpoString != tt.want {
				t.Errorf("Resposta inesperada. Esperado: %+v, Obtido: %+v", tt.want, corpoString)
				return
			}
		})
	}
}
