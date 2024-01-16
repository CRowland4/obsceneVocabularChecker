package main

import (
	"fmt"
	"os"
)

func main() {
	filename := getFilename()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	printContents(file)
	return
}

func getFilename() (filename string) {
	fmt.Scanln(&filename)
	return filename
}

func printContents(file []byte) {
	fmt.Println(string(file))
	return
}
