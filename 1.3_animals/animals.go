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
	animals := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	for {
		a, r := userRequest()
		fmt.Println(a, r)
		switch r {
		case "eat":
			animals[a].Eat()
		case "move":
			animals[a].Move()
		case "speak":
			animals[a].Speak()
		}
	}
}
