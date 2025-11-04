package todo

import (
	"os"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	var ls List
	ls.AddTask("Buy milk")

	if len(ls) != 1 {
		t.Errorf("TestAdd failed: expected list length 1, got %d", len(ls))
	}

	if ls[0].Task != "Buy milk" {
		t.Errorf("TestAdd failed: expected task %q, got %q", "Buy milk", ls[0].Task)
	}
}

func TestComplete(t *testing.T) {
	var ls List
	ls.AddTask("Call the doctor")

	err := ls.CompleteTask(0)
	if err != nil {
		t.Fatalf("TestComplete failed: unexpected error: %v", err)
	}

	if !ls[0].Done {
		t.Errorf("TestComplete failed: expected task %q to be completed", ls[0].Task)
	}
}

func TestDelete(t *testing.T) {
	var ls List
	ls.AddTask("Do the laundry")
	ls.AddTask("Take out the trash")

	err := ls.DeleteTask(0)
	if err != nil {
		t.Fatalf("TestDelete failed: unexpected error: %v", err)
	}

	if len(ls) != 1 {
		t.Errorf("TestDelete failed: expected list length 1, got %d", len(ls))
	}

	if ls[0].Task != "Take out the trash" {
		t.Errorf("TestDelete failed: expected remaining task to be %q, got %q", "Take out the trash", ls[0].Task)
	}
}

func TestSaveAndGet(t *testing.T) {
	var ls List
	ls.AddTask("Pay bills")

	tf, err := os.CreateTemp("", "todo_test_*.json")
	if err != nil {
		t.Fatalf("TestSaveAndGet failed: could not create temp file: %v", err)
	}
	defer os.Remove(tf.Name())

	if err := ls.Save(tf.Name()); err != nil {
		t.Fatalf("TestSaveAndGet failed: could not save list: %v", err)
	}

	var newList List
	if err := newList.Get(tf.Name()); err != nil {
		t.Fatalf("TestSaveAndGet failed: could not get list: %v", err)
	}

	if len(newList) != 1 {
		t.Errorf("TestSaveAndGet failed: expected list length 1, got %d", len(newList))
	}

	if newList[0].Task != "Pay bills" {
		t.Errorf("TestSaveAndGet failed: expected task %q, got %q", "Pay bills", newList[0].Task)
	}
}

func TestString(t *testing.T) {
	var ls List
	ls.AddTask("Task 1")
	ls.AddTask("Task 2")

	ls.CompleteTask(1)

	result := ls.String()

	if !strings.Contains(result, "[ ] 0: Task 1") {
		t.Errorf("TestString failed: expected incomplete task format '[ ] 0: Task 1', got:\n%s", result)
	}

	if !strings.Contains(result, "[X] 1: Task 2") {
		t.Errorf("TestString failed: expected complete task format '[X] 1: Task 2', got:\n%s", result)
	}
}
