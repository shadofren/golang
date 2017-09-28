package main

import (
	"fmt"
	"./sortings" // Local package import
	"./DS"
)

func main() {
	// list := []int {9, 10, 8, 7, 6, 5, 4, 3, 2, 1};
	// testHeapSort(list[:]);
	// testQuickSort(list[:]);
	// testLinkedList();
}

func testLinkedList() {
	sentinel := DS.Node{-1, nil, nil};
	sentinel.Next = &sentinel;
	sentinel.Prev = &sentinel;

	// Test InsertFront
	var node *DS.Node;
	for i := 1; i <= 10; i++ {
		newNode := DS.Node{i, nil, nil};
		node = &newNode;
		(&sentinel).InsertFront(&newNode);
	}
	(&sentinel).Print();

	// Test Delete (Delete the last node)
	(&sentinel).Delete(node.Next.Next);
	(&sentinel).Delete(node);
	(&sentinel).Print();

	// Test InsertBack
	(&sentinel).InsertBack(&(DS.Node{11, nil, nil}));
	(&sentinel).Print();
	
	// Test Search
	fmt.Println("search 4", (&sentinel).Search(4));
	fmt.Println("search 8", (&sentinel).Search(8));
}

func testHeapSort(list []int) {
	sortings.HeapSort(list);
	fmt.Println(list);
}

func testQuickSort(list []int) {
	sortings.QuickSort(list);
	fmt.Println(list);
}

