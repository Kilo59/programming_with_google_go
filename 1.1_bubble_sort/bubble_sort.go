package main

import (
	"fmt"
	"sort"
)

// Swap ...
// Swaps positions of 2 adjacent elements
func Swap(index int, intSlice []int) {
	fmt.Println("Swap position")
}

func sweep() {
	fmt.Println("Sweep")
}

// BubbleSort ...
func BubbleSort(intSlice []int) {
	sorted := false
	fmt.Println(intSlice)
	for sorted == false {
		for i := range intSlice {
			fmt.Println(i, intSlice[i])

		}
		sorted = sort.IntsAreSorted(intSlice)
	}
	fmt.Println("Sorted!\n", intSlice)
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
	// intSlice := userInputSlice(10)
	intSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(intSlice)
	BubbleSort(intSlice)
}
