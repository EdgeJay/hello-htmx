package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
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

	http.ServeFile(w, r, "./htmx/index.html")
}
