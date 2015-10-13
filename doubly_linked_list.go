package main

type Person struct {
	LastName  string
	FirstName string
}

type Node struct {
	Person
	next, prev *Node
}
