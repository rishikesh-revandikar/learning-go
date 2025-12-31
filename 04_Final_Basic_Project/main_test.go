package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rishikesh-revandikar/learning-g/04_Final_Basic_Project/internal/routes"
)

func TestHealthHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rr := httptest.NewRecorder()
	
	handler := http.HandlerFunc(routes.HealthHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v want %v", status, http.StatusOK)
	}

	if rr.Body.String() != "OK" {
		t.Errorf("Unexpected body: got %v want OK", rr.Body.String())
	}
}