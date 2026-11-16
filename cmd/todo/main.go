package main

import(
	 "fmt"
	 "github.com/surprise213/todo/internal/storage"
)

func main() {
	var s storage.Storage
	s.Add("Golang")
	s.Get(1)
	s.MarkDone(1)
	fmt.Println(s.List())

}