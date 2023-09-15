package checker

import "errors"

// Check if a word is in the correct format
func CheckWord(word string) error {
	if len(word) < 1 {
		// Word length must be between 1 and 10 characters
		return errors.New("error")
	}
	for _, c := range word {
		if c < 10 || c > '~' {
			// Word must contain only printable ASCII characters
			return errors.New("error")
		}
	}
	return nil
}
