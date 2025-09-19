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

	var inputBuilder strings.Builder

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading:", err)
			return
		}
		if strings.TrimSpace(text) == "exit" || strings.TrimSpace(text) == "EXIT" || strings.TrimSpace(text) == "Exit" {
			break
		}
		inputBuilder.WriteString(text)
	}

	input := inputBuilder.String()
	result := wordCounter(input, lineCountMode)

	if lineCountMode == "-l" {
		fmt.Println("Number of lines:", result)
	} else {
		fmt.Println("Number of words:", result)
	}
}

func wordCounter(text string, mode string) int {
	if mode == "-l" {
		lines := strings.Split(strings.TrimSpace(text), "\n")
		return len(lines)
	} else {
		cleanWord := strings.ReplaceAll(text, "-", " ")
		words := strings.Fields(cleanWord)
		return len(words)
	}
}
