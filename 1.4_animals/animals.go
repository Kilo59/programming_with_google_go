package main

import "fmt"

func userRequest() (query string, animalName string, infoRequest string) {
	var q string
	var an string
	var ir string
	fmt.Print(">")
	fmt.Scan(&q, &an, &ir)
	return q, an, ir
}

// Animal interface type
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow ...
type Cow struct {
	food       string
	locomotion string
	noise      string
}

// Bird ...
type Bird struct {
	food       string
	locomotion string
	noise      string
}

// Snake ...
type Snake struct {
	food       string
	locomotion string
	noise      string
}

// Eat Cow...
func (c Cow) Eat() {
	fmt.Println(c.food)
}

// Eat Bird...
func (b Bird) Eat() {
	fmt.Println(b.food)
}

// Eat Snake...
func (s Snake) Eat() {
	fmt.Println(s.food)
}

// Move Cow...
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}

// Move Bird...
func (b Bird) Move() {
	fmt.Println(b.locomotion)
}

// Move Snake...
func (s Snake) Move() {
	fmt.Println(s.locomotion)
}

// Speak Cow...
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

// Speak Bird...
func (b Bird) Speak() {
	fmt.Println(b.noise)
}

// Speak Snake...
func (s Snake) Speak() {
	fmt.Println(s.noise)
}

func main() {
	fmt.Println("newanimal <name> <animal type>")
	fmt.Println("query <name> <info>")
	animals := map[string]Animal{
		"cow":   Cow{"grass", "walk", "moo"},
		"bird":  Bird{"worms", "fly", "peep"},
		"snake": Snake{"mice", "slither", "hsss"},
	}
	namedAnimals := map[string]Animal{}
	for {
		q, an, ir := userRequest()
		switch q {
		case "newanimal":
			namedAnimals[an] = animals[ir]
			fmt.Println("Created it!")
		case "query":
			switch ir {
			case "eat":
				namedAnimals[an].Eat()
			case "move":
				namedAnimals[an].Move()
			case "speak":
				namedAnimals[an].Speak()
			}
		}

	}
}
