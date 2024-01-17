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
		word := getUserString()
		if word == "exit" {
			fmt.Println("Bye!")
			return
		}
		fmt.Println(wordCensor(tabooMap, word))
	}
}

func wordCensor(tabooMap map[string]string, word string) (censor string) {
	if censoredWord, ok := tabooMap[strings.ToLower(word)]; ok {
		return censoredWord
	}

	return word

}

func getCensorMap(file []byte) (censoredWords map[string]string) {
	taboos := strings.Split(string(file), "\n")
	censoredWords = map[string]string{}
	for _, taboo := range taboos {
		censoredWords[strings.ToLower(taboo)] = getCensoredWord(taboo)
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
