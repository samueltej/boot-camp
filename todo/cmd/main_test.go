package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".test_todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)
	err := build.Run()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot build tool %s: %s", binName, err)
		os.Exit(1)
	}

	err = os.WriteFile(fileName, []byte{}, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot create file %s", fileName)
		os.Exit(1)
	}

	fmt.Println("Running tests....")
	result := m.Run()

	fmt.Println("Cleaning up....")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "New Task"
	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-task", task)
		fmt.Println(cmd)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		fmt.Println(cmd)
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := "[ ] 0: " + task
		if !strings.Contains(string(out), expected) {
			t.Errorf("expected output to contain '%s', got %s instead", expected, string(out))
		}
	})

	t.Run("CompleteTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		fmt.Println(cmd)
		err := cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
		cmd = exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		expected := "[X] 0: " + task
		if !strings.Contains(string(out), expected) {
			t.Errorf("expected output to contain completed task '%s', got %s instead", expected, string(out))
		}
	})

	t.Run("DeleteTask", func(t *testing.T) {
		addCmd := exec.Command(cmdPath, "-task", "Task to delete")
		err := addCmd.Run()
		if err != nil {
			t.Fatal(err)
		}
		cmd := exec.Command(cmdPath, "-delete", "1")
		fmt.Println(cmd)
		err = cmd.Run()
		if err != nil {
			t.Fatal(err)
		}
	})
}