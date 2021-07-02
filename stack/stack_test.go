package stack

import (
	"fmt"
	"sync"
	"testing"
)

func TestStack(t *testing.T) {
	w := &sync.WaitGroup{}
	stack := NewStack()

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

}
