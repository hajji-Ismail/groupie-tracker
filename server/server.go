package server

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"groupie-tracker/fetching"
)

type Parse struct {
	Index  *template.Template
	Artist *template.Template

	ErrorTemp *template.Template
}

var parsing Parse

func init() {
	// Get the current working directory (project root directory)
	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	// Build the absolute path using the working directory
	Index, err := template.ParseFiles(wd + "/template/index.html")
	if err != nil {
		log.Fatalf("Error parsing index.html: %v", err)
	}
	parsing.Index = Index

	Artist, err := template.ParseFiles(wd + "/template/artist.html")
	if err != nil {
		log.Fatalf("Error parsing artist.html: %v", err)
	}
	parsing.Artist = Artist

	ErrorTemp, err := template.ParseFiles(wd + "/template/error.html")
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
