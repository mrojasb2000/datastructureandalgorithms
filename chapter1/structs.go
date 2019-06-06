package main

import (
	"fmt"
)

type student struct {
	rollNo int
	name   string
}

func main() {
	stud := student{1, "johnny"}
	fmt.Println(stud)
	// Accessing inner fields
	fmt.Println("Student name ::", stud.name)
	pstud := &stud
	// Accessing inner fields
	fmt.Println("Student name ::", pstud.name)
	// Named initialization
	fmt.Println(student{rollNo: 2, name: "Ann"})
	// Order does not matter
	fmt.Println(student{name: "Ann", rollNo: 2})
	// Default initialization of rollNo
	fmt.Println(student{name: "Alice"})
}
