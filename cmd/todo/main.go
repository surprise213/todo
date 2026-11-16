package main

import(
	 "fmt"
	 "errors"
	 "github.com/surprise213/todo/internal/storage"
)

func main() {
	var s storage.Storage
	_,err:=s.Add("Golang")
	if errors.Is(err, storage.ErrTitleEmpty){
		fmt.Println(storage.ErrTitleEmpty)
	}

	_,err=s.Get(1)
	if errors.Is(err, storage.ErrTitleEmpty){
		fmt.Println(storage.ErrTitleEmpty)
	}
	err=s.MarkDone(1)
	if errors.Is(err, storage.ErrTitleEmpty){
		fmt.Println(storage.ErrTitleEmpty)
	}
	fmt.Println(s.List())

}