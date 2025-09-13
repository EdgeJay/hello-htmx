package services

import (
	"github.com/EdgeJay/hello-htmx/data"
	"github.com/google/uuid"
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
		svc.UserTodos[sessionID] = []data.Todo{
			{ID: uuid.New().String(), Item: "Learn Go", Done: false},
			{ID: uuid.New().String(), Item: "Learn HTMX", Done: false},
			{ID: uuid.New().String(), Item: "Build something awesome!", Done: false},
		}
	}
}

func (svc *TodoService) GetTodos(sessionID string) []data.Todo {
	svc.initTodos(sessionID)
	return svc.UserTodos[sessionID]
}

func (svc *TodoService) AddTodo(sessionID string, todo string, done bool) data.Todo {
	svc.initTodos(sessionID)
	todoItem := data.Todo{
		ID:   uuid.New().String(),
		Item: todo,
		Done: done,
	}
	svc.UserTodos[sessionID] = append(svc.UserTodos[sessionID], todoItem)
	return todoItem
}
