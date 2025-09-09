package web

import (
	"html/template"
	"log"
	"net/http"
)

type Router struct {
	mux *http.ServeMux ``
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) SetupRoutes() {
	r.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		http.ServeFile(w, r, "./htmx/index.html")
	})

	// api
	r.mux.HandleFunc("POST /api/todo", func(w http.ResponseWriter, r *http.Request) {
		// read form value
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Failed to parse form", http.StatusBadRequest)
			return
		}
		todo := r.FormValue("todo")
		if todo == "" {
			http.Error(w, "Todo item cannot be empty", http.StatusBadRequest)
			return
		}

		// In a real application, you would save the todo item to a database here.
		// For this example, we'll just log it.
		log.Printf("API: Adding todo item: %s\n", todo)

		// read todo.html template file and replace {{todo}} with the actual todo item
		tpl, err := template.ParseFiles("./htmx/todo.html")
		if err != nil {
			http.Error(w, "Failed to load template", http.StatusInternalServerError)
			return
		}

		// execute template and write to response
		w.Header().Set("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)

		if err := tpl.Execute(w, map[string]string{"Todo": todo}); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	})

	r.mux.HandleFunc("DELETE /api/todo/{id}", func(w http.ResponseWriter, r *http.Request) {
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
	})

	// static assets
	r.mux.HandleFunc("GET /static/{path...}", func(w http.ResponseWriter, r *http.Request) {
		filePath := r.PathValue("path")
		log.Printf("Static asset: Received request for /static/%s\n", filePath)
		http.ServeFile(w, r, "./static/"+filePath)
	})
}

func (r *Router) Start() error {
	log.Println("Starting server on :8080")
	return http.ListenAndServe(":8080", r.mux)
}
