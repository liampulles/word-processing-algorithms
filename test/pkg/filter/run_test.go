package filter_test

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/liampulles/word-processing-algorithms/pkg/filter"
)

func TestRun_GivenSomeArgsAndInput_ShouldReturnAndOutputAsExpected(t *testing.T) {
	// Setup fixture
	var tests = []struct {
		argsFixture    []string
		inFixture      string
		expectedResult int
		expectedOutput string
	}{
		// --- Failing cases ---
		// Empty case -> error (no args)
		{
			[]string{},
			lines(),
			1,
			"",
		},
		// Empty args case -> error
		{
			[]string{},
			lines("some", "words"),
			1,
			"",
		},
		// Too many args case -> error
		{
			[]string{"palindrome", "palindrome"},
			lines("some", "words"),
			1,
			"",
		},
		// Invalid arg case -> error
		{
			[]string{"does.not.exist"},
			lines("some", "words"),
			1,
			"",
		},

		// --- Palindrome cases ---
		// No input -> No output
		{
			[]string{"palindrome"},
			lines(),
			0,
			"",
		},
		// One blank word -> No output
		{
			[]string{"palindrome"},
			lines(""),
			0,
			"",
		},
		// One non-palindrome -> No output
		{
			[]string{"palindrome"},
			lines("word"),
			0,
			"",
		},
		// One palindrome -> Palindrome out
		{
			[]string{"palindrome"},
			lines("cIvic"),
			0,
			linesTerminated("cIvic"),
		},
		// Many blanks -> No output
		{
			[]string{"palindrome"},
			lines("", ""),
			0,
			"",
		},
		// Many non-palindromes -> No output
		{
			[]string{"palindrome"},
			lines("cat", "dog"),
			0,
			"",
		},
		// Many palindromes -> Same output
		{
			[]string{"palindrome"},
			lines("a", "dEed", "cIvic", "kayAk", "Don’t nod"),
			0,
			linesTerminated("a", "dEed", "cIvic", "kayAk", "Don’t nod"),
		},
		// Mixed case
		{
			[]string{"palindrome"},
			lines("", "cIvic", "dog", "Don’t nod", "", "kayAk", "man", ""),
			0,
			linesTerminated("cIvic", "Don’t nod", "kayAk"),
		},

		// --- Abecedarian cases ---
		// No input -> No output
		{
			[]string{"abecedarian"},
			lines(),
			0,
			"",
		},
		// One blank word -> No output
		{
			[]string{"abecedarian"},
			lines(""),
			0,
			"",
		},
		// One non-abecedarian word -> No output
		{
			[]string{"abecedarian"},
			lines("cat"),
			0,
			"",
		},
		// One abecedarian word -> Copy to output
		{
			[]string{"abecedarian"},
			lines("ad"),
			0,
			linesTerminated("ad"),
		},
		// Many blanks -> No output
		{
			[]string{"abecedarian"},
			lines("", ""),
			0,
			"",
		},
		// Many non-abecedarian -> No output
		{
			[]string{"abecedarian"},
			lines("cat", "dog"),
			0,
			"",
		},
		// Many abecedarian words -> Copy all to putput
		{
			[]string{"abecedarian"},
			lines("ad", "aLmo st"),
			0,
			linesTerminated("ad", "aLmo st"),
		},
		// Mixed case
		{
			[]string{"abecedarian"},
			lines("", "ad", "dog", "AlMo st", "", "kayak", "man", ""),
			0,
			linesTerminated("ad", "AlMo st"),
		},

		// --- Distinct cases ---
		// No input -> No output
		{
			[]string{"distinct"},
			lines(),
			0,
			"",
		},
		// One blank word -> No output
		{
			[]string{"distinct"},
			lines(""),
			0,
			"",
		},
		// One non-distinct word -> No output
		{
			[]string{"distinct"},
			lines("deed"),
			0,
			"",
		},
		// One distinct word -> Copy to output
		{
			[]string{"distinct"},
			lines("cAt"),
			0,
			linesTerminated("cAt"),
		},
		// Many blank words -> No output
		{
			[]string{"distinct"},
			lines("", ""),
			0,
			"",
		},
		// Many non-distinct words -> No output
		{
			[]string{"distinct"},
			lines("deed", " mom"),
			0,
			"",
		},
		// Many distinct words -> Copy to output
		{
			[]string{"distinct"},
			lines("cAt", "Dog"),
			0,
			linesTerminated("cAt", "Dog"),
		},
		// Mixed case
		{
			[]string{"distinct"},
			lines("", "deed", "cAt", "mom", "", "Dog"),
			0,
			linesTerminated("cAt", "Dog"),
		},
	}

	for i, test := range tests {
		t.Run(fmt.Sprintf("[%d] (%s -> %v -> %s)", i, test.inFixture, test.argsFixture, test.expectedOutput), func(t *testing.T) {
			// Setup fixture
			inFixtureReader := strings.NewReader(test.inFixture)
			outFixtureWriter := bytes.NewBufferString("")

			// Exercise SUT
			actualResult := filter.Run(test.argsFixture, inFixtureReader, outFixtureWriter)

			// Verify result
			if actualResult != test.expectedResult {
				t.Errorf("Unexpected result.\nExpected: %d\nActual: %d", test.expectedResult, actualResult)
			}
			actualOutput := outFixtureWriter.String()
			if actualOutput != test.expectedOutput {
				t.Errorf("Unexpected result.\nExpected: %s\nActual: %s", test.expectedOutput, actualOutput)
			}
		})
	}
}

func lines(in ...string) string {
	return strings.Join(in, "\n")
}

func linesTerminated(in ...string) string {
	return fmt.Sprintln(lines(in...))
}
