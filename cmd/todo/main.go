package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/surprise213/todo/internal/storage"
)

func main() {
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	var s storage.Storage
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
		t, err := s.Add(task)
		if err != nil {
			return err
		}
		fmt.Printf("добавлена задача %d %s", t.ID, t.Title+"\n")

	case "list":
		if len(args) != 1 {
			return errors.New("использование: list")
		}
		var builder strings.Builder
		tasks := s.List()
		if len(tasks) == 0 {
			fmt.Println("список пуст")
			return nil
		}
		for _, t := range tasks {
			mark := "[ ]"
			if t.Done {
				mark = "[x]"
			}
			fmt.Fprintf(&builder, "%s %d. %s\n", mark, t.ID, t.Title)
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
		err = s.MarkDone(id)
		if err != nil {
			return err
		}

	case "delete":
		if len(args) != 2 {
			return errors.New("использование: delete <id>")
		}
		id, err := strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("неверный id %q: %w", args[1], err)
		}
		err = s.Delete(id)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("неизвестная команда %q", args[0])
	}
	return nil
}
