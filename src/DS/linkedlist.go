package DS

import "fmt"

type Node struct {
	Val int;
	Prev *Node;
	Next *Node;
}

type LinkedList struct {
	Sentinel *Node;
}

func (list *LinkedList) InsertFront(node *Node) {
	node.Prev = list.Sentinel;
	node.Next = list.Sentinel.Next;
	list.Sentinel.Next.Prev = node;
	list.Sentinel.Next = node;
}

func (list *LinkedList) InsertBack(node *Node) {
	node.Next = list.Sentinel;
	node.Prev = list.Sentinel.Prev;
	list.Sentinel.Prev.Next = node;
	list.Sentinel.Prev = node;
}

func (list *LinkedList) Search(val int) *Node {
	node := list.Sentinel.Next;
	for node != list.Sentinel && node.Val != val {
		node = node.Next;
	}
	return node;
}

func (list *LinkedList) Delete(val int) {
	node := list.Search(val);
	node.Next.Prev = node.Prev;
	node.Prev.Next = node.Next;
}

func (list *LinkedList) Print() {
	node := list.Sentinel.Next;
	for node != list.Sentinel {
		fmt.Print(node.Val, " ");
		node = node.Next;
	}
	fmt.Println();
}

func TestLinkedList() {
	list := LinkedList{&(Node{-1, nil, nil})}
	list.Sentinel.Next = list.Sentinel;
	list.Sentinel.Prev = list.Sentinel;

	// Test InsertFront
	for i := 1; i <= 10; i++ {
		newNode := Node{i, nil, nil};
		(&list).InsertFront(&newNode);
	}
	(&list).Print();

	// Test Delete (Delete the last node)
	(&list).Delete(8);
	(&list).Delete(7);
	(&list).Print();

	// Test InsertBack
	(&list).InsertBack(&(Node{11, nil, nil}));
	(&list).Print();
	
	// Test Search
	fmt.Println("search 4", (&list).Search(4));
	fmt.Println("search 8", (&list).Search(8));
}