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

	wordCount := 0
	lineCount := 0
	isFinished := false

	for !isFinished {
		text, _ := reader.ReadString('\n')
		sentence := strings.TrimSpace(text)

		if sentence == "exit" {
			isFinished = true
			continue
		}

		if lineCountMode == "-l" {
			lineCount++
		} else {
			words := strings.Fields(text)
			wordCount += len(words) 
		}
	}

	if lineCountMode == "-l" {
		fmt.Println("Número de líneas:", lineCount)
	} else {
		fmt.Println("Número de palabras:", wordCount)
	}
}