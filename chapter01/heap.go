package main

import (
	"container/heap"
	"fmt"
)

// IntegerHeap struct
type IntegerHeap []int

// Len gets the length of integerHeap
func (iheap IntegerHeap) Len() int {
	return len(iheap)
}

// Less checks if element of i index is less than j index
func (iheap IntegerHeap) Less(i, j int) bool {
	return iheap[i] < iheap[j]
}

// Swap the element of i and j index
func (iheap IntegerHeap) Swap(i, j int) {
	iheap[i], iheap[j] = iheap[j], iheap[i]
}

// Push - pushes the item
func (iheap *IntegerHeap) Push(heapintf interface{}) {
	*iheap = append(*iheap, heapintf.(int))
}

// Pop - pops the item from the heap
func (iheap *IntegerHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerHeap = *iheap
	n = len(previous)
	x1 = previous[n-1]
	*iheap = previous[0 : n-1]
	return x1
}

func main() {
	var intHeap *IntegerHeap = &IntegerHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	heap.Push(intHeap, 7)
	heap.Push(intHeap, 3)
	fmt.Printf("minimun: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}
}
