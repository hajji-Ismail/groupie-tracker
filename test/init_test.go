package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/server"
)



func HomeTest(t *testing.T) {
	serving := httptest.NewServer(http.HandlerFunc(server.Home))

	req := httptest.NewRequest("GET", serving.URL, nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status code %d got %d", http.StatusOK, res.StatusCode)
	}
	if res.Request.Method != http.MethodGet {
		t.Errorf("expected method %s got %s", http.MethodGet, res.Request.Method)
	}
}


func TestArtistHandler(t *testing.T) {
	// Replace the real Fetchdetails function with the mock

	// Create a new request to test the Artist handler
	req := httptest.NewRequest("GET", "/artist?Artist=1", nil) // artist ID is 1
	w := httptest.NewRecorder()

	// Call the Artist handler
	handler := http.HandlerFunc(server.Artist)
	handler.ServeHTTP(w, req)

	// Check if the response code is correct (200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	
	expected := "Queen" 
	if !contains(w.Body.String(), expected) {
		t.Errorf("Expected body to contain %s, got %s", expected, w.Body.String())
	}
}

// Utility function to check if the response body contains a substring
func contains(body string, expected string) bool {
	return strings.Contains(body, expected)
}
