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
			fmt.Printf("Found : %v\n", name)
			return result
		}
	}
	fmt.Printf("Not Found : %v\n", name)
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

		if prev_node != nil { // Not the First Node
			prev_node.next = next_node
		} else {
			l.head = next_node
		}
		if next_node != nil { // Not the Last Node
			next_node.prev = prev_node
		} else {
			l.tail = prev_node
		}
		success = true
	}
	if success {
		fmt.Printf("Removed: %v\n", name)
	} else {
		fmt.Printf("Failed Removing: %v\n", name)
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

	mfList := new(List)
	mfList.Push(Person{Name: "Kengo", Age: 27})
	mfList.Push(Person{Name: "Ichikawa", Age: 33})
	mfList.Push(Person{Name: "Asano", Age: 35})

	processed := make(map[*Node]bool)

	fmt.Println("First time through list...")
	for n := mfList.First(); n != nil; n = n.Next() {
		fmt.Printf("%s\n", n.Name)
		if processed[n] {
			fmt.Printf("%s as been processed\n", n.Name)
		}
		processed[n] = true
	}

	fmt.Println(dashes)
	fmt.Println("Second time through list...")
	for n := mfList.First(); n != nil; n = n.Next() {
		fmt.Printf("%v", n.Name)
		if processed[n] {
			fmt.Println(" has benn processed")
		} else {
			fmt.Println()
		}
		processed[n] = true
	}

	fmt.Println(dashes)
	// var found_node *Node
	name_to_find := "Kengo"
	mfList.Find(name_to_find)
	// Not in the List
	name_to_find = "Tsuji"
	mfList.Find(name_to_find)

	fmt.Println(dashes)
	name_to_remove := "Kengo"
	mfList.Delete(name_to_remove)
	mfList.Delete(name_to_remove)

	fmt.Println(dashes)
	fmt.Println("*Pop each value in the list...")
	for person, err := mfList.Pop(); err == nil; person, err = mfList.Pop() {
		fmt.Printf("%v\n", person)
	}
	//fmt.Println(mfList.Pop())
}
