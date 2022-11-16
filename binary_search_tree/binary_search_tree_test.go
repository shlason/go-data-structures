package binarySearchTree

import (
	"fmt"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {
	bst := New()

	t.Error("testing")

	bst.Insert(11)
	bst.Insert(7)
	bst.Insert(15)
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(9)
	bst.Insert(8)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(13)
	bst.Insert(12)
	bst.Insert(14)
	bst.Insert(20)
	bst.Insert(18)
	bst.Insert(25)

	bst.PreOrderTraverse()
	fmt.Println("-------------------------------------")
	bst.Remove(15)
	bst.Remove(13)
	bst.PreOrderTraverse()
}
