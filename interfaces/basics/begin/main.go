// interfaces/basics/begin/main.go
package main

import "fmt"

// define a humanoid interface with speak and walk methods returning string
type humanoid interface {
	speak() string
	walk() string
}

// define a person type that implements humanoid interface
type person struct {
	name string
}

func (p person) speak () string {return "Person speaks"}
func (p person) walk () string {return "Person walks"}

// implement the Stringer interface for the person type
func (p person) String () string {return fmt.Sprintf("My name, is %s",p.name)}

// define a dog type that can walk but not speak
// type dog struct {}

// func (d dog) walk () string {return "Dog walks"}

func main() {
	// invoke with a person
	person1 := person{name:"James"}
	doHumanThings(person1)

	// can we invoke with a dog?
	// doHumanThings(dog{})

	fmt.Println(person1)
}

func doHumanThings(h humanoid) {
	fmt.Println(h.speak())
	fmt.Println(h.walk())
}
