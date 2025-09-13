package handlers

import (
	"log"
	"net/http"

	mw "github.com/EdgeJay/hello-htmx/middlewares"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// read path value
	todoId := r.PathValue("id")
	if todoId == "" {
		http.Error(w, "Todo ID cannot be empty", http.StatusBadRequest)
		return
	}

	// get todo service from context
	todoSvc := mw.GetTodoService(r)
	if todoSvc == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// get session id
	sessionID := mw.GetSessionID(r)

	// In a real application, you would delete the todo item from a database here.
	// For this example, we'll just log it.
	log.Printf("API: Deleting todo item with ID: %s\n", todoId)

	todoSvc.DeleteTodo(sessionID, todoId)

	w.WriteHeader(http.StatusOK)
}
