package storage

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

type MemoryStorage struct {
	tasks  []Task
	nextID int
}

var ErrTitleEmpty = errors.New("title cannot be empty")
var ErrNotFound = errors.New("task not found")

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{}
}

func (s *MemoryStorage) Add(title string) (Task, error) {

	if title == "" {
		return Task{}, ErrTitleEmpty
	}
	task := Task{
		ID:    s.nextID + 1,
		Title: title,
		Done:  false,
	}
	s.tasks = append(s.tasks, task)
	s.nextID += 1
	return task, nil
}

func (s *MemoryStorage) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, ErrNotFound
}

func (s *MemoryStorage) MarkDone(id int) error {
	for i := 0; i < len(s.tasks); i++ {
		if s.tasks[i].ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return ErrNotFound
}

func (s *MemoryStorage) List() []Task {
	var c []Task
	c = append(c, s.tasks...)
	return c
}

func (s *MemoryStorage) Delete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}

func (t Task) String() string {
	var builder strings.Builder
	mark := "[ ]"
	if t.Done {
		mark = "[x]"
	}
	fmt.Fprintf(&builder, "%s %d. %s", mark, t.ID, t.Title)
	res := builder.String()
	return res
}
