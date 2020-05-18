package filter

import (
	"strings"

	"github.com/liampulles/word-processing-algorithms/pkg/common"
)

// IsDistinct returns true if each alphanumeric character in a word is unique
func IsDistinct(word string) bool {
	// An empty string is not distinct (I've decided :P)
	if word == "" {
		return false
	}

	// Strip non-alphanumeric characters and make the same case.
	upperLetters := strings.ToUpper(common.AlphanumericOnly(word))

	// Add all characters to a "set" and see if the length matches
	set := make(map[rune]bool)
	for _, char := range upperLetters {
		set[char] = true
	}
	return len(set) == len(upperLetters)
}
