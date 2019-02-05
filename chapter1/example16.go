package main

import (
	"fmt"
)

func main() {
	i := 10
	fmt.Println("Value of i before increment is:", i)
	IncrementPassByPointer(&i)
	fmt.Println("Value of i after increment is:", i)
}

func IncrementPassByPointer(ptr *int) {
	(*ptr)++
}
