//HTTP Server for debuging purpose
package main

import (
	"fmt"
	"log"
	"net/http"
)

//handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL.String(), r.Proto)
	//print headers to the response
	for header, value := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", header, value)
	}
	//print the host
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	//print remote address
	fmt.Fprintf(w, "Remote address = %q\n", r.RemoteAddr)
	//parse the form arguments /index.htlm?pretty=true
	if err := r.ParseForm(); err != nil {
		//err is visible only in the scope of the if's block
		log.Print(err)
	}
	//now print the form
	for form, value := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", form, value)
	}
}

func main() {
	//HandleFunc registers the handler function for the given pattern in the DefaultServeMux.
	//It connects a handler function to incomingURLs thatbeg in with /,whichi s all URLs
	http.HandleFunc("/", handler) //each request calls handler
	//Start to listen on localhost:8000 and than call Serve with handler to handle the incomming requests
	// In our case this handler is the DefaultServeMux because the nil parrameter passed as handler
	err := http.ListenAndServe("localhost:8000", nil)
	log.Fatal(err) // log the exit error
}
