package binarytree

import (
	"testing"
)

func TestBinaryTree(t *testing.T) {
	tree := &BinaryTree{}
	tree.Insert(100)
	tree.Insert(20)
	tree.Insert(10)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(200)
	tree.Insert(110)

	PrintTree(tree.head, 0, 'H')
}
