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
			{ID: uuid.New().String(), Item: "Learn Go", Done: true},
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

func (svc *TodoService) UpdateTodo(sessionID string, todoId string, newItem string) data.Todo {
	svc.initTodos(sessionID)
	for idx, todo := range svc.UserTodos[sessionID] {
		if todo.ID == todoId {
			todo.Item = newItem
			svc.UserTodos[sessionID][idx] = todo
			return todo
		}
	}
	return data.Todo{}
}

func (svc *TodoService) ToggleTodo(sessionID string, todoId string) data.Todo {
	svc.initTodos(sessionID)
	for idx, todo := range svc.UserTodos[sessionID] {
		if todo.ID == todoId {
			todo.Done = !todo.Done
			svc.UserTodos[sessionID][idx] = todo
			return todo
		}
	}
	return data.Todo{}
}

func (svc *TodoService) DeleteTodo(sessionID string, todoId string) {
	todos := svc.GetTodos(sessionID)
	for idx, t := range todos {
		if t.ID == todoId {
			svc.UserTodos[sessionID] = append(todos[:idx], todos[idx+1:]...)
			break
		}
	}
}
