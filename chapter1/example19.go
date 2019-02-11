package main

import (
	"fmt"
)

type MyInt int

func (data MyInt) incremental1() {
	data = data + 1
}

func (data *MyInt) incremental2() {
	*data = *data + 1
}

func main() {
	var data MyInt = 1
	fmt.Println("Value before incremental1() call:", data)
	data.incremental1()
	fmt.Println("Value before incremental1() call:", data)
	data.incremental2()
	fmt.Println("Value before incremental2() call:", data)
}
