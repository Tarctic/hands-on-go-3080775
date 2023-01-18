// types/conversion/begin/main.go
package main

import "fmt"

func main() {
	// declare float and convert it to an unsigned int
	a := 43.4
	b := uint(a)
	fmt.Printf("a=%v (%T), b=%v (%T)",a,a,b,b)
}
