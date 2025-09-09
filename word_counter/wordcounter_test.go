
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
	result := wordCounter("\n\n", "")
	want := 0
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
		t.Errorf("Error en TestSingleLine: esperado %d, obtenido %d", want, result)
	}
}

func TestMultipleLines(t *testing.T) {
	lines := []string{
		"Primera línea",
		"Segunda línea",
		"Tercera línea",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error en TestMultipleLines: esperado %d, obtenido %d", want, got)
	}
}

func TestBreakLines(t *testing.T) {
	lines := []string{
		"Line1",
		"",  //en main lo valida
		"Line2",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error en TestBreakLines: esperado %d, obtenido %d", want, got)
	}
}

func TestExitLines(t *testing.T) {
	lines := []string{
		"Hola",
		"EXIT", //en main lo valida
		"Adiós",
	}

	want := 3
	got := 0

	for _, line := range lines {
		got += wordCounter(line, "-l")
	}

	if got != want {
		t.Errorf("Error en TestExitLines: esperado %d, obtenido %d", want, got)
	}
}
