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
	if list.Head == nil {
		panic("cannot delete from an empty list")
	}

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

func (list *SingleLinkedList) deleteLast() int {
	if list.Tail == nil {
		panic("cannot delete from an empty list")
	}

	if list.Size == 1 {
		val := list.Head.Data
		list.deleteFirst()
		return val
	}

	val := list.Tail.Data
	tempNode := list.Head

	for i := 0; i < list.Size-2; i++ {
		tempNode = tempNode.Next
	}

	tempNode.Next = nil
	list.Tail = tempNode
	list.Size--

	return val
}

func (list *SingleLinkedList) insert(val, pos int) {
	node := &Node{Data: val}

	if pos < 0 || pos >= list.Size {
		panic("position out of range")
	}

	if pos == 0 {
		list.insertFirst(val)
		return
	}

	if pos == list.Size-1 {
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

func (list *SingleLinkedList) delete(pos int) int {
	if pos < 0 || pos >= list.Size {
		panic("position out of range")
	}

	if pos == 0 {
		val := list.Head.Data
		list.deleteFirst()
		return val
	}

	if pos == list.Size-1 {
		val := list.Tail.Data
		list.deleteLast()
		return val
	}

	prevNode := list.Head

	for i := 0; i < pos-1; i++ {
		prevNode = prevNode.Next
	}

	toDelete := prevNode.Next
	val := toDelete.Data

	prevNode.Next = toDelete.Next
	list.Size--

	return val
}

func (list *SingleLinkedList) findByValue(val int) *Node {
	node := list.Head

	for node != nil {
		if node.Data == val {
			return node
		}
		node = node.Next
	}

	return nil
}

func (list *SingleLinkedList) displayList() {
	temp := list.Head
	for temp != nil {
		fmt.Printf("%d -> ", temp.Data)
		temp = temp.Next
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
	fmt.Printf("Delete First: %d\n", numbersList.deleteFirst())
	fmt.Printf("Delete Last: %d\n", numbersList.deleteLast())
	numbersList.displayList()
	fmt.Printf("Delete at Index 2: %d\n", numbersList.delete(2))
	numbersList.displayList()
	fmt.Printf("Node with data 2: %v\n", numbersList.findByValue(2))

}
