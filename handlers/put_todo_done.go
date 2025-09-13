package handlers

import (
	"html/template"
	"log"
	"net/http"
	"reflect"

	mw "github.com/EdgeJay/hello-htmx/middlewares"
)

func PutTodoDone(w http.ResponseWriter, r *http.Request) {
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

	log.Printf("API: Toggling todo item with ID: %s\n", todoId)

	todo := todoSvc.ToggleTodo(sessionID, todoId)
	if reflect.ValueOf(todo).IsZero() {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	// read todo.html template file
	tpl, err := template.ParseFiles("./htmx/todo.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// execute template and write to response, replacing {{.ID}} {{.Item}} with Todo
	if err := tpl.Execute(w, todo); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
