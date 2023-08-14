package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {

	// instantiate Vertex struct
	v := Vertex{1, 2}
	// pointer p points to v struct
	p := &v
	// change struct attribute value using the pointer
	p.X = 1e9
	// print struct using the pointer
	fmt.Println(*p)

	// q points to struct.X
	q := &p.X
	//change struct.X value using q pointer
	*q = 1
	// print struct.X using q pointer
	fmt.Println(*q)
	// print struct using p pointer
	fmt.Println(*p)
}
