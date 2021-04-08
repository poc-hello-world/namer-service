package main

import (
	"log"
	"net/http"
)

// BuildID is set by CI
var BuildID string

func setServiceHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Reply-Service", "namer-service")
		if BuildID == "" {
			w.Header().Set("X-Version", "dev")
		} else {
			w.Header().Set("X-Version", BuildID)
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Accept-Language, Content-Type, X-Version, X-Reply-Service")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func main() {
	globalmux := http.NewServeMux()

	globalmux.HandleFunc("/", NameHandler)
	globalmux.HandleFunc("/status/alive", AliveHandler)
	globalmux.HandleFunc("/status/ready", HealthyHandler)

	log.Fatal(http.ListenAndServe(":5003", setServiceHeader(globalmux)))
}
