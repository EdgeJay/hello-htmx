package handlers

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/EdgeJay/hello-htmx/data"
	mw "github.com/EdgeJay/hello-htmx/middlewares"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
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

	// get user's todo list from in-memory store
	todoSvc := mw.GetTodoService(r)
	if todoSvc == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	todos := todoSvc.GetTodos(sessionID)
	log.Printf("User %s has %d todos\n", sessionID, len(todos))

	tpl, err := template.ParseFiles(
		"./htmx/index.html",
		"./htmx/todo.html",
	)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)

	// execute template and write to response
	if err := tpl.ExecuteTemplate(w, "index.html", map[string][]data.Todo{
		"Todos": todos,
	}); err != nil {
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
