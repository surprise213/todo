package main

import (
	"errors"
	"fmt"

	"github.com/surprise213/todo/internal/storage"
)

func HandleErrTitleEmpty(err error) {
	if errors.Is(err, storage.ErrTitleEmpty) {
		fmt.Println(storage.ErrTitleEmpty)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}
}
func HandleErrNotFound(err error) {
	if errors.Is(err, storage.ErrNotFound) {
		fmt.Println(storage.ErrNotFound)
	} else if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {
	var s storage.Storage
	_, err := s.Add("Golang")
	HandleErrTitleEmpty(err)
	_, err = s.Add("C++")
	HandleErrTitleEmpty(err)
	_, err = s.Add("C#")
	HandleErrTitleEmpty(err)

	_, err = s.Get(1)
	HandleErrNotFound(err)

	err = s.MarkDone(1)
	HandleErrNotFound(err)

	_, err = s.Get(42)
	HandleErrNotFound(err)

	err = s.Delete(2)
	HandleErrNotFound(err)

	fmt.Println(s.List())

	err = s.Delete(2)
	HandleErrNotFound(err)

	fmt.Println(s.List())

	_, err = s.Add("Python")
	HandleErrTitleEmpty(err)

	fmt.Println(s.List())
}
