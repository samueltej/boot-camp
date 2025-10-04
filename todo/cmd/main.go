package main

import (
	"fmt"
	"os"
	"strings"
	"todo"
)

const fileName = ".todo.json"

func main() {
	var list todo.List

	err := list.Get(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading file: %v\n", err)
		os.Exit(1)
	}
	
	args := os.Args[1:]

	if len(args) == 0 {
		for _, item := range list {
			fmt.Println(item.Task)
		}
		return
	}

	task := strings.Join(args, " ")
	list.AddTask(task)

	err = list.Save(fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error saving file: %v\n", err)
		os.Exit(1)
	}
}
