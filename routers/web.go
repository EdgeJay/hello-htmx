package routers

import (
	"log"
	"net/http"

	"github.com/EdgeJay/hello-htmx/handlers"
	mw "github.com/EdgeJay/hello-htmx/middlewares"
	"github.com/EdgeJay/hello-htmx/services"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) SetupRoutes() {
	r.mux.HandleFunc("GET /", handlers.GetIndex)

	// api
	r.mux.Handle("POST /api/todo", mw.CheckSession(http.HandlerFunc(handlers.PostTodo)))
	r.mux.Handle("PUT /api/todo/{id}/done", mw.CheckSession(http.HandlerFunc(handlers.PutTodoDone)))
	r.mux.Handle("DELETE /api/todo/{id}", mw.CheckSession(http.HandlerFunc(handlers.DeleteTodo)))

	// static assets
	r.mux.HandleFunc("GET /static/{path...}", handlers.GetStatic)
}

func (r *Router) Start() error {
	log.Println("Starting server on :8080")
	todoService := services.NewTodoService()
	return http.ListenAndServe(":8080", mw.WithTodoService(todoService)(r.mux))
}
