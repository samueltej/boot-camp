package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello World!!"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)

	port := flag.Int("p", 8080, "port of server")
	flag.Parse()

	addr := fmt.Sprintf(":%d", *port)

	log.Fatal(http.ListenAndServe(addr, mux))
}