package main

import (
	"log"
	"net/http"
	"server_go/handlers"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/home", handlers.HomeHandler{})
	mux.Handle("/about", http.HandlerFunc(handlers.AboutHandler))
	mux.HandleFunc("/help", handlers.HelpHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
