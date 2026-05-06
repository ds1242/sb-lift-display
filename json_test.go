package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRespondWithJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondWithJSON(rr, http.StatusOK, map[string]string{"key": "value"})
	if rr.Code != 200 {
		t.Errorf("Result was incorrect, got %d, want %d\n", rr.Code, 200)
	}
	if rr.Header().Get("Content-Type") != "application/json" {
		t.Errorf("Expected Content-Type application/json, got %s\n", rr.Header().Get("Content-Type"))
	}
	expected := `{"key":"value"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s\n", expected, rr.Body.String())
	}

}

func TestRespondWithErrorTable(t *testing.T) {
	tests := []struct {
		name    string
		message string
		body    string
		want    int
	}{
		{"internal service error", "something broke", `{"error":"something broke"}`, http.StatusInternalServerError},
		{"bad request error", "bad request", `{"error":"bad request"}`, http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			RespondWithError(rr, tt.want, tt.message)
			if rr.Code != tt.want {
				t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
			}
			if rr.Body.String() != tt.body {
				t.Errorf("Expected body %s, got %s", tt.message, rr.Body.String())
			}
		})
	}
}
