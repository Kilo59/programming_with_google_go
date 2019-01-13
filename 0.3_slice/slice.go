package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
)

func getUserInput(promptMessage string) (userInt int) {
	var userInputValue string

	fmt.Println(promptMessage)
	fmt.Scanln(&userInputValue)

	userInt, err := strconv.Atoi(userInputValue)
	if err != nil {
		if userInputValue == "X" {
			fmt.Println("Exiting loop")
			os.Exit(0)
		}
		fmt.Println("Could not parse Integer")
		os.Exit(1)
	}
	return userInt
}

func main() {
	// capacity of 8 to account for number of tests required in review criteria
	// Underlying array will not need to expand unless more than 5 integers are entered
	var integerSlice = make([]int, 3, 8)

	for {
		userInput := getUserInput("Enter an Integer (X to quit):")
		integerSlice = append(integerSlice, userInput)
		sort.Ints(integerSlice)
		fmt.Println(integerSlice)
	}

}
