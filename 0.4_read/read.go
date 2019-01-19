package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type name struct {
	fname string
	lname string
}

func getUserInput(promptMessage string) (userInputStr string) {
	var userInputValue string
	fmt.Println(promptMessage)
	fmt.Scanln(&userInputValue)
	return userInputValue
}

func stopOnError(err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage)
		panic(err)
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
	lines := readLines(getUserInput("Filename?:"))
	var nameSlice = make([]name, 0, len(lines))

	// could have done this during readLines() function
	for _, v := range lines {
		var n name
		names := strings.Fields(v)
		n.fname = names[0]
		n.lname = names[1]
		nameSlice = append(nameSlice, n)
	}

	fmt.Println("# First\tLast\n------------")
	for i, v := range nameSlice {
		fmt.Printf("%d %s %s\n", i, v.fname, v.lname)
	}
}
