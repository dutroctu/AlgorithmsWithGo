package main

import "fmt"

type DoublyLinkedListNode struct {
	Data int
	Next *DoublyLinkedListNode
	Prev *DoublyLinkedListNode
}

type DoublyLinkedList struct {
	Head *DoublyLinkedListNode
	Tail *DoublyLinkedListNode
}

func (l *DoublyLinkedList) First() *DoublyLinkedListNode {
	return l.Head
}

func (n *DoublyLinkedListNode) Nexte() *DoublyLinkedListNode {
	return n.Next
}

func (l *DoublyLinkedList) AddHead(data int) {
	newNode := &DoublyLinkedListNode{Data: data, Prev: nil}
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	newNode.Next = l.Head
	l.Head = newNode
	if l.Tail == nil {
		l.Tail = newNode
	}
}

func (l *DoublyLinkedList) AddTail(data int) {
	newNode := &DoublyLinkedListNode{Data: data, Next: nil}
	if l.Tail != nil {
		l.Tail.Next = newNode
	}
	l.Tail = newNode
	if l.Head == nil {
		l.Head = newNode
	}
}

func (l *DoublyLinkedList) Find(data int) *DoublyLinkedListNode {

	current := l.Head

	for current != nil {
		if current.Data == data {
			return current
		}
		current = current.Next
	}
	return nil
}

func (l *DoublyLinkedList) Remove(data int) bool {
	found := l.Find(data)
	if found == nil {
		return false
	}

	next := found.Next
	prev := found.Prev

	/* remove HEAD */
	if prev == nil {
		l.Head = next
		if l.Head != nil {
			l.Head.Prev = nil
		}
	} else {
		prev.Next = next

	}

	/* remove TAIL */
	if next == nil {
		l.Tail = prev
		if l.Tail != nil {
			l.Tail.Next = nil
		}
	} else {
		next.Prev = prev
	}

	return true
}

func (l *DoublyLinkedList) printf() {
	fmt.Println("Printing...")
	for n := l.First(); n != nil; n = n.Nexte() {
		fmt.Printf("%v\n", n.Data)
	}
}

func main() {
	l := new(DoublyLinkedList)

	l.AddHead(17)
	l.AddHead(12)
	l.AddTail(1991)

	l.printf()

	node := l.Find(1992)
	if node != nil {
		fmt.Printf("%v is exist\n", node.Data)
	} else {
		fmt.Printf("is not exist\n")
	}

	if ok := l.Remove(17); ok {
		l.printf()
		l.AddHead(17)
		l.printf()
	}

	head := l.Head.Data
	if ok := l.Remove(head); ok {
		l.printf()
		l.AddHead(head)
		l.printf()
	}

	tail := l.Tail.Data
	if ok := l.Remove(tail); ok {
		l.printf()
		l.AddTail(tail)
		l.printf()
	}

	if ok := l.Remove(1992); ok {
		l.printf()
	}
}
