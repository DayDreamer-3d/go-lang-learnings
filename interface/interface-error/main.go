package main

import (
	"fmt"
	"math"
)

// like exception class in Python
// this can also act like a logger formatter
type NegativeNumber struct{}

// error method is a default method
// like __str__, __init__ in python
func (err *NegativeNumber) Error() string {
	return "Cannot have a number less than 0."
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		return math.Sqrt(x), nil
	} else {
		return 0, &NegativeNumber{}
	}
}

func main() {
	nums := []float64{4, 1, 0, -2}
	for _, num := range nums {
		numSqrt, err := Sqrt(num)
		if err == nil {
			fmt.Printf("Sqrt of %f is %f\n", num, numSqrt)
		} else {
			fmt.Println(err)
		}
	}
}
