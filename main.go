package main

import (
	"encoding/json"
	"net/http"
	"os"
	"time"
)

func envOr(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}

func main() {
	port := envOr("PORT", "8080")
	appid := envOr("ALB_APP_ID", "APP-CHECKOUT-004")
	envName := envOr("ALB_ENV", "prod-mumbai")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(map[string]any{
			"service": "checkout-service",
			"appid":   appid,
			"env":     envName,
			"time":    time.Now().UTC().Format(time.RFC3339),
		})
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok\n"))
	})

	srv := &http.Server{
		Addr:              ":" + port,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		os.Exit(1)
	}
}
