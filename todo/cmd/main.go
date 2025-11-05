package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const defaultFileName = ".todo.json"

func main() {
	filename := getFileName()

	list := flag.Bool("list", false, "List incomplete tasks")
	task := flag.String("task", "", "Add a new task")
	complete := flag.Int("complete", -1, "Complete a task by its number")
	delete := flag.Int("delete", -1, "Delete a task by its number")
	flag.Parse()

	var l todo.List

	if err := l.Get(filename); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if *list {
		fmt.Print(&l)
		return
	}

	if *complete != -1 {
		if err := l.CompleteTask(*complete - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *delete != -1 {
		if err := l.DeleteTask(*delete - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *task != "" {
		l.AddTask(*task)
		if err := l.Save(filename); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("Error: you must specify at least one valid flag (-list, -task, -complete, -delete)")
	os.Exit(1)
}


func getFileName() string {
	if envFile := os.Getenv("TODO_FILENAME"); envFile != "" {
		return envFile
	}
	return defaultFileName
}