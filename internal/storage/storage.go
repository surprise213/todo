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
	tasks []Task
	nextId int
}

var ErrTitleEmpty = errors.New("Title cannot be empty")
var ErrNotFound = errors.New("not found")

func (s *Storage) Add(title string) (Task, error) {
	
	if title == "" {
		return Task{}, ErrTitleEmpty
	}
	task := Task{
		ID:    s.nextId + 1,
		Title: title,
		Done:  false,
	}
	s.tasks = append(s.tasks, task)
	s.nextId+=1
	return task, nil
}

func (s *Storage) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID==id{
			return task, nil
		}
	}
	return Task{}, ErrNotFound
}

func (s *Storage) MarkDone(id int) error{
	for i:=0;i<len(s.tasks);i++{
		if s.tasks[i].ID==id{
			s.tasks[i].Done=true
			return nil
		}
	}
	return ErrNotFound
}

func (s *Storage) List() []Task {
	var c []Task
	c=append(c, s.tasks...)
	return c
}
