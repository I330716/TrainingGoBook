// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	//It connects a handler function to incomingURLs thatbeg in with /,whichi s all URLs
	http.HandleFunc("/", handler) //each request calls handler
	//Start to listen on localhost:8000 and than call Serve with handler to handle the incomming requests
	// In our case this handler is the DefaultServeMux because the nil parrameter passed as handler
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err) // log the exit error
}

// handler echoes the PAth component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
