// types/maps/mutating/begin/main.go
package main

import "fmt"

type author struct {
	name string
}

func main() {
	var authors = map[string]author{
		"ma": {name: "Marcus Aurelius"},
		"tm": {name: "Toni Morrison"},
	}
	fmt.Printf("%v\n", authors)

	// change the value of an author in the map
	authors["tm"] = author{name: "Toni MyMom"}

	// print the map
	fmt.Printf("%v\n", authors)

	// delete a key from the map
	delete(authors, "tm")

	// print the map
	fmt.Printf("%v\n", authors)
}
