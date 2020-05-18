package pkg

// WordFilter performs some check on a word and returns true if it passes,
// false otherwise
type WordFilter func(string) bool

// LineHandler takes a string and process it in some way to optionally return a result
type LineHandler func(string) *string

// LineHandlerDescriptor adds metadata to a LineHandler
type LineHandlerDescriptor struct {
	Name        string
	Description string
	LineHandler LineHandler
}
