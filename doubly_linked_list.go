package main

import (
	"errors"
	"fmt"
	"strings"
)

type Person struct {
	Name string
	Age  int
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

// Append Node
func (l *List) Push(p Person) *List {
	n := &Node{Person: p}
	if l.head == nil { // First Node
		l.head = n
	} else {
		l.tail.next = n // link to next
		n.prev = l.tail // link to previous
	}
	l.tail = n // update tail to new node
	return l
}

// Find Node
func (l *List) Find(name string) *Node {
	found := false
	var result *Node = nil
	for n := l.First(); n != nil && !found; n = n.Next() {
		if n.Person.Name == name {
			found = true
			result = n
			break
		}
	}

	return result
}

// Delete Node
func (l *List) Delete(name string) bool {
	success := false
	node2del := l.Find(name)
	if node2del != nil {
		fmt.Println("Delete - Found: ", name)
		prev_node := node2del.prev
		next_node := node2del.next

		prev_node.next = node2del.next
		next_node.prev = node2del.prev
		success = true
	}
	return success
}

var errEmpty = errors.New("Error - List is empty")

// Pop last item from list
func (l *List) Pop() (p Person, err error) {
	if l.tail == nil {
		err = errEmpty
	} else {
		p = l.tail.Person
		l.tail = l.tail.prev
		if l.tail == nil {
			l.head = nil
		}
	}
	return p, err
}

func main() {
	dashes := strings.Repeat("-", 50)

	l := new(List)
	l.Push(Person{Name: "Kengo", Age: 27})
	l.Push(Person{Name: "Ichikawa", Age: 33})
	l.Push(Person{Name: "Asano", Age: 35})

	processed := make(map[*Node]bool)

	fmt.Println("First time through list...")
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf("%s\n", n.Name)
		if processed[n] {
			fmt.Printf("%s as been processed\n", n.Name)
		}
		processed[n] = true
	}

	fmt.Println(dashes)
	fmt.Println("Second time through list...")
	for n := l.First(); n != nil; n = n.Next() {
		fmt.Printf("%v", n.Name)
		if processed[n] {
			fmt.Println("has benn processed")
		} else {
			fmt.Println()
		}
		processed[n] = true
	}

	fmt.Println(dashes)
	var found_node *Node
	name_to_find := "Kengo"
	found_node = l.Find(name_to_find)
	if found_node == nil {
		fmt.Println("Not Found : %v\n", name_to_find)
	} else {
		fmt.Printf("Found: %v\n", name_to_find)
	}
	// Not in the List
	name_to_find = "Tsuji"
	found_node = l.Find(name_to_find)
	if found_node == nil {
		fmt.Println("Not Found : %v\n", name_to_find)
	} else {
		fmt.Printf("Found: %v\n", name_to_find)
	}

}
