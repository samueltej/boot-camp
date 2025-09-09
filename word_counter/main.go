package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	var lineCountMode string
	if len(os.Args) > 1 {
		lineCountMode = os.Args[1]
	} else {
		lineCountMode = ""
	}
	if lineCountMode == "-l" {
		fmt.Println("Modo conteo de líneas. Escribe tu texto (una o varias líneas). Escribe 'exit' para terminar:")
	} else {
		fmt.Println("Modo conteo de palabras. Escribe tu texto (una o varias líneas). Escribe 'exit' para terminar:")
	}

	contador := 0
	isFinished := false

	for !isFinished {

		text, _ := reader.ReadString('\n')
		sentence := strings.TrimSpace(text)

		if sentence == "Exit" || sentence == "exit" ||sentence == "EXIT" { 
			isFinished = true 
			continue
		 }

		result := wordCounter(sentence, lineCountMode)
		contador += result
	}

	if lineCountMode == "-l" {
		fmt.Println("Número de líneas:", contador)
	} else {
		fmt.Println("Número de palabras:", contador)
	}
}

func wordCounter(word string, mode string) int {
	
	if mode == "-l" {
		return 1
	} else {
		cleanWord := strings.ReplaceAll(word, "-", " ")
		words := strings.Fields(cleanWord)
		return len(words)
	}

}
