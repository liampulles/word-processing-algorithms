package filter

import (
	"strings"

	"github.com/liampulles/word-processing-algorithms/pkg/common"
)

// IsAbecedarian returns true if the letters are in alphabetical order
func IsAbecedarian(word string) bool {
	// An empty string is not abacedarian (I've decided :P)
	if word == "" {
		return false
	}

	// Strip non-alphanumeric characters and make the same case.
	upperLetters := strings.ToUpper(common.AlphanumericOnly(word))

	// While successive letters are "larger" than the previous, continue
	for i := 0; i < len(upperLetters)-1; i++ {
		curr, next := upperLetters[i], upperLetters[i+1]
		if next < curr {
			return false
		}
	}
	return true
}
