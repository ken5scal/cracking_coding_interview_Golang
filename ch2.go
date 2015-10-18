package main

import (
	"fmt"
	"strings"

	list "github.com/ken5scal/linked_list"
)

func doublyLinkedListTest() {
	dashes := strings.Repeat("-", 50)

	mfList := new(list.DoublyLinkedList)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})

	processed := make(map[*list.Node]bool)

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

func singlyLinkedListTest() {
	dashes := strings.Repeat("-", 50)

	mfList := new(list.SinglyLinkedList)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})

	processed := make(map[*list.Node]bool)

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
	name_to_find := "Kengo"
	mfList.Find(name_to_find)
	// Not in the List
	name_to_find = "Tsuji"
	mfList.Find(name_to_find)

	fmt.Println(dashes)

	name_to_remove := "Ichikawa"
	mfList.Delete(name_to_remove)
	mfList.Delete(name_to_remove)

	fmt.Println(dashes)

	fmt.Println("*Pop each value in the list...")
	//for person, err := mfList.Pop(); err == nil; person, err = mfList.Pop() {
	//	fmt.Printf("%v\n", person)
	//}
	fmt.Println(mfList.Pop()) //
	fmt.Println(mfList.Pop()) //
	fmt.Println(mfList.Pop()) //
	fmt.Println(mfList.Pop()) //
}

func singlyLinkedListReverseTest() {
	mfList := new(list.SinglyLinkedList)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})
	mfList.Reverse()
	fmt.Println(mfList.First().Name)
	fmt.Println(mfList.First().Next().Name)
	fmt.Println(mfList.First().Next().Next().Name)
}

func main() {
	//doublyLinkedListTest()
	//singlyLinkedListTest()
	singlyLinkedListReverseTest()
}
