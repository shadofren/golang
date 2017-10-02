package main

import (
	"./sortings" // Local package import
	// "./DS"
	// "./others"
	// "./search"
)

func main() {
	list := []int {9, 10, 8, 7, 6, 5, 4, 3, 2, 1};
	sortings.TestHeapSort(list[:]);
	// sortings.TestQuickSort(list[:]);
	// DS.TestLinkedList();
	// DS.TestStack();
	// DS.TestQueue();
	// others.TestPermutation();
	// search.TestBinarySearchI();
	// search.TestBinarySearchR();
}




