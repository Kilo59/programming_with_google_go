package main

import (
	"fmt"
	"strings"
)

func findChars(targetChars string, strHaystack string) (found bool) {
	runeTargetChars := []rune(targetChars)
	runeHaystack := []rune(strHaystack)

	//check first character
	if runeTargetChars[0] != runeHaystack[0] {
		return false
	}
	//check last character
	if runeTargetChars[len(runeTargetChars)-1] != runeHaystack[len(runeHaystack)-1] {
		return false
	}
	//checkin for presence of middle character
	if strings.ContainsRune(strHaystack, runeTargetChars[1]) {
		return true
	}
	return false
}

func main() {

	var userInput string
	fmt.Println("Enter Text:")
	fmt.Scan(&userInput)

	if findChars("ian", strings.ToLower(userInput)) == true {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}

}
