package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHealthRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		name       string
		method     string
		url        string
		wantStatus int
	}{
		{
			name:       "GET /health returns 200",
			method:     http.MethodGet,
			url:        "/health",
			wantStatus: http.StatusOK,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tc.method, tc.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Errorf("status: got %d, want %d", w.Code, tc.wantStatus)
			}
		})
	}
}

func TestHomeRoute(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		name         string
		wantStatus   int
		wantContains []string
		wantAbsent   []string
	}{
		{
			name:       "GET / returns 200 with full landing page",
			wantStatus: http.StatusOK,
			wantContains: []string{
				"AgentClinic",
				"Relief for the Overworked AI",
				"Book a Session",
				"Describe Your Ailment",
				"Choose a Therapy",
				"Book Your Session",
				"2026 AgentClinic",
				"htmx",
			},
		},
		{
			name:         "GET / includes HTML title tag with site name",
			wantStatus:   http.StatusOK,
			wantContains: []string{"<title>AgentClinic</title>"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/", nil)
			router.ServeHTTP(w, req)

			if w.Code != tc.wantStatus {
				t.Errorf("status: got %d, want %d", w.Code, tc.wantStatus)
			}

			body := w.Body.String()
			for _, want := range tc.wantContains {
				if !strings.Contains(body, want) {
					t.Errorf("body missing %q", want)
				}
			}
			for _, absent := range tc.wantAbsent {
				if strings.Contains(body, absent) {
					t.Errorf("body should not contain %q", absent)
				}
			}
		})
	}
}

func TestUnknownRouteReturns404(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := setupRouter()

	tests := []struct {
		name string
		url  string
	}{
		{name: "missing page", url: "/does-not-exist"},
		{name: "admin path", url: "/admin"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, tc.url, nil)
			router.ServeHTTP(w, req)

			if w.Code != http.StatusNotFound {
				t.Errorf("%s: status: got %d, want 404", tc.url, w.Code)
			}
		})
	}
}
