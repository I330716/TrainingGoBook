// Fetchall fetchs URLs in parallel and reports their times and size.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a go routine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	//take current time
	start := time.Now()
	//make sequest and take the response
	response, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	//copy the body data in to /dev/null
	nbytes, err := io.Copy(ioutil.Discard, response.Body)
	response.Body.Close() //don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	//time after start in seconds
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
