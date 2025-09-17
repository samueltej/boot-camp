package main

import "testing"

// Test Words

func TestSingleSentence(t *testing.T) {
	result := wordCounter("This is a test", "")
	want := 4
	if result != want {
		t.Errorf("Error in TestSingleSentence: expected %d, got %d", want, result)
	}
}

func TestMultipleSentences(t *testing.T) {
	result := wordCounter("Go is fun. Testing is important.", "")
	want := 6
	if result != want {
		t.Errorf("Error in TestMultipleSentences: expected %d, got %d", want, result)
	}
}

func TestSingleWord(t *testing.T) {
	result := wordCounter("Hello", "")
	want := 1
	if result != want {
		t.Errorf("Error in TestSingleWord: expected %d, got %d", want, result)
	}
}

func TestComposedWord(t *testing.T) {
	result := wordCounter("read-only", "")
	want := 2
	if result != want {
		t.Errorf("Error in TestComposedWord: expected %d, got %d", want, result)
	}
}

func TestBreakWords(t *testing.T) {
	result := wordCounter("\n\n", "")
	want := 0
	if result != want {
		t.Errorf("Error in TestBreakWords: expected %d, got %d", want, result)
	}
}

func TestWordExit(t *testing.T) {
	result := wordCounter("EXIT", "") // validated in main
	want := 1
	if result != want {
		t.Errorf("Error in TestWordExit: expected %d, got %d", want, result)
	}
}

// Test Lines

func TestSingleLine(t *testing.T) {
	result := wordCounter("This is one line", "-l")
	want := 1
	if result != want {
		t.Errorf("Error in TestSingleLine: expected %d, got %d", want, result)
	}
}

func TestMultipleLines(t *testing.T) {
	lines := []string{
		"First line",
		"Second line",
		"Third line",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error in TestMultipleLines: expected %d, got %d", want, got)
	}
}

func TestBreakLines(t *testing.T) {
	lines := []string{
		"Line1",
		"", // validated in main
		"Line2",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error in TestBreakLines: expected %d, got %d", want, got)
	}
}

func TestExitLines(t *testing.T) {
	lines := []string{
		"Hello",
		"EXIT", // validated in main
		"Goodbye",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error in TestExitLines: expected %d, got %d", want, got)
	}
}
