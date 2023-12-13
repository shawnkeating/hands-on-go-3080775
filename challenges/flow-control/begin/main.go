// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	// use os package to read the file specified as a command line argument

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(fmt.Errorf("Failed to read file: %s", err))
	}

	// convert the bytes to a string
	str := string(file)

	// initialize a map to store the counts
	m := map[string]int{"words": 0, "letters": 0, "symbols": 0, "numbers": 0}

	// use the standard library utility package that can help us split the string into words
	words := strings.Fields(str)

	// capture the length of the words slice
	numWords := len(words)
	m["words"] = numWords

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	letters := 0
	numbers := 0
	symbols := 0
	for _, v := range words {
		for _, k := range v {
			switch {
			case unicode.IsNumber(k) == true:
				numbers++
			case unicode.IsPunct(k) == true:
				symbols++
			default:
				letters++
			}

		}
	}

	m["letters"] = letters
	m["numbers"] = numbers
	m["symbols"] = symbols

	// dump the map to the console using the spew package
	spew.Dump(m)
}
