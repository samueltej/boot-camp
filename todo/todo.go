package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (ls *List) AddTask(task string) {
	i := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}
	*ls = append(*ls, i)
}

func (ls *List) CompleteTask(index int) error {
	if index < 0 || index >= len(*ls) {
		return fmt.Errorf("invalid index %d", index)
	}
	(*ls)[index].Done = true
	(*ls)[index].CompletedAt = time.Now()
	return nil
}

func (ls *List) DeleteTask(index int) error {
	if index < 0 || index >= len(*ls) {
		return fmt.Errorf("invalid index %d", index)
	}

	*ls = append((*ls)[:index], (*ls)[index+1:]...)
	return nil
}

func (ls *List) Save(filename string) error {
	data, err := json.Marshal(ls)
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (ls *List) Get(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        if os.IsNotExist(err) {
            *ls = List{} 
            return nil
        }
        return err
    }
    if len(data) == 0 {
        *ls = List{} 
    }

    err = json.Unmarshal(data, ls)
    if err != nil {
        return err
    }

    return nil
}

func (ls *List) String() string {
	var formatted strings.Builder
	
	for i, task := range *ls {
		prefix := "[ ]"
		if task.Done {
			prefix = "[X]"
		}
		fmt.Fprintf(&formatted, "%s %d: %s\n", prefix, i, task.Task)
	}
	
	return formatted.String()
}