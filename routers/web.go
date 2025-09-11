package routers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/EdgeJay/hello-htmx/handlers"
	"github.com/EdgeJay/hello-htmx/middlewares"
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
	r.mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		sessionID := ""

		// get session id
		cookie, err := r.Cookie("session_id")
		if err != nil {
			// start cookie-based session
			sessionID = uuid.New().String() // Generate a unique session ID

			cookie := &http.Cookie{
				Name:     "session_id",
				Value:    sessionID,
				Path:     "/",
				Expires:  time.Now().Add(7 * 24 * time.Hour), // Set cookie expiry
				HttpOnly: true,                               // Prevent JavaScript access
				SameSite: http.SameSiteLaxMode,               // Prevent CSRF attacks
			}

			http.SetCookie(w, cookie)
		} else {
			sessionID = cookie.Value
			log.Printf("Session ID: %s\n", sessionID)
		}

		http.ServeFile(w, r, "./htmx/index.html")
	})

	// api
	r.mux.HandleFunc("POST /api/todo", handlers.PostTodo)

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
	todoService := services.NewTodoService()
	return http.ListenAndServe(":8080", middlewares.WithServices(todoService)(r.mux))
}
