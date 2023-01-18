// packages/basics/main.go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func greet() string {
	return "HELLO"
}

func greetWithName(name string) string {
	return "HELLO " + name
}

func greetWithNameAndAge(name string, age int) (greeting string) {
	greeting = "Hello " + name + " of age " + strconv.Itoa(age)
	return
}

func divide(a, b int) (int, error) {
	if b==0 {
		return 0, errors.New("cannot divide by Z E R 0")
	}
	return a/b, nil
}

func main() {
	fmt.Println(greet())
	fmt.Println(greetWithName("Zainab"))
	fmt.Println(greetWithNameAndAge("Zainab", 19))
	fmt.Printf("Today is a %s \n", time.Now().Weekday())
	fmt.Println(divide(5,0))
	fmt.Println(divide(51,20))
}
