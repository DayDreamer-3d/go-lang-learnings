package main

import (
	"fmt"
	"log"
)

type Animal struct {
	name string
	kind string
	age  uint8
}

func (a Animal) Speak() {
	switch a.kind {
	case "dog":
		{
			fmt.Println("woof !!")
		}
	case "cat":
		{
			fmt.Println("meow !")
		}
	default:
		{
			log.Fatal("Not supported Animal.")
		}
	}
}

func main() {

	dog := Animal{
		name: "cuddle",
		kind: "dog",
		age:  3,
	}
	dog.Speak()

	human := Animal{
		name: "John",
		kind: "human",
		age:  4,
	}
	human.Speak()
}
