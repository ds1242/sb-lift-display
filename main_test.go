package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandleGetLifts_Authorized(t *testing.T) {
	cache := &Cache{
		liftStatus: &LiftStatus{Lift: []Lift{{ID: "1"}}},
		fetchedAt:  time.Now(),
	}
	handler := handleGetLifts(cache, "secret-key")
	req := httptest.NewRequest("GET", "/api/lifts", nil)
	req.Header.Set("X-API-Key", "secret-key")
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected 200, got %d", rr.Code)
	}
}

func TestHandleGetLifts_Unauthorized(t *testing.T) {
	handler := handleGetLifts(&Cache{}, "secret-key")
	req := httptest.NewRequest("GET", "/api/lifts", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusUnauthorized {
		t.Errorf("Expected 401, got %d", rr.Code)
	}
}
