package main

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.

import (
	"bufio" //bufferes input output for NewScanner(*os.File)
	"fmt"   //formating package e.g. Printf
	"os"    //for os.Stin which is the standart input
)

func main() {
	// make an empty map which holds a set ofkey/value pairsand provides cons tant-t imeoperat ions tostore
	counts := make(map[string]int)
	// NewScanner returns a new Scanner to read from r.
	// The split function defaults to ScanLines.
	input := bufio.NewScanner(os.Stdin) //*bufio.Scanner()
	//while loop. until there is new line -> true or false
	for input.Scan() {
		//exponds to
		//line := input.Text() //-> reads the next line and removes the newline character from the end
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	// Range-based for lo op
	// Each iteration produces two results, a key and the value of the map element for that key.
	for line, numberOfAppirances := range counts {
		if numberOfAppirances > 1 {
			fmt.Printf("%d\t%s\n", numberOfAppirances, line)
		}
	}
}

// Scan advances the Scanner to the next token, which will then be
// available through the Bytes or Text method. It returns false when the
// scan stops, either by reaching the end of the input or an error.
// After Scan returns false, the Err method will return any error that
// occurred during scanning, except that if it was io.EOF, Err
// will return nil.
// Scan panics if the split function returns 100 empty tokens without
// advancing the input. This is a common error mode for scanners.
