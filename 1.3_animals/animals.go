package main

import "fmt"

func userRequest() (animal string, attributeRequest string) {
	var a string
	var r string
	fmt.Print(">")
	fmt.Scan(&a, &r)
	return a, r
}

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

func main() {
	for {
		a, r := userRequest()
		fmt.Println(a, r)
	}
}
