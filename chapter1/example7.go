package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	isEven(rand.Intn(100))
}

func isEven(value int) {
	switch {
	case value%2 == 0:
		fmt.Printf("%d is even", value)
	default:
		fmt.Printf("%d is old", value)
	}
}
