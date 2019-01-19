package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type name struct {
	fname string
	lname string
}

func stopOnError(err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage)
		os.Exit(1)
	}
}

func readLines(fileName string) (contents []byte) {
	contents, err := ioutil.ReadFile(fileName)
	stopOnError(err, "Error reading file")
	return contents
}

func main() {
	// var nameSlice = make([]name, 10)
	fileContents := readLines("contacts.txt")

	for i, b := range fileContents {
		fmt.Println(i, string(b))
	}

	// fmt.Println(fileContents)
	// fmt.Println(string(fileContents))
}
