package main

import (
	"errors"
	"fmt"

	"github.com/surprise213/todo/internal/storage"
)

func main() {
	var s storage.Storage
	_, err := s.Add("Golang")
	if errors.Is(err, storage.ErrTitleEmpty) {
		fmt.Println(storage.ErrTitleEmpty)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}
	_, err = s.Get(1)
	if errors.Is(err, storage.ErrNotFound) {
		fmt.Println(storage.ErrNotFound)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}

	err = s.MarkDone(1)
	if errors.Is(err, storage.ErrNotFound) {
		fmt.Println(storage.ErrNotFound)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(s.List())

	_, err = s.Get(42)
	if errors.Is(err, storage.ErrNotFound) {
		fmt.Println(storage.ErrNotFound)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}
}
