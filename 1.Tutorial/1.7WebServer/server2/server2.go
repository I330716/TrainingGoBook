// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count uint

func main() {
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	//It connects a handler function to incomingURLs thatbeg in with /,whichi s all URLs
	http.HandleFunc("/", handler)      //each request calls handler
	http.HandleFunc("/count", counter) // add another handle function to the DefaultServeMux
	//Start to listen on localhost:8000 and than call Serve with handler to handle the incomming requests
	// In our case this handler is the DefaultServeMux because the nil parrameter passed as handler
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err) // log the exit error
}

// handler echoes the PAth component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	/*Behind the scenes, the server runs the handler for each incomingrequest in a sep arategoroutine so
	that it can serve multiple requests simultaneously.How ever, if two concur rentrequests try to
	update count at the same time, it might not be incremented consistently.
	There for we are using mutex to avoid race condition */
	mu.Lock()
	count++
	mu.Unlock()
	//writes to the response
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	mu.Lock()                           // avoid race condition
	fmt.Fprintf(w, "Count %d\n", count) // write to the response stream
	mu.Unlock()
}
