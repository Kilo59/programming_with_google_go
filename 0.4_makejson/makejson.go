package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func getUserInput(promptMessage string) (userInputStr string) {
	var userInputValue string

	fmt.Println(promptMessage)
	// Scan breaks on spaces
	itemsScanned, err := fmt.Scanln(&userInputValue)
	if err != nil {
		fmt.Println("\tCould not parse input!")
		fmt.Printf("\tItems scanned %d\n", itemsScanned)
		os.Exit(1)
	}
	return userInputValue
}

func mapToJSON(mapInput map[string]string) (jsonB []byte) {
	jsonB, err := json.Marshal(mapInput)
	if err != nil {
		fmt.Println("\tCould not marshal mapInput object!")
		os.Exit(1)
	}
	return jsonB
}

func main() {
	addressMap := map[string]string{
		"name":    getUserInput("Enter a name:"),
		"address": getUserInput("Enter an address:")}
	// remove debug statement
	fmt.Println(addressMap)

	addressMapJSON := mapToJSON(addressMap)
	fmt.Println(string(addressMapJSON))
}
