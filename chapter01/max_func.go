package main

import (
	"fmt"
)

func main() {
	fmt.Println(max(10, 20))
	numbers := []int{321, 22, 555, 1, 77, 2, 8, 9, 12, 144}
	fmt.Println("Max number: ", maxNumber(numbers))

}

func maxNumber(numbers []int) int {
	var maxNumber = 0
	lenNumbers := len(numbers)
	for i := 0; i < lenNumbers; i++ {
		if pos := i + 1; pos < lenNumbers {
			//fmt.Println(numbers[i], " : ", maxNumber, "= ", max(numbers[i], maxNumber))
			maxNumber = max(numbers[i], maxNumber)
		} else {
			maxNumber = max(numbers[lenNumbers-1], maxNumber)
		}
	}
	return maxNumber
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
