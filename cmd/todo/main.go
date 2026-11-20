package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/surprise213/todo/internal/storage"
)

type Repository interface {
	Add(title string) (storage.Task, error)
	List() []storage.Task
	MarkDone(id int) error
	Delete(id int) error
}

var _ Repository = (*storage.MemoryStorage)(nil)
var _ Repository = (*storage.FileStorage)(nil)

func main() {
	repo, err := storage.NewFileStorage("tasks.json")
	if err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}

	if err := run(os.Args[1:], repo); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		if errors.Is(err, storage.ErrNotFound) {
			os.Exit(2)
		}
		os.Exit(1)
	}
}

func run(args []string, repo Repository) error {
	if len(args) == 0 {
		return errors.New("не указана команда\nиспользование: add, list, done, delete")
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			return errors.New("использование: add <задание>")
		}
		task := strings.Join(args[1:], " ")
		task = strings.TrimSpace(task)
		t, err := repo.Add(task)
		if err != nil {
			return fmt.Errorf("add: %w", err)
		}
		fmt.Printf("добавлена задача %d %s", t.ID, t.Title+"\n")

	case "list":
		if len(args) != 1 {
			return errors.New("использование: list")
		}
		var builder strings.Builder
		tasks := repo.List()
		if len(tasks) == 0 {
			fmt.Println("список пуст")
			return nil
		}
		for _, t := range tasks {
			fmt.Fprintln(&builder, t)
		}
		res := builder.String()
		fmt.Print(res)

	case "done":
		if len(args) != 2 {
			return errors.New("использование: done <id>")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("неверный id %q: %w", args[1], err)
		}
		err = repo.MarkDone(id)
		if err != nil {
			return fmt.Errorf("done: %w", err)
		}

	case "delete":
		if len(args) != 2 {
			return errors.New("использование: delete <id>")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("неверный id %q: %w", args[1], err)
		}
		err = repo.Delete(id)
		if err != nil {
			return fmt.Errorf("delete: %w", err)
		}
	default:
		return fmt.Errorf("неизвестная команда %q", args[0])
	}
	return nil
}
