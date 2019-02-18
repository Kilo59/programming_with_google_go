package main

import (
	"fmt"
	"sort"
)

// Swap ...
// Swaps positions of 2 adjacent elements
// Assignment Requirement
func Swap(index int, intSlice []int) {
	intSlice[index+1], intSlice[index] = intSlice[index], intSlice[index+1]
}

// BubbleSort ...
// Assignment Requirement
func BubbleSort(intSlice []int) {
	sorted := false
	n := len(intSlice)
	sweeps := 1
	for sorted == false {
		for i := 0; i < n-1; i++ {
			if intSlice[i+1] < intSlice[i] {
				Swap(i, intSlice)
			}
		}
		sweeps++
		sorted = sort.IntsAreSorted(intSlice)
	}
	fmt.Printf("Sorted! %d sweeps\n", sweeps)
	fmt.Println(intSlice)
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
	BubbleSort(intSlice)
}
