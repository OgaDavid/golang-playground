package main

type Node struct {
	Data int
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
	Size int
}

func (list *LinkedList) insertFirst(val int) {
	node := &Node{Data: val}

	node.Next = list.Head
	list.Head = node

	if list.Tail == nil {
		list.Tail = node
	}

	list.Size++
}

func (list *LinkedList) insertLast(val int) {
	node := &Node{Data: val}

	if list.Tail == nil {
		list.insertFirst(val)
		return
	}

	list.Tail.Next = node
	list.Tail = node

	list.Size++
}

func (list *LinkedList) insert(val, pos int) {
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

func main() {

}
