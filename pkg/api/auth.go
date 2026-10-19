package api

import (
	"net/http"
	"os"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		password := os.Getenv("TODO_PASSWORD")

		if password != "" {
			cookie, err := r.Cookie("token")
			if err != nil || !validToken(cookie.Value, password) {
				writeJSON(w, http.StatusUnauthorized, map[string]string{
					"error": "Требуется аутентификация",
				})
				return
			}
		}

		next(w, r)
	}
}
