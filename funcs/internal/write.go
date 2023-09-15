package internal

import (
	"strings"
)

func Write(word string, hash map[rune][][]string) string {
	// Read input from command line argument
	word = strings.ReplaceAll(word, "\r", "")
	word = strings.ReplaceAll(word, "\n", "\\n")
	inputWords := strings.Split(word, "\\n")
	res := ""
	for _, word := range inputWords {
		if word != "" {
			// If the input word is not empty, print all possible permutations
			wordRunes := []rune(word)
			// lastChar := rune(word[len(word)-1])
			for i := 0; i < 8; i++ {
				// For each character in the word, print out all possible permutations from the hash map
				for index, v := range wordRunes {
					wordList := hash[v][i]
					for _, el := range wordList {
						// fmt.Print(el)
						res += el
					}
					if index == len(wordRunes)-1 {
						// If the current character is the last character in the word, print a newline character
						// fmt.Print("\n")
						res += "\n"
					}
				}
			}
		} else {
			// If the input word is empty, print a newline character
			// fmt.Print("\n")
			res += "\n"
		}
	}
	return res
}
