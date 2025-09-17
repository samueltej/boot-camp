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
		fmt.Println("Line count mode. Write your text (one or several lines). Type 'exit' to finish:")
	} else {
		fmt.Println("Word count mode. Write your text (one or several lines). Type 'exit' to finish:")
	}

	contador := 0
	isFinished := false

	for !isFinished {

		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading:", err)
			return
		}
		sentence := strings.TrimSpace(text)

		if sentence == "Exit" || sentence == "exit" || sentence == "EXIT" {
			isFinished = true
			continue
		}

		result := wordCounter(sentence, lineCountMode)
		contador += result
	}

	if lineCountMode == "-l" {
		fmt.Println("Number of lines:", contador)
	} else {
		fmt.Println("Number of words:", contador)
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
