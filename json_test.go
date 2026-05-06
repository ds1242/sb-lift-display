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

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondWithError(rr, http.StatusBadRequest, "bad request")
	if rr.Code != http.StatusBadRequest {
		t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, rr.Code)
	}
	expected := `{"error":"bad request"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}


func TestRespondWithError_ServerError(t *testing.T) {
    rr := httptest.NewRecorder()
    RespondWithError(rr, http.StatusInternalServerError, "something broke")
    if rr.Code != http.StatusInternalServerError { 
		t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, rr.Code)
	}
    expected := `{"error":"something broke"}`
	if rr.Body.String() != expected {
		t.Errorf("Expected body %s, got %s", expected, rr.Body.String())
	}
}