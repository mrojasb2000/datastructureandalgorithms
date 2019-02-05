package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(10, 20))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
