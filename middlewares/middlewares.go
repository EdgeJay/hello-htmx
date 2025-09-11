package middlewares

import (
	"context"
	"net/http"

	"github.com/EdgeJay/hello-htmx/services"
)

type ServiceKey string

const todoServiceKey = ServiceKey("todoService")

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
