/*
The following script is a Go program that serves static files from a specified directory and listens on a specified port. It accepts two command-line flags:

    dir: Specifies the directory to serve static files from. The default is ./_site.
    port: Specifies the port to listen on. The default is 3000.

To use the program, run it from the command line and pass in the required flags. For example:

> go run main.go -dir=./public -port=8080

The program creates an HTTP file server using the http.FileServer function, which serves files from the specified directory. It then registers the file server
to handle requests on the root URL path. Finally, it starts listening on the specified port using the http.ListenAndServe function. If an error occurs, it logs
the error and exits.
*/
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
