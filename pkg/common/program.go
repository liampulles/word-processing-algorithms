package common

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/liampulles/word-processing-algorithms/pkg"
)

var logger = log.New(os.Stderr, "", 0)

// Run selects a handler using args, then processes it for each line of in, and
// writes the output of the selected handler to out
func Run(args []string, in io.Reader, out io.Writer, namedHandlers []pkg.LineHandlerDescriptor) int {
	// What algorithm to use?
	handler, err := selectAlgorithm(args, namedHandlers)
	if err != nil {
		logger.Printf("encountered error: %v", err)
		return 1
	}

	// Algorithm lines
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		str := handler(line)
		if str == nil {
			continue
		}
		if _, err := io.WriteString(out, fmt.Sprintln(*str)); err != nil {
			logger.Printf("encountered error: %v", err)
			return 1
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Printf("encountered error: %v", err)
		return 1
	}

	return 0
}

func selectAlgorithm(args []string, handlers []pkg.LineHandlerDescriptor) (pkg.LineHandler, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf(usage(handlers))
	}
	opt := args[0]
	for _, handler := range handlers {
		if strings.ToUpper(handler.Name) == strings.ToUpper(opt) {
			return handler.LineHandler, nil
		}
	}
	return nil, fmt.Errorf(usage(handlers))
}

func usage(handlers []pkg.LineHandlerDescriptor) string {
	usage := "You must provide one argument for the program to use. Available algorithms:\n\n"
	for _, handler := range handlers {
		usage += fmt.Sprintf("%s\t%s\n", handler.Name, handler.Description)
	}
	return usage
}
