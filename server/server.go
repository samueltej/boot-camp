package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func jsonReply(w http.ResponseWriter, r *http.Request, status int, payload *todoResponse) {
	data, err := json.Marshal(payload)
	if err != nil {
		errorReply(w, r, http.StatusInternalServerError, err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

func textReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)
	w.Write([]byte(payload))
}

func errorReply(w http.ResponseWriter, r *http.Request, status int, payload string) {
	log.Printf("Error %d: %s", status, payload)
	http.Error(w, payload, status)
}

func newMux(dataFile string) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.Handle("/todo", router(dataFile))
	mux.Handle("/todo/", router(dataFile))
	return mux
}
