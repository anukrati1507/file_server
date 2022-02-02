package main

import (
	"net/http"
	"fmt"
)

func setupRoutes() {
	http.HandleFunc("/put", Put)
	http.HandleFunc("/get", Get)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/list", List)
	http.HandleFunc("/rename", Rename)
	http.ListenAndServe(":5500", nil)
}

func main() {
	fmt.Println("File - Server")
	setupRoutes()
}