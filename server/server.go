package main

import (
	"log"
	"net/http"
)

func textReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(status)
	w.Write([]byte(payload))
}

func errorReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	log.Printf("Error %d: %s", status, payload)
	http.Error(w, payload, status)
}

func newMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	return mux
}