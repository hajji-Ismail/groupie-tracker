package server

import (
	"html/template"
	"log"
	"net/http"

	"groupie-tracker/fetching"
)

type Parse struct {
	Index     *template.Template
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
