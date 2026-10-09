package handler

import "net/http"

func NewRouter(health *HealthHandler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", health.Check)

	return mux
}
