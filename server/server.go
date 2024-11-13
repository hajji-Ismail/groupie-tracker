package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"groupie-tracker/fetching"
)

type Parse struct {
	Index  *template.Template
	Artist *template.Template

	ErrorTemp *template.Template
}

var parsing Parse

func init() {
	// Get the current working directory
	workingDir, err := os.Getwd()
	if strings.HasSuffix(workingDir, "/test") {
		workingDir = strings.TrimSuffix(workingDir, "/test")
	}
	if err != nil {
		log.Fatalf("Error getting current working directory: %v", err)
	}

	// Construct the absolute path for index.html
	indexPath := filepath.Join(workingDir, "template", "index.html")
	Index, err := template.ParseFiles(indexPath)
	if err != nil {
		log.Fatalf("Error parsing index.html: %v", err)
	}
	parsing.Index = Index

	// Construct the absolute path for artist.html
	artistPath := filepath.Join(workingDir, "template", "artist.html")
	Artist, err := template.ParseFiles(artistPath)
	if err != nil {
		log.Fatalf("Error parsing artist.html: %v", err)
	}
	parsing.Artist = Artist

	// Construct the absolute path for error.html
	errorPath := filepath.Join(workingDir, "template", "error.html")
	ErrorTemp, err := template.ParseFiles(errorPath)
	if err != nil {
		log.Fatalf("Error parsing error.html: %v", err)
	}
	parsing.ErrorTemp = ErrorTemp
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsing.ErrorTemp.Execute(w, "Method Not Allowed")
		return // Return after handling non-GET request
	}
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		parsing.ErrorTemp.Execute(w, "notFound")
		return // Return after handling 404 error
	}
	data := fetching.Artists
	// Use the global data variable instead of declaring a new one
	err := parsing.Index.Execute(w, data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsing.ErrorTemp.Execute(w, "Internal Server Error")
		return
	}
}

func Artist(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsing.ErrorTemp.Execute(w, "Method Not Allowed")
		return
	}

	// Extract the artist ID from the URL
	artistIDStr := r.URL.Query().Get("Artist")
	if artistIDStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		parsing.ErrorTemp.Execute(w, "Artist ID is required")
		return
	}
	Artist, err := fetching.Fetchdetails(artistIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parsing.ErrorTemp.Execute(w, "Bad Request")
		return
	}

	errr := parsing.Artist.Execute(w, Artist)
	if errr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsing.ErrorTemp.Execute(w, "Internal Server Error")
		return
	}
}

func ServStatic(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for the CSS directory itself
	if r.URL.Path == "/static/" || r.URL.Path == "/static" {
		w.WriteHeader(http.StatusNotFound)
		parsing.ErrorTemp.Execute(w, "NOT FOUND")
		return
	}
	_, err := os.Stat(r.URL.Path[1:])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		parsing.ErrorTemp.Execute(w, "NOT FOUND")
		return
	}

	// Serve the CSS file if it's a valid request
	http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
}
