package main

import (
	"fmt"
	"os"
)

func checkUserArgs() (userArgs []string) {
	if len(os.Args) < 2 {
		fmt.Println("Please rerun slice.go providing integers as arguments")
		fmt.Println("Example: go run slice.go 1 22 58")
		os.Exit(1)
	}
	return os.Args[1:]
}

func main() {
	userArgs := checkUserArgs()
	fmt.Println(userArgs)
	fmt.Println(userArgs[0])
	fmt.Println(userArgs[1])
}
