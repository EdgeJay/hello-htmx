package services

import (
	"github.com/EdgeJay/hello-htmx/data"
)

type TodoService struct {
	UserTodos map[string][]data.Todo
}

func NewTodoService() *TodoService {
	return &TodoService{
		UserTodos: make(map[string][]data.Todo, 0),
	}
}

func (svc *TodoService) initTodos(sessionID string) {
	if _, exists := svc.UserTodos[sessionID]; !exists {
		svc.UserTodos[sessionID] = []data.Todo{}
	}
}

func (svc *TodoService) GetTodos(sessionID string) []data.Todo {
	svc.initTodos(sessionID)
	return svc.UserTodos[sessionID]
}

func (svc *TodoService) AddTodo(sessionID string, todo data.Todo) {
	svc.initTodos(sessionID)
	svc.UserTodos[sessionID] = append(svc.UserTodos[sessionID], todo)
}
