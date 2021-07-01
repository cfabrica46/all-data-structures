package linkedlist

import (
	"fmt"
	"sync"
)

//Como se podria aprovechar una interface?
//type LinkedList interface {
//	Len() int
//	String() string
//	InsertToStart(int)
//	InsertToEnd(int)
//	InsertAfterTo(int, int)
//	Delete(int)
//}

type SimpleNode struct {
	data int
	next *SimpleNode
}

type SimpleLinkedList struct {
	m          *sync.Mutex
	length     int
	head, tail *SimpleNode
}

func NewLinkedList() *SimpleLinkedList {
	return &SimpleLinkedList{m: &sync.Mutex{}}
}

func (l SimpleLinkedList) Len() int {
	return l.length
}

func (l SimpleLinkedList) String() string {
	var s string
	for l.head != nil {
		s += fmt.Sprintf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	return s
}

func (l *SimpleLinkedList) InsertToStart(d int) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{data: d}
	n.next = l.head
	l.head = n
}

func (l *SimpleLinkedList) InsertToEnd(d int) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{data: d}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		l.tail = n
		l.length++
	}
}

func (l *SimpleLinkedList) InsertAfterTo(key, d int) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{data: d}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	curr := l.head
	for curr.data != key {
		curr = curr.next
	}
	if curr == nil {
		l.tail.next = n
		l.tail = n
		return
	}
	n.next = curr.next
	curr.next = n
}

func (l *SimpleLinkedList) Delete(key int) (curr *SimpleNode) {
	l.m.Lock()
	defer l.m.Unlock()
	if l.head.data == key {
		l.head = l.head.next
		l.length--
		return
	}
	var prev *SimpleNode = nil
	curr = l.head
	for curr != nil && curr.data != key {
		prev = curr
		curr = curr.next
	}
	if curr == nil {
		return curr
	}
	if curr == l.tail {
		l.tail = prev
	}
	prev.next = curr.next
	l.length--
	return
}

//-------------------------------------------------------------------------------------------

type DoubleNode struct {
	prev, next *DoubleNode
	data       int
}

type DoubleLinkedList struct {
	m          *sync.Mutex
	length     int
	head, tail *DoubleNode
}

func NewDoubleLinkedList() *DoubleLinkedList {
	return &DoubleLinkedList{m: &sync.Mutex{}}
}

func (l DoubleLinkedList) Len() int {
	return l.length
}

func (l DoubleLinkedList) String() string {
	var s string

	for l.head != nil {
		s += fmt.Sprintf("%v -> ", l.head.data)
		l.head = l.head.next
	}
	return s
}

func (l *DoubleLinkedList) InsertToStart(d int) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{data: d}

	n.next = l.head
	l.head.prev = n
	l.head = n
}

func (l *DoubleLinkedList) InsertToEnd(d int) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{data: d}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.next = n
		n.prev = l.tail
		l.tail = n
		l.length++
	}
}

func (l *DoubleLinkedList) InsertAfterTo(key int, d int) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{data: d}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	curr := l.head
	for curr.data != key {
		curr = curr.next
	}
	if curr == nil {
		n.prev = l.tail
		l.tail.next = n
		l.tail = n
		return
	}
	n.next = curr.next
	n.prev = curr
	if curr.next != nil {
		curr.next.prev = n
	}
	curr.next = n
}

func (l *DoubleLinkedList) Delete(key int) (curr *DoubleNode) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.head.data == key {
		l.head = l.head.next
		l.head.prev = nil
		return
	}
	curr = l.head
	for curr != nil && curr.data != key {
		curr = curr.next
	}
	if curr == nil {
		return
	}
	if curr == l.tail {
		curr.prev.next = nil
		l.length--
		return
	}
	curr.prev.next = curr.next
	l.length--
	return
}
