package storage

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

type Storage struct {
	tasks []Task
}

func (s *Storage) Add(title string) (Task, error) {
	if title == "" {
		return Task{}, fmt.Errorf("Title cannot be empty")
	}
	task := Task{
		ID:    len(s.tasks) + 1,
		Title: title,
		Done:  false,
	}
	s.tasks = append(s.tasks, task)
	return task, nil
}

func (s Storage) Get(id int) (Task, error) {
	for _, task := range s.tasks {
		if task.ID==id{
			return task, nil
		}
	}
	return Task{}, fmt.Errorf("Not found")
}

func (s *Storage) MarkDone(id int) error{
	for _, task := range s.tasks{
		if task.ID==id{
			task.Done=true
			return nil
		}
	}
	return fmt.Errorf("Not found")
}

func (s Storage) List() []Task {
	return s.tasks
}
