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
	Data interface{}
	Next *SimpleNode
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
		s += fmt.Sprintf("%v -> ", l.head.Data)
		l.head = l.head.Next
	}
	return s
}

func (l *SimpleLinkedList) InsertToStart(d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{Data: d}
	n.Next = l.head
	l.head = n
}

func (l *SimpleLinkedList) InsertToEnd(d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{Data: d}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.Next = n
		l.tail = n
		l.length++
	}
}

func (l *SimpleLinkedList) InsertAfterTo(key, d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()
	n := &SimpleNode{Data: d}
	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	curr := l.head
	for curr.Data != key {
		curr = curr.Next
	}
	if curr == nil {
		l.tail.Next = n
		l.tail = n
		return
	}
	n.Next = curr.Next
	curr.Next = n
}

func (l *SimpleLinkedList) GetHead() *SimpleNode {
	return l.head
}

func (l *SimpleLinkedList) GetTail() *SimpleNode {
	return l.tail
}

func (l SimpleLinkedList) Get(key interface{}) interface{} {

	for l.head != nil {
		if l.head.Data == key {
			return l.head
		}
		l.head = l.head.Next
	}
	return nil
}

func (l *SimpleLinkedList) Delete(key interface{}) (curr *SimpleNode) {
	l.m.Lock()
	defer l.m.Unlock()
	if l.head.Data == key {
		l.head = l.head.Next
		l.length--
		return
	}
	var prev *SimpleNode = nil
	curr = l.head
	for curr != nil && curr.Data != key {
		prev = curr
		curr = curr.Next
	}
	if curr == nil {
		return curr
	}
	if curr == l.tail {
		l.tail = prev
	}
	prev.Next = curr.Next
	l.length--

	return
}

//-------------------------------------------------------------------------------------------

type DoubleNode struct {
	prev, Next *DoubleNode
	Data       interface{}
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
		s += fmt.Sprintf("%v -> ", l.head.Data)
		l.head = l.head.Next
	}
	return s
}

func (l *DoubleLinkedList) InsertToStart(d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{Data: d}

	n.Next = l.head
	l.head.prev = n
	l.head = n
}

func (l *DoubleLinkedList) InsertToEnd(d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{Data: d}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
	} else {
		l.tail.Next = n
		n.prev = l.tail
		l.tail = n
		l.length++
	}
}

func (l *DoubleLinkedList) InsertAfterTo(key, d interface{}) {
	l.m.Lock()
	defer l.m.Unlock()

	n := &DoubleNode{Data: d}

	if l.head == nil {
		l.head = n
		l.tail = n
		l.length++
		return
	}
	curr := l.head
	for curr.Data != key {
		curr = curr.Next
	}
	if curr == nil {
		n.prev = l.tail
		l.tail.Next = n
		l.tail = n
		return
	}
	n.Next = curr.Next
	n.prev = curr
	if curr.Next != nil {
		curr.Next.prev = n
	}
	curr.Next = n
}

func (l *DoubleLinkedList) GetHead() *DoubleNode {
	return l.head
}

func (l *DoubleLinkedList) GetTail() *DoubleNode {
	return l.tail
}

func (l DoubleLinkedList) Get(key interface{}) interface{} {

	for l.head != nil {
		if l.head.Data == key {
			return l.head
		}
		l.head = l.head.Next
	}
	return nil
}

func (l *DoubleLinkedList) Delete(key int) (curr *DoubleNode) {
	l.m.Lock()
	defer l.m.Unlock()

	if l.head.Data == key {
		l.head = l.head.Next
		l.head.prev = nil
		return
	}
	curr = l.head
	for curr != nil && curr.Data != key {
		curr = curr.Next
	}
	if curr == nil {
		return
	}
	if curr == l.tail {
		curr.prev.Next = nil
		l.length--
		return
	}
	curr.prev.Next = curr.Next
	l.length--
	return
}
