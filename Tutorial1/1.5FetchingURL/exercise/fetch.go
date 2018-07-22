// Feth prints  the content found at a URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		//Exercise 1.8 Modyfy fetch to add prefix http:// if it is missing
		if !strings.HasPrefix(url, "http://") || strings.HasPrefix(url, "https://") {
			url = "http://" + url
		}
		// makes http request and if there is no err it returns
		// the result in response struct
		respoonse, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//don't know if closing without using the body is requared
			respoonse.Body.Close()
			continue
		}
		//Exercise 1.9 Modify fetch to show the response status
		fmt.Fprintln(os.Stdout, "Status: "+respoonse.Status)
		// Exercise 1.7 avoid buffer copy of the response by showing it on stdout directly
		// type Writer(os.Stdout) is interface that wraps the basic Write method
		// type Reader(resp.Body) is the interface that wraps the basic Read methond
		_, err = io.Copy(os.Stdout, respoonse.Body)
		respoonse.Body.Close() // close to not leack resources
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			continue
		}
	}
}
