package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	countLines, countBytes := parseFlags()
	reader := bufio.NewReader(os.Stdin)

	printModeMessage(countLines, countBytes)

	input := readInput(reader)

	result := modeCounter(input, countLines, countBytes)
	printResult(result, countLines, countBytes)
}

func parseFlags() (bool, bool) {
	countLines := flag.Bool("l", false, "count lines")
	countBytes := flag.Bool("b", false, "count bytes")
	flag.Parse()
	return *countLines, *countBytes
}

func printModeMessage(countLines, countBytes bool) {
	if countLines {
		fmt.Println("Line count mode. Write your text (one or several lines). Type 'exit' to finish:")
	} else if countBytes {
		fmt.Println("Byte count mode. Write your text (one or several lines). Type 'exit' to finish:")
	} else {
		fmt.Println("Word count mode. Write your text (one or several lines). Type 'exit' to finish:")
	}
}

func readInput(reader *bufio.Reader) string {
	var inputBuilder strings.Builder

	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error while reading:", err)
			return ""
		}
		if isExit(text) {
			break
		}
		inputBuilder.WriteString(text)
	}

	return inputBuilder.String()
}

func isExit(text string) bool {
	exitOptions := []string{"exit", "EXIT", "Exit"}
	for _, option := range exitOptions {
		if strings.TrimSpace(text) == option {
			return true
		}
	}
	return false
}

func printResult(result int, countLines, countBytes bool) {
	if countLines {
		fmt.Println("Number of Lines:", result)
	} else if countBytes {
		fmt.Println("Number of Bytes", result)
	} else {
		fmt.Println("Number of words:", result)
	}
}

func modeCounter(text string, countLines, countBytes bool) int {
	if countLines {
		lines := strings.Split(text, "\n")
		if strings.HasSuffix(text, "\n") {
			return len(lines) - 1
		}
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
