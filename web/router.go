package web

import (
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
