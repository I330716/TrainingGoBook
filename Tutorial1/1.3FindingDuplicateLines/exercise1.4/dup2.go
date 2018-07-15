// Dup2 prints the count and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

//LineMeta saves meta data for given line as number of accurences and in which file appears
type LineMeta struct {
	numberOfAccurrences int
	//this will play role of set
	fileOfAccurrences map[string]bool
}

func main() {
	counts := make(map[string]LineMeta)
	files := os.Args[1:]
	//if there is no command line arguments
	if len(files) == 0 {
		countLine(os.Stdin, counts)
	} else {
		for _, filename := range files {
			file, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v'n", err)
			}
			countLine(file, counts)
			file.Close()
		}
	}
	for line, lineMeta := range counts {
		if lineMeta.numberOfAccurrences > 1 {
			fmt.Printf("%d\t%s\n", lineMeta.numberOfAccurrences, line)
			fmt.Printf("Found in:")
			for filename := range lineMeta.fileOfAccurrences {
				fmt.Printf(" %s", filename)
			}
			fmt.Println()
		}
	}
}

func countLine(file *os.File, counts map[string]LineMeta) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		//Here we get copy of the LineMeta not a reference
		lineMeta, ok := counts[text]
		if !ok {
			//if there is no meta data for this line it means
			//that the line have not yet accured and we must make
			// new map for it unless we want runtime exeption "assignment to entry in nil map"
			lineMeta.fileOfAccurrences = make(map[string]bool)
		}
		lineMeta.numberOfAccurrences++
		//add the name of the file in our Set of strings
		lineMeta.fileOfAccurrences[file.Name()] = true
		counts[text] = lineMeta
	}

}
