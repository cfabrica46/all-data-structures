package stack

import (
	"fmt"
	"sync"
)

type Node struct {
	prev *Node
	data int
}

type Stack struct {
	m      *sync.Mutex
	length int
	head   *Node
}

func NewStack() *Stack {
	return &Stack{m: &sync.Mutex{}}
}

func (s Stack) Len() int {
	return s.length
}

func (s Stack) IsEmpty() (empty bool) {
	if s.head == nil {
		empty = true
	}
	return
}

func (s Stack) String() string {
	var b string
	for s.head != nil {
		b += fmt.Sprintf("<- %v  ", s.head.data)
		s.head = s.head.prev
	}
	return b
}

func (s *Stack) Push(d int) {
	s.m.Lock()
	defer s.m.Unlock()
	n := &Node{data: d}
	if s.IsEmpty() {
		s.head = n
		s.length++
	} else {
		n.prev = s.head
		s.head = n
		s.length++
	}
}

func (s *Stack) Pop() *Node {
	s.m.Lock()
	defer s.m.Unlock()

	if s.IsEmpty() {
		return nil
	}
	NodeDeleted := s.head
	s.head = s.head.prev
	return NodeDeleted
}
