package handlers

import (
	"log"
	"net/http"
)

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// read path value
	id := r.PathValue("id")
	if id == "" {
		http.Error(w, "Todo ID cannot be empty", http.StatusBadRequest)
		return
	}

	// In a real application, you would delete the todo item from a database here.
	// For this example, we'll just log it.
	log.Printf("API: Deleting todo item with ID: %s\n", id)

	w.WriteHeader(http.StatusNoContent)
}
