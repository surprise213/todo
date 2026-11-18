package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/surprise213/todo/internal/storage"
)

var s = storage.Storage{}

// func HandleErr(err error) {
// 	if errors.Is(err, storage.ErrTitleEmpty) {
// 		fmt.Println(storage.ErrTitleEmpty)
// 	} else if errors.Is(err, storage.ErrNotFound) {
// 		fmt.Println(storage.ErrNotFound)
// 	} else if err != nil {
// 		fmt.Println("Error: ", err)
// 	}
// }

func main() {
	s.Add("Golang")
	if err := run(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, "ошибка:", err)
		os.Exit(1)
	}
}

func run(args []string) error {
	if len(args) == 0 {
		return errors.New("не указана команда\nиcпoльзoвaниe: add, list, done, delete")
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
		fmt.Printf("добавлена задача %x %s", t.ID, t.Title)

	case "list":
		if len(args) != 1 {
			return errors.New("использование: list")
		}
		var builder strings.Builder
		tasks := s.List()
		for i := range tasks {
			if tasks[i].Done {
				builder.WriteString("[+] " + strconv.Itoa(tasks[i].ID) + ". " + tasks[i].Title + "\n")
			} else {
				builder.WriteString("[ ] " + strconv.Itoa(tasks[i].ID) + ". " + tasks[i].Title + "\n")
			}
		}
		res := builder.String()
		fmt.Println(res)

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
			return fmt.Errorf("после delete должно быть id")
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
