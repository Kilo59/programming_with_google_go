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
	file, err := os.Open(fileName)
	stopOnError(err, "Error reading file")
	defer file.Close()

	// Use the size of the file to determine length of the slice
	fileInfo, _ := file.Stat()
	var strSlice = make([]string, 0, fileInfo.Size())

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
