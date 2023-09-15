package internal

import (
	"errors"
	"os"
	"strings"
)

func CreateHash(format string) (map[rune][][]string, error) {
	// Read the file based on the specified format
	file, err := os.ReadFile("funcs/internal/data/" + format + ".txt")
	if err != nil {
		return map[rune][][]string{}, errors.New("error")
	}
	// Replace all "\r" characters with empty strings and split the contents into individual words
	content := strings.ReplaceAll(string(file), "\r", "")
	words := strings.Split(content, "\n")
	// Create a list of lists, where each inner list contains all the permutations of a word
	wordList := make([][]string, len(words))
	for i, word := range words {
		wordList[i] = strings.Split(word, "\n")
	}
	// Create a hash map where the key is a rune (a character) and the value is a list of permutations
	hash := make(map[rune][][]string)
	for v := ' '; v <= '~'; v++ {
		i := int(v-31) * 9
		for j := i - 8; j < i; j++ {
			hash[v] = append(hash[v], wordList[j])
		}
	}
	return hash, nil
}
