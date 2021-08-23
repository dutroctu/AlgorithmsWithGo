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

func main() {
	l := new(DoublyLinkedList)

	l.AddHead(17)
	l.AddHead(12)
	l.AddTail(1991)

	for n := l.First(); n != nil; n = n.Nexte() {
		fmt.Printf("%v\n", n.Data)
	}
}
