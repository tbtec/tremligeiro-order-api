package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

func TestNewHTTPServerRoutes(t *testing.T) {
	// Mocks mínimos para container e config
	mockContainer := &container.Container{}
	mockConfig := env.Config{Port: 8080}

	server := New(mockContainer, mockConfig)
	app := server.Server

	tests := []struct {
		method string
		route  string
		status int
	}{
		{"GET", "/live", http.StatusOK},
		{"POST", "/api/v1/order", http.StatusBadRequest}, // Espera erro por falta de body válido
		{"GET", "/api/v1/order", http.StatusBadRequest},
		{"GET", "/api/v1/order/123", http.StatusBadRequest},
		{"PUT", "/api/v1/order/123", http.StatusBadRequest},
		{"POST", "/api/v1/order/123/checkout", http.StatusBadRequest},
		{"GET", "/not-found", http.StatusNotFound},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(tt.method, tt.route, nil)
		resp, err := app.Test(req, -1)
		assert.NoError(t, err, tt.route)
		assert.Equal(t, tt.status, resp.StatusCode, tt.route)
	}
}
