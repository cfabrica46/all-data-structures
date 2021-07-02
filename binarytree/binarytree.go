package binarytree

import "fmt"

type binaryNode struct {
	data        int
	left, right *binaryNode
}

type BinaryTree struct {
	head *binaryNode
}

func (t BinaryTree) IsEmpty() bool {
	return t.head == nil
}

func (t *BinaryTree) Insert(data int) {
	if t.IsEmpty() {
		t.head = &binaryNode{data: data, left: nil, right: nil}
	} else {
		t.head.insert(data)
	}
}

func (n *binaryNode) insert(data int) {
	if n == nil {
		return
	}

	if data <= n.data {
		if n.left == nil {
			n.left = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.left.insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &binaryNode{data: data, left: nil, right: nil}
		} else {
			n.right.insert(data)
		}
	}
}

func PrintTree(n *binaryNode, space int, head rune) {
	if n == nil {
		return
	}
	for i := 0; i < space; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%c:%v\n", head, n.data)
	PrintTree(n.left, space+2, 'L')
	PrintTree(n.right, space+2, 'R')
}
