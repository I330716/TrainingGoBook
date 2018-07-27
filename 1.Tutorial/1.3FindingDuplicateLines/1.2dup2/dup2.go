// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) < 1 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			file, err := os.Open(arg)
			if err != nil {
				//%v means display value in a natural format
				fmt.Fprintf(os.Stderr, "%v'n", err)

				//go exactly on the next interation
				continue
			}
			countLines(file, counts)
			//we must free the resources of the file by closing it
			file.Close()
		}
	}

	for line, numberOFAppirances := range counts {
		if numberOFAppirances > 1 {
			fmt.Printf("%d\t%s\n", numberOFAppirances, line)
		}
	}
}

//Here counts is reference of map and what we do in the function body with
//this map will efect the map which caller have passed as an argument
func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)
	//while loop until there is new line
	for input.Scan() {
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
}
