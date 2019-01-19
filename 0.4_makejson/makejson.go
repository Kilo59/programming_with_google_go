package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func getUserInputLine(promptMessage string) (userInputStr string) {
	// using bufio to avoid "problems" with fmt.Scan and spaces
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(promptMessage)
	scanner.Scan()
	return scanner.Text()
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
		"name":    getUserInputLine("Enter a name:"),
		"address": getUserInputLine("Enter an address:")}

	addressMapJSON := mapToJSON(addressMap)
	fmt.Println(string(addressMapJSON))
}
