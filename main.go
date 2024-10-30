package main

import (
	"groupie-tracker/server"
	"net/http"
)
 
func main(){
	http.HandleFunc("/", server.Home)
	http.ListenAndServe(":5050",nil)
}