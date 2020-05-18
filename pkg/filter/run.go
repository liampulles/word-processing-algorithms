package filter

import (
	"io"
	"log"
	"os"

	"github.com/liampulles/word-processing-algorithms/pkg/common"

	"github.com/liampulles/word-processing-algorithms/pkg"
)

var namedFilters = []pkg.LineHandlerDescriptor{
	{
		Name:        "abecedarian",
		Description: "Selects lines whose letters read in alphabetical order",
		LineHandler: handlerWrapper(IsAbecedarian),
	},
	{
		Name:        "distinct",
		Description: "Selects lines whose alphanumeric characters are distinct",
		LineHandler: handlerWrapper(IsDistinct),
	},
	{
		Name:        "palindrome",
		Description: "Selects lines which read the same back to front",
		LineHandler: handlerWrapper(IsPalindrome),
	},
}

var logger = log.New(os.Stderr, "", 0)

// Run filters in to out via the filter selected by parsing args
func Run(args []string, in io.Reader, out io.Writer) int {
	return common.Run(args, in, out, namedFilters)
}

func handlerWrapper(filter pkg.WordFilter) pkg.LineHandler {
	return func(line string) *string {
		if filter(line) {
			return &line
		}
		return nil
	}
}
