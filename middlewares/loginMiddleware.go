package middlewares

import (
	"encoding/json"
	"fmt"
	"go-todos/literals"
	"go-todos/models"
	"net/http"
)

func LoginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		session, ok := ctx.Value(literals.RequestSessionUserKey).(*models.Session)
		if !ok || session.UserID == "" {
			fmt.Println("login required")
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(401)
			json.NewEncoder(w).Encode("Login required")
			return
		}
		next.ServeHTTP(w, r)
	})
}
