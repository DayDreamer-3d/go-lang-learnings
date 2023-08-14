// interface in go
// https://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go

package main

import (
	"fmt"
	"math"
)

// Geometry interface says any type which has area method
// and returns float can be of Geometry type
type Geometry interface {
	area() float64
	geoTyp() string
}

type Rect struct {
	length  int64
	breadth int64
	kind    string
}

func (r *Rect) area() float64 {
	return float64(r.breadth) * float64(r.length)
}

func (r *Rect) geoTyp() string {
	return r.kind
}

type Circle struct {
	radius int64
	kind   string
}

func (c *Circle) area() float64 {
	return math.Pi * math.Pow(float64(c.radius), 2)
}

func (c *Circle) geoTyp() string {
	return c.kind
}

func GeometryKind(g Geometry) string {
	return g.geoTyp()
}

func main() {

	// create a geometries array
	// with first element being Rect Address
	// second element being Circle Address
	// remember: we have defined area method on pointer
	// but we have option of having array elements as pointers or actual values
	// because even pointer methods can be invoked by values
	geometries := []Geometry{&Rect{3, 2, "Rectangle"}, &Circle{3, "Circle"}}
	for _, geometry := range geometries {
		// an external function can use Geometry interface
		// and process the inputs
		fmt.Println("Geometry Type: ", GeometryKind(geometry))
		// here we do not care what geometry we have
		// we just care it has area method [duck typing]
		fmt.Println("Area: ", geometry.area())
	}
}
