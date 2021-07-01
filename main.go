package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/cfabrica46/all-data-structures/linkedlist"
	"github.com/cfabrica46/all-data-structures/queue"
	"github.com/cfabrica46/all-data-structures/stack"
)

func main() {

	w := &sync.WaitGroup{}

	fmt.Println("~~~~~All Data Structures~~~~~")
	fmt.Println()
	fmt.Println("~~~Simple Linked List~~~ ")
	simpleList := linkedlist.NewLinkedList()
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

	fmt.Println()
	fmt.Println("~~~Double Linked List~~~")

	doubleList := linkedlist.NewDoubleLinkedList()

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

	fmt.Println()
	fmt.Println("~~~Stack~~~")
	stack := stack.NewStack()

	for i := 0; i < 10; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			stack.Push(i)
		}(i)

	}
	w.Wait()
	fmt.Println(stack)

	fmt.Println("Pop:")
	for i := 0; i < stack.Len(); i++ {
		if stack.Pop() == nil {
			fmt.Println("The Stack Is Empty")
		} else {
			fmt.Println(stack)
		}
	}

	fmt.Println()
	fmt.Println("~~~Queue~~~")

	JQ := queue.NewJobQueue()
	for i := 0; i < 6; i++ {
		err := JQ.JoinIntoQueue(queue.NewJob(time.Second))
		if err != nil {
			fmt.Println("The queue is full")
		}
	}
	fmt.Println(JQ)
	fmt.Scanln()
}
