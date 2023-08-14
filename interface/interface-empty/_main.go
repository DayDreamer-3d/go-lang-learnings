// Empty interfaces used when we do not know the type of the method parameter
// just as object in python
// https://go.dev/tour/methods/14
package main

import (
	"fmt"
	"log"
)

type Human struct {
	kind string
	name string
}

type Fish struct {
	kind string
	name string
}

type MyFloat float64

func instantiate(i interface{}) (newInstance interface{}) {
	// this switch case demo type assertion
	// this switch case check the type of the "i"
	// this statement is very specific for switch case
	switch i.(type) {
	case Human:
		newInstance = Human{"Human", "John"}
	case Fish:
		newInstance = Fish{"Fish", "Smelly"}
	case MyFloat:
		newInstance = 4.1
	case int:
		newInstance = 2
	default:
		log.Fatal("Wrong type.")
	}
	// implicit return statement
	// because we defined returned variable name
	// in the function definition
	return
}

//
// func updateInstance(i interface{}) {
// 	switch i.(type) {
// 	case Human:
// this will not work
// empty interface is a method set and not a field set
// to update a field we require the actual concrete type
// stackoverflow question: https://stackoverflow.com/questions/59976812/empty-interfaces-in-golang
// 		i.name = "Terry"
// 	}
// }

func main() {
	// define an array of instances
	// with different type of instances
	// Note: "instance{}" is a type therefore,
	// we need additional {} to make it an array
	// Also Note: this array holds actual values and not the address
	instances := []interface{}{
		Human{},
		Fish{},
		MyFloat(0),
		0,
	}

	// create a new instance based on the instance type
	for index, instance := range instances {
		newInstance := instantiate(instance)
		fmt.Printf("(%v, %T)\n", newInstance, newInstance)
		instances[index] = newInstance
	}

	// create a new array addresses of instance pointers
	var instanceAddresses []interface{}
	for _, instance := range instances {
		instanceAddresses = append(instanceAddresses, &instance)
	}

	fmt.Println(instanceAddresses)
}
