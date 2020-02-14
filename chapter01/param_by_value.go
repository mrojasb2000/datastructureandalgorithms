package main

import (
	"fmt"
)

// IncrementPassByValue test pass parameters
func IncrementPassByValue(x int) {
	fmt.Println("Incrment value...")
	x++
}

func main() {
	i := 10
	fmt.Println("Value of i before increment is:", i)
	IncrementPassByValue(i)
	fmt.Println("Value of i after increment is:", i)
}
