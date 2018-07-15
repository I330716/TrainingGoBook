//Reads file names from the command line and print how many times
//a line is repeathed if any.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		//open a file and read it at once
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			//if there is an error print it and continuo to the nex file
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		//Lets count the duplicated lines
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	//now lets print the result
	for line, numberOfAppirances := range counts {
		//print lines that appears more than once
		if numberOfAppirances > 1 {
			fmt.Printf("%d\t%s\n", numberOfAppirances, line)
		}
	}
}
