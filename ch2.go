package main

import (
	"fmt"
	"strings"

	list "github.com/ken5scal/linked_list"
)

func problem1() {
	fmt.Printf("Problem 2.1\n")
	dashes := strings.Repeat("-", 50)
	fmt.Println(dashes)
	mfList := new(list.DoublyLinkedList)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	fmt.Println(dashes)
}

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

func singlyLinkedListRemoveDuplicatesTest() {
	mfList := new(list.SinglyLinkedList)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.RemoveDuplicates()

	for n := mfList.First(); n != nil; n = n.Next() {
		fmt.Printf("%s\n", n.Name)
	}
}

var Kengo = &list.Person{Name: "Kengo", Age: 27}
var Ichikawa = &list.Person{Name: "Ichikawa", Age: 33}
var Asano = &list.Person{Name: "Asano", Age: 35}
var Tsuji = &list.Person{Name: "Tsuji", Age: 40}
var Tsuzuki = &list.Person{Name: "Tsuzuki", Age: 42}

func SinglyLinkedListFindInReverseOrder(idx int) {
	mfList := new(list.SinglyLinkedList)
	//	mfList.Push(Kengo)
	//	mfList.Push(Ichikawa)
	//	mfList.Push(Asano)
	//	mfList.Push(Tsuji)
	//	mfLIst.Push(Tsuzuki)
	mfList.Push(list.Person{Name: "Kengo", Age: 27})
	mfList.Push(list.Person{Name: "Ichikawa", Age: 33})
	mfList.Push(list.Person{Name: "Asano", Age: 35})
	mfList.Push(list.Person{Name: "Tsuji", Age: 40})
	mfList.Push(list.Person{Name: "Tsuzuki", Age: 42})
	fmt.Println(mfList.FindInReverseOrder(idx).Name)
}

func main() {
	//doublyLinkedListTest()
	//singlyLinkedListTest()
	dashes := strings.Repeat("-", 20)
	fmt.Println(dashes + " Reversing Linked List " + dashes)
	singlyLinkedListReverseTest()
	fmt.Println(dashes + " Remove Duplicates in Linked List " + dashes)
	singlyLinkedListRemoveDuplicatesTest()
	fmt.Println(dashes + " Find specified index in reverse order " + dashes)
	SinglyLinkedListFindInReverseOrder(0)
}
