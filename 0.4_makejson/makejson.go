package main

import (
	"fmt"
	"os"
)

func getUserInput(promptMessage string) (userInputStr string) {
	var userInputValue string

	fmt.Println(promptMessage)
	itemsScanned, err := fmt.Scanln(&userInputValue)
	if err != nil {
		fmt.Println("\tCould not parse input")
		fmt.Printf("\tItems scanned %d\n", itemsScanned)
		os.Exit(1)
	}
	return userInputValue
}

func main() {
	addressMap := map[string]string{
		"name":    getUserInput("Enter a name:"),
		"address": getUserInput("Enter an address:")}
	fmt.Println(addressMap)
}
