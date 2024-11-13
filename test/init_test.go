package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"groupie-tracker/server"
)

func TestQuen(t *testing.T) {
	req := httptest.NewRequest("GET", "/artist?Artist=1", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(server.Artist)
	handler.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	expected := []string{"Freddie Mercury", "Brian May", "John Daecon", "Roger Meddows-Taylor", "Mike Grose", "Barry Mitchell", "Doug Fogie"}
	for _, val := range expected {
		if !contains(w.Body.String(), val) {
			t.Errorf("Expected body to contain %s, got %s", val, w.Body.String())
		}
	}
}

func TestGorillaz(t *testing.T) {
	req := httptest.NewRequest("GET", "/artist?Artist=39", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(server.Artist)
	handler.ServeHTTP(w, req)

	// Check if the response code is correct (200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	expected := "26-03-2001"

	if !contains(w.Body.String(), expected) {
		t.Errorf("Expected body to contain %s, got %s", expected, w.Body.String())
	}
}

func TestTravis(t *testing.T) {
	req := httptest.NewRequest("GET", "/artist?Artist=30", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(server.Artist)
	handler.ServeHTTP(w, req)

	// Check if the response code is correct (200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	expected := []string{"santiago-chile", "sao_paulo-brazil", "los_angeles-usa", "houston-usa", "atlanta-usa", "new_orleans-usa", "philadelphia-usa", "london-uk", "frauenfeld-switzerland", "turku-finland"}
	for _, val := range expected {
		if !contains(w.Body.String(), val) {
			t.Errorf("Expected body to contain %s, got %s", val, w.Body.String())
		}
	}
}

func TestFOO(t *testing.T) {
	req := httptest.NewRequest("GET", "/artist?Artist=51", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(server.Artist)
	handler.ServeHTTP(w, req)

	// Check if the response code is correct (200 OK)
	if w.Code != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", w.Code)
	}

	expected := []string{"Dave Grohl", "Nate Mendel", "Taylor Hawkins", "Chris Shiflett", "Chris Shiflett", "Pat Smear", "Rami Jaffee"}

	for _, val := range expected {
		if !contains(w.Body.String(), val) {
			t.Errorf("Expected body to contain %s, got %s", expected, w.Body.String())
		}
	}
}

func contains(body string, expected string) bool {
	return strings.Contains(body, expected)
}
