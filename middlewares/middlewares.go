package middlewares

import (
	"context"
	"net/http"

	"github.com/EdgeJay/hello-htmx/services"
)

type ServiceKey string

type SessionIdKey string

const todoServiceKey = ServiceKey("todoService")

const sessionIdKey = SessionIdKey("session_id")

func CheckSession(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get session id
		cookie, err := r.Cookie("session_id")
		if err != nil {
			http.Error(w, "Session not found", http.StatusUnauthorized)
			return
		}

		// add session id to context
		ctx := context.WithValue(r.Context(), sessionIdKey, cookie.Value)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetSessionID(r *http.Request) string {
	if sid, ok := r.Context().Value(sessionIdKey).(string); ok {
		return sid
	}
	return ""
}

func WithTodoService(todoSvc *services.TodoService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Inject services into request context
			ctx := context.WithValue(r.Context(), todoServiceKey, todoSvc)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func GetTodoService(r *http.Request) *services.TodoService {
	if svc, ok := r.Context().Value(todoServiceKey).(*services.TodoService); ok {
		return svc
	}
	return nil
}
