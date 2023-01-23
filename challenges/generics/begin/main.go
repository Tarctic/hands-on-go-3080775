// challenges/generics/begin/main.go
package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Part 1: print function refactoring

// non-generic print functions

// func printString(s string) { fmt.Println(s) }

// func printInt(i int) { fmt.Println(i) }

// func printBool(b bool) { fmt.Println(b) }

func printAny[T any](val T) {
	fmt.Println(val)
}

// Part 2 sum function refactoring

type numeric interface {constraints.Float | constraints.Integer}

// sum sums a slice of any type
func sumAny[T numeric](numbers ...T) (T) {
	var sum T
	for _, n := range numbers {
		sum += n
	}
	return sum
}

func main() {
	// call non-generic print functions
	// printString("Hello")
	// printInt(42)
	// printBool(true)

	// call generic printAny function for each value above
  printAny("Hello")
  printAny(42)
  printAny(true)

	// call generics sum function
	fmt.Println(sumAny(1,2,3))

}
