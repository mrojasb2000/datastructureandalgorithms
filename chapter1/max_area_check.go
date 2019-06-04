package main

import "fmt"

func maxAreaCheck(lenght, width, limit int) bool {
	if area := lenght * width; area < limit {
		return true
	}
	return false
}

func main() {
	fmt.Println(maxAreaCheck(12, 3, 4))
}
