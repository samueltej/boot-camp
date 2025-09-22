package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	countLines := flag.Bool("l", false, "count lines")
	countBytes := flag.Bool("b", false, "count lines")
	flag.Parse()

	reader := bufio.NewReader(os.Stdin)

	if *countLines {
		fmt.Println("Line count mode. Write your text (one or several lines). Type 'exit' to finish:")
	} else if *countBytes {
		fmt.Println("Byte count mode. Write your text (one or several lines). Type 'exit' to finish:")
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
	result := modeCounter(input, *countLines, *countBytes)

	if *countLines {
		fmt.Println("Number of Lines:", result)
	} else if *countBytes {
		fmt.Println("Number of Bytes", result)
	} else {
		fmt.Println("Number of words:", result)
	}
}

func modeCounter(text string, countLines, countBytes bool) int {
	if countLines {
		lines := strings.Split(text, "\n")
		return len(lines)
	} else if countBytes {
		noNewLines := strings.ReplaceAll(text, "\r", "")
		noNewLines = strings.ReplaceAll(noNewLines, "\n", "")
		return len(noNewLines)
	} else {
		cleanWord := strings.ReplaceAll(text, "-", " ")
		words := strings.Fields(cleanWord)
		return len(words)
	}
}
