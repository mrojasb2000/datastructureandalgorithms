package main

import (
	"fmt"
	"log"
)

func main() {
	var square int
	var cube int
	var err error
	square, cube, err = powerSeries(3)

	if err != nil {
		log.Fatal("Error power series")
	}
	fmt.Println("Square: ", square, "Cube: ", cube)
}

func powerSeries(a int) (int, int, error) {
	return a * a, a * a * a, nil
}
