package main

import "fmt"

// func getUserInput(promptMessage string) (userInt float64) {
// 	var userInputValue string
// 	fmt.Print(promptMessage)
// 	fmt.Scanln(&userInputValue)
// 	userInt, err := strconv.ParseFloat(userInputValue, 64)
// 	if err != nil {
// 		fmt.Println("Could not parse Integer")
// 		os.Exit(1)
// 	}
// 	return userInt
// }

// func getUserInputLine(promptMessage string) (userInputStr string) {
// 	// using bufio to avoid "problems" with fmt.Scan and spaces
// 	scanner := bufio.NewScanner(os.Stdin)
// 	fmt.Println(promptMessage)
// 	scanner.Scan()
// 	return scanner.Text()
// }

// Animal ...
type Animal struct {
	food       string
	locomotion string
	noise      string
}

// Eat ...
func (a Animal) Eat() {
	fmt.Println(a.food)
}

// Move ...
func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

// Speak ...
func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {}
