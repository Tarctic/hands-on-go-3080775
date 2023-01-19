// flow-control/defer-panic-recover/begin/main.go
package main

import "fmt"

func cleanup(msg string) {
	fmt.Println(msg)
}

func main() {
	// defer function calls
	defer cleanup("first") // First in, last out
	defer cleanup("second") // defer means to keep for later

	fmt.Println("Main print") // normal print statement, not deferred

	// defer recovery
	// since it is kept for last, it sees the error and recovers from it
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic:",err) 
		}
	}()

	// panic
	panic("Oh no! I'm panicking! O_O") // causes an error
}
