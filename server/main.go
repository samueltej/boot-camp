package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	dataFile := flag.String("f", "datafile.json", "todo data file")
	host := flag.String("h", "localhost", "server host")
	port := flag.Int("p", 8080, "server port")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d", *host, *port)

	server := &http.Server{
		Addr:    addr,
		Handler: newMux(*dataFile),
	}

	log.Printf("Server listening on http://%s", addr)
	log.Fatal(server.ListenAndServe())
}
