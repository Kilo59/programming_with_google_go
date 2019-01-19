package main

import (
	"fmt"
	"os"
)

func stopOnError(err error, errorMessage string) {
	if err != nil {
		fmt.Println(errorMessage)
		os.Exit(1)
	}
}

func main() {
	fmt.Println("read")
	f, err := os.Create("outfile.txt")
	stopOnError(err, "Could not create file")

	barr := []byte{1, 2, 3}
	nb1, err := f.Write(barr)
	stopOnError(err, "Could not write byte array")
	nb2, err := f.WriteString("Hello World")
	stopOnError(err, "Could not write string")
	fmt.Println(nb1, nb2)
}
