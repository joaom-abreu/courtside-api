package handler

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type fakePinger struct {
	err error
}

func (f fakePinger) Ping(_ context.Context) error {
	return f.err
}

func TestHealthHandler_Check(t *testing.T) {
	tests := []struct {
		name         string
		pingErr      error
		wantCode     int
		wantStatus   string
		wantDatabase string
	}{
		{
			name:         "database up",
			pingErr:      nil,
			wantCode:     http.StatusOK,
			wantStatus:   "ok",
			wantDatabase: "up",
		},
		{
			name:         "database down",
			pingErr:      errors.New("connection refused"),
			wantCode:     http.StatusServiceUnavailable,
			wantStatus:   "unavailable",
			wantDatabase: "down",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			router := NewRouter(NewHealthHandler(fakePinger{err: tt.pingErr}))
			req := httptest.NewRequest(http.MethodGet, "/health", nil)
			rec := httptest.NewRecorder()

			router.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("expected status code %d, got %d", tt.wantCode, rec.Code)
			}

			var body healthResponse
			if err := json.NewDecoder(rec.Body).Decode(&body); err != nil {
				t.Fatalf("could not decode response: %v", err)
			}

			if body.Status != tt.wantStatus {
				t.Errorf("expected status %q, got %q", tt.wantStatus, body.Status)
			}
			if body.Database != tt.wantDatabase {
				t.Errorf("expected database %q, got %q", tt.wantDatabase, body.Database)
			}
		})
	}
}
