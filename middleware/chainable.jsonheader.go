package middleware

import "net/http"

//JSONHeader chainable middleware example
func JSONHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "aplication/json")
		next.ServeHTTP(w, r)
	}
}
