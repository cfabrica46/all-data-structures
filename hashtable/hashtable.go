package hashtable

import (
	"fmt"
	"math"

	"github.com/cfabrica46/all-data-structures/linkedlist"
)

const arrayLength = 100

type HashTable struct {
	data [arrayLength]*linkedlist.SimpleLinkedList
}

type listData struct {
	key   string
	value interface{}
}

func (h *HashTable) String() string {
	var s string

	for i := range h.data {
		if h.data[i] != nil {
			s += fmt.Sprintf("%v", h.data[i])
		}
	}

	return s
}

func NewHashTable() *HashTable {
	return &HashTable{
		data: [arrayLength]*linkedlist.SimpleLinkedList{},
	}
}

func hashFunction(s string) int {
	h := 0
	for pos, char := range s {
		h += int(char) * int(math.Pow(31, float64(len(s)-pos+1)))
	}
	return h
}

func getIndex(k int) int {
	return k % arrayLength
}

func (h *HashTable) Set(k string, v interface{}) *HashTable {

	index := getIndex(hashFunction(k))

	if h.data[index] == nil {
		h.data[index] = linkedlist.NewLinkedList()
		h.data[index].InsertToStart(listData{k, v})
	} else {
		n := h.data[index].GetHead()
		for {
			if n == nil {
				h.data[index].InsertToStart(listData{k, v})
				break
			}

			d := n.Data.(listData)
			if d.key == k {
				d.value = v
				break
			}
			n = n.Next
		}
	}

	return h
}

func (h *HashTable) Get(k string) interface{} {
	index := getIndex(hashFunction(k))
	l := h.data[index]

	if l == nil {
		return nil
	}

	n := l.GetHead()

	for {
		if n == nil {
			return nil
		}

		d := n.Data.(listData)
		if d.key == k {
			return d.value
		}
		n = n.Next
	}

}

func (h *HashTable) Delete(k string) interface{} {
	index := getIndex(hashFunction(k))
	l := h.data[index]

	if l == nil {
		return nil
	}

	n := l.GetHead()

	for {
		if n == nil {
			return nil
		}

		d := n.Data.(listData)
		if d.key == k {
			l.Delete(d)
			return d.value
		}
		n = n.Next
	}

}
