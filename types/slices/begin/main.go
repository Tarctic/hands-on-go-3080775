// types/slices/begin/main.go
package main

import "fmt"

func main() {
	// declare a slice string using the make function
	var names = make([]string, 0)

	// append 3 names to the slice
	names = append(names, "Sami")
	names = append(names, "Rumi")
	names = append(names, "Hemi")

	// print the slice
	fmt.Println(names)

	// slice the slice using syntax slice[low:high]
	fmt.Println(names[1:3]) // [Jane Mary]
	fmt.Println(names[1:]) // [Jane Mary]
	fmt.Println(names[:1]) // [John]
	fmt.Println(names[:]) // [John Jane Mary]
}
