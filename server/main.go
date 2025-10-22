package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	
	host := flag.String("h", "localhost", "host fot default")
	port := flag.Int("p", 8080, "port of server")
	flag.Parse()

	addr := fmt.Sprintf("%s:%d",*host, *port)

	server := &http.Server{
		Addr: addr,
		Handler: newMux(),
	}

	log.Fatal(server.ListenAndServe())
}