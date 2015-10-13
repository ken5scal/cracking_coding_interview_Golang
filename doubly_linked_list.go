package main

type Person struct {
	LastName  string
	FirstName string
}

type Node struct {
	Person
	next, prev *Node
}

type List struct {
	head, tail *Node
}

func (l *List) First() *Node {
	return l.head
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}
