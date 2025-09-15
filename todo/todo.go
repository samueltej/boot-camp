package todo

import (
	"time"
	"errors"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (ls *List) AddList(task string) {
	i := item{
		Task:      task,
		Done:      false,
		CreatedAt: time.Now(),
	}
	*ls = append(*ls, i)
}
func (ls *List) CompleteTask(index int) error {
	if index < 0 || index >= len(*ls) {
		return errors.New("index don't exists")
	}
	(*ls)[index].Done = true
	(*ls)[index].CompletedAt = time.Now()
	return nil
}

func (ls *List) DeleteTask(index int) error {
	if index < 0 || index >= len(*ls) {
		return errors.New("index don't exists")
	}

	*ls = append((*ls)[:index], (*ls)[index+1:]...)
	return nil
}
