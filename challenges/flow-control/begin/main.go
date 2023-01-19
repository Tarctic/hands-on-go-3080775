// challenges/flow-control/begin/main.go
package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	// handle any panics that might occur anywhere in this function
	defer func() {
		if err := recover(); err != nil {
			fmt.Print("Recovered from error:", err)
		}
	}()

	// use os package to read the file specified as a command line argument
	file := os.Args
	bytes, err := os.ReadFile(file[1])
	if err != nil {
		panic(fmt.Errorf("could not read file error: %s", err))
	}

	// convert the bytes to a string
	str := string(bytes)

	// initialize a map to store the counts
	m := map[string]int{"letters": 0, "numbers": 0, "symbols": 0, "words": 0}

	// use the standard library utility package that can help us split the string into words
	words := strings.Split(str, " ")

	// capture the length of the words slice
	m["words"] = len(words)

	// use for-range to iterate over the words slice and for each word, count the number of letters, numbers, and symbols, storing them in the map
	for _, word := range words {
		for _, i := range word {
			ascii := int(i)
			if (ascii >= 65 && ascii <= 90) || (ascii >= 97 && ascii <= 122) {
				m["letters"]++
			} else if ascii >= 48 && ascii <= 57 {
				m["numbers"]++
			} else {
				m["symbols"]++
			}
		}
	}

	// dump the map to the console using the spew package
	spew.Dump(m)
}
