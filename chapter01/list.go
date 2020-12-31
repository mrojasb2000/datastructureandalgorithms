package main

import (
	"container/list"
	"fmt"
)

func main() {
	var intlist list.List

	fmt.Println("FIFO")
	// FIFO
	intlist.PushBack(11)
	intlist.PushBack(23)
	intlist.PushBack(34)

	for element := intlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}

	// Clean list
	intlist.Init()
	fmt.Println("LIFO")

	// LIFO
	intlist.PushFront(11)
	intlist.PushFront(23)
	intlist.PushFront(34)

	for element := intlist.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}
