package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type fileData struct {
	NextID int    `json:"next_id"`
	Tasks  []Task `json:"tasks"`
}
type FileStorage struct {
	path string
	mem  *MemoryStorage
}

func NewFileStorage(path string) (*FileStorage, error) {
	fs := &FileStorage{
		path: path,
		mem:  NewMemoryStorage(),
	}
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return fs, nil
	} else if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
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
	if err := os.WriteFile(f.path, data, 0644); err != nil {
		return fmt.Errorf("write %s: %w", f.path, err)
	}
	return nil
}

func (f *FileStorage) Add(title string) (Task, error) {
	t, err := f.mem.Add(title)
	if err != nil {
		return Task{}, err
	}
	if err := f.save(); err != nil {
		return Task{}, fmt.Errorf("save: %w", err)
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
		return fmt.Errorf("save %d: %w", id, err)
	}
	return nil
}
func (f *FileStorage) Delete(id int) error {
	err := f.mem.Delete(id)
	if err != nil {
		return err
	}
	if err := f.save(); err != nil {
		return fmt.Errorf("save %d: %w", id, err)
	}
	return nil
}
