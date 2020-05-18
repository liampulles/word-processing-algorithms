package common

import (
	"regexp"
)

var alphanumericRegex, _ = regexp.Compile("[^a-zA-Z0-9]+")

// AlphanumericOnly strips non-alphanumeric characters from a string
func AlphanumericOnly(in string) string {
	return alphanumericRegex.ReplaceAllString(in, "")
}
