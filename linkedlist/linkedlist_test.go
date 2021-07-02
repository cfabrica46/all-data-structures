package linkedlist

import (
	"fmt"
	"sync"
	"testing"
)

func TestSimpleLinkedList(t *testing.T) {
	w := &sync.WaitGroup{}
	fmt.Println("~~~Simple Linked List~~~ ")
	simpleList := NewLinkedList()
	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			simpleList.InsertToEnd(i)
		}(i)
	}
	w.Wait()
	fmt.Println(simpleList)
	simpleList.Delete(9)
	fmt.Println(simpleList)
	simpleList.InsertToStart(999)
	fmt.Println(simpleList)
	simpleList.InsertToEnd(123)
	fmt.Println(simpleList)
	simpleList.InsertAfterTo(3, 32)
	fmt.Println(simpleList)
}

func TestDoubleLinkedList(t *testing.T) {
	w := &sync.WaitGroup{}
	fmt.Println("~~~Double Linked List~~~")

	doubleList := NewDoubleLinkedList()

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			doubleList.InsertToEnd(i)
		}(i)
	}

	w.Wait()
	fmt.Println(doubleList)
	doubleList.InsertToStart(999)
	fmt.Println(doubleList)
	doubleList.Delete(3)
	fmt.Println(doubleList)
	doubleList.InsertAfterTo(5, 789)
	fmt.Println(doubleList)
}
