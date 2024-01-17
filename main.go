package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := getUserString()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	tabooMap := getCensorMap(file)
	for {
		switch word := getUserString(); word {
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println(censorFilter(tabooMap, word))
		}
	}
}

func censorFilter(tabooMap map[string]bool, word string) (censor string) {
	if _, ok := tabooMap[strings.ToLower(word)]; ok {
		return getCensoredWord(word)
	}

	return word

}

func getCensorMap(file []byte) (censoredWords map[string]bool) {
	taboos := strings.Split(string(file), "\n")
	censoredWords = map[string]bool{}
	for _, taboo := range taboos {
		censoredWords[strings.ToLower(taboo)] = true
	}

	return censoredWords
}

func getCensoredWord(word string) (censoredWord string) {
	for _ = range word {
		censoredWord += "*"
	}

	return censoredWord
}

func getUserString() (str string) {
	fmt.Scanln(&str)
	return str
}
