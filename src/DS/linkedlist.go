package DS

import "fmt"

type Node struct {
	Val int;
	Prev *Node;
	Next *Node;
}

func (sentinel *Node) InsertFront(node *Node) {
	node.Prev = sentinel;
	node.Next = sentinel.Next;
	sentinel.Next.Prev = node;
	sentinel.Next = node;
}

func (sentinel *Node) InsertBack(node *Node) {
	node.Next = sentinel;
	node.Prev = sentinel.Prev;
	sentinel.Prev.Next = node;
	sentinel.Prev = node;
}

func (sentinel *Node) Search(val int) *Node {
	node := sentinel.Next;
	for node != sentinel && node.Val != val {
		node = node.Next;
	}
	return node;
}

func (sentinel *Node) Delete(node *Node) {
	node.Next.Prev = node.Prev;
	node.Prev.Next = node.Next;
}

func (sentinel *Node) Print() {
	node := sentinel.Next;
	for node != sentinel {
		fmt.Print(node.Val, " ");
		node = node.Next;
	}
	fmt.Println();
}

func TestLinkedList() {
	sentinel := Node{-1, nil, nil};
	sentinel.Next = &sentinel;
	sentinel.Prev = &sentinel;

	// Test InsertFront
	var node *Node;
	for i := 1; i <= 10; i++ {
		newNode := Node{i, nil, nil};
		node = &newNode;
		(&sentinel).InsertFront(&newNode);
	}
	(&sentinel).Print();

	// Test Delete (Delete the last node)
	(&sentinel).Delete(node.Next.Next);
	(&sentinel).Delete(node);
	(&sentinel).Print();

	// Test InsertBack
	(&sentinel).InsertBack(&(Node{11, nil, nil}));
	(&sentinel).Print();
	
	// Test Search
	fmt.Println("search 4", (&sentinel).Search(4));
	fmt.Println("search 8", (&sentinel).Search(8));
}