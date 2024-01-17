package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := getUserString()
	word := getUserString()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(isWordTaboo(file, word))
	return
}

func isWordTaboo(file []byte, word string) (isTaboo bool) {
	taboos := strings.Split(string(file), "\n")
	tabooWords := map[string]bool{}
	for _, taboo := range taboos {
		tabooWords[strings.ToLower(taboo)] = true
	}

	if _, ok := tabooWords[strings.ToLower(word)]; ok {
		return true
	}

	return false

}

func getUserString() (str string) {
	fmt.Scanln(&str)
	return str
}
