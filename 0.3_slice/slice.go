package main

import (
	"fmt"
	"os"
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

}
