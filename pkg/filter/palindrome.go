package filter

import (
	"strings"

	"github.com/liampulles/word-processing-algorithms/pkg/common"
)

// IsPalindrome returns true if word reads the same back to front, false otherwise
func IsPalindrome(word string) bool {
	// An empty string is not a palindrome (I've decided :P)
	if word == "" {
		return false
	}

	// Strip non-alphanumeric characters and make the same case.
	upperLetters := strings.ToUpper(common.AlphanumericOnly(word))

	// While front and back pointers are not equal or crossed, match the character they point to.
	for frontPtr, backPtr := 0, len(upperLetters)-1; frontPtr < backPtr; frontPtr, backPtr = frontPtr+1, backPtr-1 {
		if upperLetters[frontPtr] != upperLetters[backPtr] {
			return false
		}
	}
	return true
}
