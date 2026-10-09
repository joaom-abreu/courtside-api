package handler

import (
	"context"
	"net/http"
	"time"
)

const healthCheckTimeout = 2 * time.Second

type DatabasePinger interface {
	Ping(ctx context.Context) error
}

type HealthHandler struct {
	db DatabasePinger
}

func NewHealthHandler(db DatabasePinger) *HealthHandler {
	return &HealthHandler{db: db}
}

type healthResponse struct {
	Status   string `json:"status"`
	Database string `json:"database"`
}

func (h *HealthHandler) Check(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), healthCheckTimeout)
	defer cancel()

	if err := h.db.Ping(ctx); err != nil {
		writeJSON(w, http.StatusServiceUnavailable, healthResponse{Status: "unavailable", Database: "down"})
		return
	}

	writeJSON(w, http.StatusOK, healthResponse{Status: "ok", Database: "up"})
}
