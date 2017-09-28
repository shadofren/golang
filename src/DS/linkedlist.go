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