package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/mafewo/meliexercise/config"
	"github.com/mafewo/meliexercise/toolkit"
)

//IPFilter filter request IP
func IPFilter(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if (config.IPFilter == "whitelist") || (config.IPFilter == "blacklist") {

			ipRemoteAddr := r.RemoteAddr
			ips := strings.Split(r.Header.Get("X-Forwarded-For"), ",")
			var evalIP string

			if ips[0] != "" {
				//ip fowarded
				ipRemoteAddr = ips[0]
			}

			evalIP = strings.Split(ipRemoteAddr, ":")[0]

			ex, _ := toolkit.InArray(evalIP, config.IPList)

			switch config.IPFilter {
			case "whitelist":
				if !ex {
					w.WriteHeader(http.StatusForbidden)
					log.Printf("Request rechazado: La IP no está en la Whitelist : %v", evalIP)
					return
				}
			case "blacklist":
				if ex {
					w.WriteHeader(http.StatusForbidden)
					log.Printf("Request rechazado: La ip está en la Blacklist : %v", evalIP)
					return
				}
			}

		}

		//next
		next.ServeHTTP(w, r)
	}
}
