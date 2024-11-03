package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/server"
)

func main() {
	http.HandleFunc("/", server.Home)
	http.HandleFunc("/Artist",server.GetArtistByID)
	http.HandleFunc("/static/" , server.ServStatic)
	fmt.Println("the server is running on the port 5050 -->> http://localhost:5050")
	http.ListenAndServe(":5050", nil)
}
