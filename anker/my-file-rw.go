package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "I am learning Go programming language."
	file, err := os.Create("./file.txt")
	checkError(err)
	length, err := io.WriteString(file, content)
	checkError(err)
	fmt.Printf("Wrote %d bytes to file\n", length)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	defer readFile("./file.txt")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(fileName string) {
	file, err := os.Open(fileName)
	checkError(err)
	data := make([]byte, 100)
	count, err := file.Read(data)
	checkError(err)
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}
