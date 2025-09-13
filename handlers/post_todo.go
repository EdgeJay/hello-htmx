package handlers

import (
	"html/template"
	"log"
	"net/http"

	mw "github.com/EdgeJay/hello-htmx/middlewares"
)

func PostTodo(w http.ResponseWriter, r *http.Request) {
	// get todo service from context
	todoSvc := mw.GetTodoService(r)
	if todoSvc == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// get session id
	sessionID := mw.GetSessionID(r)

	// read form value
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	todoItem := r.FormValue("todo")
	if todoItem == "" {
		http.Error(w, "Todo item cannot be empty", http.StatusBadRequest)
		return
	}

	log.Printf("API: Adding todo item: %s\n", todoItem)

	// read todo.html template file
	tpl, err := template.ParseFiles("./htmx/todo.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// save into in-memory store
	todo := todoSvc.AddTodo(sessionID, todoItem, false)

	// execute template and write to response, replacing {{.ID}} {{.Item}} with Todo
	if err := tpl.Execute(w, todo); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	} else {

	}
}
