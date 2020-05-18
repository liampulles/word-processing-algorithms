package main

import (
	"os"

	"github.com/liampulles/word-processing-algorithms/pkg/filter"
)

func main() {
	args := os.Args
	in := os.Stdin
	defer in.Close()
	out := os.Stdout
	defer out.Close()

	result := filter.Run(args[1:], in, out)
	os.Exit(result)
}
