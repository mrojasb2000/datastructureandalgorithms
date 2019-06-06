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

// IncrementPassByPointer calling
func IncrementPassByPointer(ptr *int) {
	fmt.Println("-----------------------")
	fmt.Println("Address::", ptr)
	fmt.Println("Values::", *ptr)
	fmt.Println("-----------------------")
	(*ptr)++
	fmt.Println("Address::", ptr)
	fmt.Println("Values::", *ptr)
	fmt.Println("-----------------------")
}
