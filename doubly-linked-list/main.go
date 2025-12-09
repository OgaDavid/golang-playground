package main

import "fmt"

type Node struct {
	Data int
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Head *Node
	Size int
}

func (list *DoublyLinkedList) insertFirst(val int) {
	node := &Node{Data: val}

	node.Next = list.Head
	if list.Head != nil {
		list.Head.Prev = node
	}
	list.Head = node
}

func (list *DoublyLinkedList) displayList() {
	temp := list.Head

	for temp != nil {
		fmt.Printf("%d -> ", temp.Data)
		temp = temp.Next
	}

	fmt.Println("END")
}

func main() {
	var numbersDoublyList DoublyLinkedList

	numbersDoublyList.insertFirst(2)
	numbersDoublyList.insertFirst(9)
	numbersDoublyList.insertFirst(7)
	numbersDoublyList.displayList()

}
