// Use pointer for method receivers

package main

import (
	"fmt"
	"log"
)

type Animal struct {
	name    string
	kind    string
	counter uint8
}

// pointer as a receiver
// basically, this method will take animal pointer
func (a *Animal) SpeakUsingPointer() {
	switch a.kind {
	case "dog":
		{
			fmt.Println("woof !!")
			a.counter++
		}
	default:
		{
			log.Fatal("Unsupported Animal.")
		}
	}
}

// this method is using pass by value parameter
func (a Animal) SpeakUsingValue() {
	switch a.kind {
	case "dog":
		{
			fmt.Println("woof !!")
			// this warning indicates counter will not have an affect
			// because "a" is pass by value
			// therefore, we need a return
			a.counter++
		}
	default:
		{
			log.Fatal("Unsupported.")
		}
	}
}

func main() {

	dog := Animal{
		name: "cuddle",
		kind: "dog",
	}
	// this method call is using a pointer and not the value
	// pointer methods use value or pointer object (dog or &dog)
	dog.SpeakUsingPointer()
	// this counter will increment and it will become 1
	// because we are using pointer
	fmt.Println(dog.counter)
	fmt.Println("counter got incremented to 1 because of pointer receiver (method).")
	fmt.Println("")

	dog.SpeakUsingValue()
	// dog.counter will remain 1 because we are passing value
	fmt.Println(dog.counter)
	fmt.Println("no counter incremented because of value receiver (method).")
	fmt.Println("")

	// this method call is using a pointer and not the value
	// pointer methods use value or pointer object (dog or &dog)
	p := &dog
	p.SpeakUsingPointer()
	// this counter will increment by and equal 2
	fmt.Println(dog.counter)
	fmt.Println("counter got incremented to 2 because of pointer receiver (method).")
	fmt.Println("")

	// pointer can call value method receiver
	// why ??
	p.SpeakUsingValue()

}
