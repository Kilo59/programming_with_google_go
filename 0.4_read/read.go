package main

import (
	"bufio"
	"fmt"
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

func readLines(fileName string) (linesSlice []string) {
	// TODO: check the size of the file before making slice
	var strSlice = make([]string, 0, 100)

	file, err := os.Open(fileName)
	stopOnError(err, "Error reading file")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		strSlice = append(strSlice, scanner.Text())
	}
	return strSlice
}

func main() {
	// var nameSlice = make([]name, 10)
	lines := readLines("contacts.txt")

	for i, v := range lines {
		fmt.Println(i, v)
	}

}
