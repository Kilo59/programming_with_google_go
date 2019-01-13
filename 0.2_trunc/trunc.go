package main

import (
	"fmt"
	"os"
	"strconv"
)

func truncate(floatNumber float32) (truncatedNumber int) {
	return int(floatNumber)
}

func main() {

	var usrFloat string

	//if no arguments are passed prompt the user to enter a float
	if len(os.Args) > 1 {
		usrFloat = os.Args[1]
	} else {
		fmt.Print("Enter a float: ")
		fmt.Scanln(&usrFloat)
	}

	//convert argument to a float type.
	argAsFloat, err := strconv.ParseFloat(usrFloat, 32)
	if err != nil {
		fmt.Println("Could not parse argument as Float")
		os.Exit(1)
	}

	fmt.Println("Truncating", usrFloat)
	// strconv.ParseFloat always returns a float64, convert to float32.
	fmt.Println(truncate(float32(argAsFloat)))
}
