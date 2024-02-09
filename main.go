package main

import (
	"fmt"
	"os"
	"strings"
	"bufio"
)

func main() {
	fmt.Println("Enter a filepath to your list of taboo words (formatted as one word per line): ")
	filename := getUserString()

	file, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	tabooMap := getCensorMap(file)
	for {
		fmt.Println("\nEnter something to censor: ")
		switch sentence := getUserString(); sentence {
		case "exit":
			fmt.Println("Bye!")
			return
		default:
			fmt.Println("\nCensored: " + censorFilter(tabooMap, sentence))
		}
	}
}

func censorFilter(tabooMap map[string]bool, sentence string) (censor string) {
	var periodFlag bool
	if strings.Contains(sentence, ".") {
		sentence = strings.Replace(sentence, ".", "", 1)
		periodFlag = true
	} else {
		periodFlag = false
	}
	words := strings.Split(sentence, " ")

	var censorArray []string
	for _, word := range words {
		if _, ok := tabooMap[strings.ToLower(word)]; ok {
			censorArray = append(censorArray, getCensoredWord(word))
		} else {
			censorArray = append(censorArray, word)
		}

	}

	censor = strings.Join(censorArray, " ")
	if periodFlag {
		censor += "."
	}

	return censor

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
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	str = scanner.Text()
	return str
}
