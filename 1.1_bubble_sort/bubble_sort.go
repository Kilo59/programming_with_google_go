package main

import (
	"fmt"
)

// Swap ...
// Swaps positions of 2 adjacent elements
func Swap(index int, intSlice []int) {
	fmt.Println("Swap position")
}

// BubbleSort ...
func BubbleSort(intSlice []int) {
	fmt.Println(intSlice)
}

func stopOnError(err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage)
		panic(err)
	}
}

func userInputSlice(size int) (intSlice []int) {
	fmt.Printf("Enter up to %d integers, one per line\nEnter non integer when finished:\n", size)
	inputInts := make([]int, size)
	for i := range inputInts {
		_, err := fmt.Scan(&inputInts[i])
		if err != nil {
			break
		}
	}
	return inputInts
}

func main() {
	intSlice := userInputSlice(10)
	fmt.Println(intSlice)
}
