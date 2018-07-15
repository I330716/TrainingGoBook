// Feth prints  the content found at a URL
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	for _, url := range os.Args[1:] {
		// makes http request and if there is no err it returns
		// the result in response struct
		respoonse, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			//don't know if closing without using the body is requared
			respoonse.Body.Close()
			continue
		}
		//read hole body and return it as sequence of bytes
		body, err := ioutil.ReadAll(respoonse.Body)
		respoonse.Body.Close() // close to not leack resources
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v", url, err)
			continue
		}
		fmt.Printf("%s", body)
	}
}