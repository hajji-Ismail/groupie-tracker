package server

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"

	"groupie-tracker/fetching"
	"groupie-tracker/models"
)

type Parse struct {
	Index *template.Template

	ErrorTemp *template.Template
}

var parsing Parse

func init() {
	Index, err := template.ParseFiles("template/index.html")
	if err != nil {
		log.Fatal("I can't parse the index.html file")
		return
	}
	parsing.Index = Index

	ErrorTemp, err := template.ParseFiles("template/error.html")
	if err != nil {
		log.Fatal("I can't parse the error.html file")
		return
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
	}
}

func GetArtistByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		parsing.ErrorTemp.Execute(w, "Method Not Allowed")
		return
	}

	// Extract the artist ID from the URL
	artistIDStr := r.URL.Path[len("/artist/"):] // Assuming the URL is like /artist/{id}
	artistID, err := strconv.Atoi(artistIDStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		parsing.ErrorTemp.Execute(w, "Invalid Artist ID")
		return
	}

	// Fetch the artist data using the artist ID
	var artist *models.Artist
	artists := fetching.Artists // Assuming this is of type *[]models.Artist

	// Dereference the pointer to access the slice
	for _, a := range *artists {
		if a.Id == artistID { // Assuming each artist has an ID field
			artist = &a
			break
		}
	}

	if artist == nil {
		w.WriteHeader(http.StatusNotFound)
		parsing.ErrorTemp.Execute(w, "Artist Not Found")
		return
	}

	// Render the artist data using a template (you may need to create this template)
	err = parsing.Index.Execute(w, artist) // Use a different template if necessary
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		parsing.ErrorTemp.Execute(w, "Internal Server Error")
	}
}

func ServStatic(w http.ResponseWriter, r *http.Request) {
	// Check if the request is for the CSS directory itself
	if r.URL.Path == "/css/" || r.URL.Path == "/css" {
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
