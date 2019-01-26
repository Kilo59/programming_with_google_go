package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn ...
// Returns a function with computes displacement as a function of time.
func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + vo*t + so
	}
	return fn
}

func main() {
	fn := GenDisplaceFn(10, 2, 1)
	fmt.Println(fn(3))
	fmt.Println(fn(5))

}
