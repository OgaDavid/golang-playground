package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type SingleLinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (list *SingleLinkedList) insertFirst(val int) {
	node := &Node{Data: val}

	node.Next = list.Head
	list.Head = node

	if list.Tail == nil {
		list.Tail = node
	}

	list.Size++
}

func (list *SingleLinkedList) deleteFirst() int {
	val := list.Head.Data
	list.Head = list.Head.Next

	// if head becomes nil after calling next, update tail to nil too
	if list.Head == nil {
		list.Tail = nil
	}

	list.Size--
	return val
}

func (list *SingleLinkedList) insertLast(val int) {
	node := &Node{Data: val}

	if list.Tail == nil {
		list.insertFirst(val)
		return
	}

	list.Tail.Next = node
	list.Tail = node

	list.Size++
}

func (list *SingleLinkedList) insert(val, pos int) {
	node := &Node{Data: val}

	if pos < 0 || pos > list.Size {
		panic("position out of range")
	}

	if pos == 0 {
		list.insertFirst(val)
		return
	}

	if pos == list.Size {
		list.insertLast(val)
		return
	}

	tempNode := list.Head

	for i := 0; i < pos-1; i++ {
		tempNode = tempNode.Next
	}

	holdNext := tempNode.Next
	tempNode.Next = node
	node.Next = holdNext
	list.Size++
}

func (list *SingleLinkedList) displayList() {
	for list.Head != nil {
		fmt.Printf("%d -> ", list.Head.Data)
		list.Head = list.Head.Next
	}

	fmt.Println("END")
}

func main() {
	var numbersList SingleLinkedList

	numbersList.insertFirst(5)
	numbersList.insertFirst(1)
	numbersList.insertFirst(3)
	numbersList.insertFirst(8)
	numbersList.insertLast(15)
	numbersList.insert(7, 2)

	numbersList.displayList()
}
