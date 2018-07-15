package middleware

import (
	"net/http"
	"strings"

	"github.com/mafewo/meliexercise/config"
)

//CORS general middleware
func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		AllowOrigin := strings.Join(config.AllowOrigin, ",")
		AllowMethods := strings.Join(config.AllowMethods, ",")
		AllowHeaders := strings.Join(config.AllowHeaders, ",")

		w.Header().Set("Access-Control-Allow-Origin", AllowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", AllowMethods)
		w.Header().Set("Access-Control-Allow-Headers", AllowHeaders)
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
