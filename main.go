package main

import (
	"net/http"

	"groupie-tracker/server"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.ListenAndServe(":5050", nil)
}
