package main

import (
	"fmt"
	"sort"
)

// Swap ...
// Swaps positions of 2 adjacent elements
// Assignment Requirement
func Swap(index int, intSlice []int) {
	fmt.Println("Swap position")
}

func sweep(intSlice *[]int) {
	fmt.Println("Sweep")
}

// BubbleSort ...
// Assignment Requirement
func BubbleSort(intSlice []int) {
	sorted := false
	n := len(intSlice)
	for sorted == false {
		for i := 0; i < n-1; i++ {
			if intSlice[i+1] < intSlice[i] {
				intSlice[i+1], intSlice[i] = intSlice[i], intSlice[i+1]
			}
		}
		fmt.Println("Sweep")
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
	intSlice := []int{6, 2, 5, 4, 1}
	// fmt.Println(intSlice)
	BubbleSort(intSlice)
}
