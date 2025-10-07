package main

import (
	"flag"
	"fmt"
	"os"
	"todo"
)

const fileName = ".todo.json"

func main() {

	list := flag.Bool("list", false, "List incomplete tasks")
	task := flag.String("task", "", "Add a new task")
	complete := flag.Int("complete", -1, "Complete a task by its number")
	delete := flag.Int("delete", -1, "Delete a task by its number")
	flag.Parse()

	var l todo.List


	if err := l.Get(fileName); err != nil {
		fmt.Fprintf(os.Stderr, "Error loading tasks: %v\n", err)
		os.Exit(1)
	}

	if *list {
		for _, t := range l {
			if !t.Done {
				fmt.Printf("Title: %s, Done: %t, CreatedAt: %s, CompletedAt: %s\n",
					t.Task, t.Done, t.CreatedAt, t.CompletedAt)
			}
		}
		return
	}


	if *complete != -1 {
		if err := l.CompleteTask(*complete - 1); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing task: %v\n", err)
			os.Exit(1)
		}
		if err := l.Save(fileName); err != nil {
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
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *task != "" {
		l.AddTask(*task)
		if err := l.Save(fileName); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving tasks: %v\n", err)
			os.Exit(1)
		}
		return
	}

	fmt.Println("Error: you must specify at least one valid flag (-list, -task, -complete, -delete)")
	os.Exit(1)

}
