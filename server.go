package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define command-line flags
	dir := flag.String("dir", "./_site", "Directory to serve static files from")
	port := flag.String("port", "3000", "Port to listen on")

	// Parse command-line flags
	flag.Parse()

	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	log.Printf("Listening on :%s...", *port)
	err := http.ListenAndServe(":"+*port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
