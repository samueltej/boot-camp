package main

import "testing"

// Test Words

func TestSingleSentence(t *testing.T) {
	result := wordCounter("This is a test", "")
	want := 4
	if result != want {
		t.Errorf("Error en TestSingleSentence: esperado %d, obtenido %d", want, result)
	}
}

func TestMultipleSentences(t *testing.T) {
	result := wordCounter("Go is fun. Testing is important.", "")
	want := 6
	if result != want {
		t.Errorf("Error en TestMultipleSentences: esperado %d, obtenido %d", want, result)
	}
}

func TestSingleWord(t *testing.T) {
	result := wordCounter("Hello", "")
	want := 1
	if result != want {
		t.Errorf("Error en TestSingleWord: esperado %d, obtenido %d", want, result)
	}
}

func TestComposedWord(t *testing.T) {
	result := wordCounter("read-only", "")
	want := 2
	if result != want {
		t.Errorf("Error en TestComposedWord: esperado %d, obtenido %d", want, result)
	}
}

func TestBreakWords(t *testing.T) {
	result := wordCounter("\n\n frase", "")
	want := 1
	if result != want {
		t.Errorf("Error en TestBreakWords: esperado %d, obtenido %d", want, result)
	}
}

func TestWordExit(t *testing.T) {
	result := wordCounter("EXIT", "") //en main se valida
	want := 1
	if result != want {
		t.Errorf("Error en TestWordExit: esperado %d, obtenido %d", want, result)
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
	text := "First line\nSecond line\nThird line"

	want := 3
	got := wordCounter(text, "-l")

	if got != want {
		t.Errorf("Error in TestMultipleLines: expected %d, got %d", want, got)
	}
}

func TestBreakLines(t *testing.T) {
	text := "Line1\n\nLine2"

	want := 3
	got := wordCounter(text, "-l")

	if got != want {
		t.Errorf("Error in TestBreakLines: expected %d, got %d", want, got)
	}
}

func TestExitLines(t *testing.T) {
	text := "Hello\nEXIT\nBye"

	want := 3 // EXIT es validado en main
	got := wordCounter(text, "-l")

	if got != want {
		t.Errorf("Error in TestExitLines: expected %d, got %d", want, got)
	}
}
