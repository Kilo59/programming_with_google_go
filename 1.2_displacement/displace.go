package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func getUserInput(promptMessage string) (userInt float64) {
	var userInputValue string
	fmt.Print(promptMessage)
	fmt.Scanln(&userInputValue)
	userInt, err := strconv.ParseFloat(userInputValue, 64)
	if err != nil {
		fmt.Println("Could not parse Integer")
		os.Exit(1)
	}
	return userInt
}

// GenDisplaceFn ...
// Returns a function with computes displacement as a function of time.
func GenDisplaceFn(a float64, vo float64, so float64) func(float64) float64 {
	fn := func(t float64) float64 {
		return (a*math.Pow(t, 2))/2 + vo*t + so
	}
	return fn
}

func main() {
	a := getUserInput("Acceleration:")
	vo := getUserInput("Initial Velocity:")
	so := getUserInput("Initial Displacement:")
	fn := GenDisplaceFn(a, vo, so)
	time := getUserInput("Time:")
	fmt.Println(fn(time))
}
