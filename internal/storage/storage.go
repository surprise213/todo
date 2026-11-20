package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
type fileData struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}
type FileStorage struct {
	path string
	mem  *MemoryStorage
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
	return Task{}, fmt.Errorf("task %d: %w", id, ErrNotFound)
}

func (s *MemoryStorage) MarkDone(id int) error {
	for i := 0; i < len(s.tasks); i++ {
		if s.tasks[i].ID == id {
			s.tasks[i].Done = true
			return nil
		}
	}
	return fmt.Errorf("task %d: %w", id, ErrNotFound)
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
	return fmt.Errorf("task %d: %w", id, ErrNotFound)
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

func NewFileStorage(path string) (*FileStorage, error) {
	fs := &FileStorage{
		path: path,
		mem:  NewMemoryStorage(),
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return fs, nil
	} else {
		return nil, fmt.Errorf("error: %w", err)
	}
	fd := fileData{}
	err = json.Unmarshal(data, &fd)
	if err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	fs.mem.tasks = fd.Tasks
	fs.mem.nextID = fd.NextID
	return fs, nil
}
func (f *FileStorage) save() error {
	fd := fileData{
		NextID: f.mem.nextID,
		Tasks:  f.mem.tasks,
	}
	data, err := json.MarshalIndent(fd, "", "  ")
	if err != nil {
		return fmt.Errorf("save %s: %w", f.path, err)
	}
	return os.WriteFile(f.path, data, 0644)
}

func (f *FileStorage) Add(title string) (Task, error) {
	t, err := f.mem.Add(title)
	if err != nil {
		return Task{}, err
	}
	if err := f.save(); err != nil {
		return Task{}, fmt.Errorf("add: %w", err)
	}
	return t, nil
}
func (f *FileStorage) List() []Task {
	return f.mem.List()
}
func (f *FileStorage) MarkDone(id int) error {
	err := f.mem.MarkDone(id)
	if err != nil {
		return err
	}
	if err := f.save(); err != nil {
		return fmt.Errorf("Markdone %d: %w", id, err)
	}
	return nil
}
func (f *FileStorage) Delete(id int) error {
	err := f.mem.Delete(id)
	if err != nil {
		return err
	}
	if err := f.save(); err != nil {
		return fmt.Errorf("Delete %d: %w", id, err)
	}
	return nil
}
