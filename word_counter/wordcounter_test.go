package main

import (
	"testing"
)

func TestWordCounter(t *testing.T) {
	cases := []struct {
		test string
		text string
		want int
	}{
		{"TestSingleSentence", "This is a test", 4},
		{"TestMultipleSentences", "Go is fun. Testing is important.", 6},
		{"TestSingleWord", "hello", 1},
		{"TestComposedWord", "read-only", 2},
		{"TestBreakWords", "\n\n word", 1},
		{"TestWordExit", "Exit", 1}, //
	}

	for _, c := range cases {
		got := modeCounter(c.text, false, false)
		if got != c.want {
			t.Errorf("Fail in %s: expected [%d], got [%d]", c.test, c.want, got)
		}
	}
}

func TestLineCounter(t *testing.T) {
	cases := []struct {
		test string
		text string
		want int
	}{
		{"TestSingleLine", "This is one line", 1},
		{"TestMultipleLines", "First line\nSecond line\nThird line", 3},
		{"TestBreakLines", "Line1\n\nLine2", 3},
		{"TestExitLines", "Hello\nEXIT\nBye", 3},
	}

	for _, c := range cases {
		got := modeCounter(c.text, true, false)
		if got != c.want {
			t.Errorf("Fail in %s: expected [%d], got [%d]", c.test, c.want, got)
		}
	}
}

func TestByteCounter(t *testing.T) {
	cases := []struct {
		test string
		text string
		want int
	}{
		{"TestEmptyString", "", 0},
		{"TestSimpleText", "GoLang", 6},
		{"TestSentence", "Hello World!", 12},
		{"TestWithSpaces", " count bytes ", 13},
		{"TestMultilineBytes", "one line\nsecond line", 19},
	}

	for _, c := range cases {
		got := modeCounter(c.text, false, true)
		if got != c.want {
			t.Errorf("Fail in %s: expected [%d], got [%d]", c.test, c.want, got)
		}
	}
}
