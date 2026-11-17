package storage

import (
	"errors"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type Storage struct {
	tasks  []Task
	nextID int
}

var ErrTitleEmpty = errors.New("title cannot be empty")
var ErrNotFound = errors.New("task not found")

func (s *Storage) Add(title string) (Task, error) {

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

func (s *Storage) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, ErrNotFound
}

func (s *Storage) MarkDone(id int) error {
	for i := 0; i < len(s.tasks); i++ {
		if s.tasks[i].ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return ErrNotFound
}

func (s *Storage) List() []Task {
	var c []Task
	c = append(c, s.tasks...)
	return c
}

func (s *Storage) Delete(id int) error {
	for i := range s.tasks {
		if s.tasks[i].ID == id {
			s.tasks = append(s.tasks[:i], s.tasks[i+1:]...)
			return nil
		}
	}
	return ErrNotFound
}
