package binarySearchTree

import "fmt"

type key byte

type node struct {
	key
	left  *node
	right *node
}

func (n *node) insertNode(k key) {
	if k >= n.key {
		if n.right == nil {
			n.right = &node{
				key: k,
			}
		} else {
			n.right.insertNode(k)
		}
	} else {
		if n.left == nil {
			n.left = &node{
				key: k,
			}
		} else {
			n.left.insertNode(k)
		}
	}
}

func (n *node) inOrderTraverseNode() {
	// if n.left == nil {
	// 	fmt.Println(n.key)
	// 	return
	// } else {
	// 	n.left.inOrderTraverseNode()
	// 	fmt.Println(n.key)
	// }
	// if n.right != nil {
	// 	n.right.inOrderTraverseNode()
	// }
	if n == nil {
		return
	}
	n.left.inOrderTraverseNode()
	fmt.Println(n.key)
	n.right.inOrderTraverseNode()
}

func (n *node) preOrderTraverseNode() {
	// fmt.Println(n.key)
	// if n.left != nil {
	// 	n.left.preOrderTraverseNode()
	// }
	// if n.right != nil {
	// 	n.right.preOrderTraverseNode()
	// }
	if n == nil {
		return
	}
	fmt.Println(n.key)
	n.left.preOrderTraverseNode()
	n.right.preOrderTraverseNode()
}

func (n *node) postOrderTraverseNode() {
	// if n.left == nil {
	// 	fmt.Println(n.key)
	// } else {
	// 	n.left.postOrderTraverseNode()
	// }
	// if n.right != nil {
	// 	n.right.postOrderTraverseNode()
	// 	fmt.Println(n.key)
	// }
	if n == nil {
		return
	}
	n.left.postOrderTraverseNode()
	n.right.postOrderTraverseNode()
	fmt.Println(n.key)
}

type binarySearchTree struct {
	root *node
}

func New() *binarySearchTree {
	return new(binarySearchTree)
}

func (bst *binarySearchTree) Insert(k key) {
	if bst.root == nil {
		bst.root = &node{
			key: k,
		}
		return
	}

	bst.root.insertNode(k)
}

func (bst *binarySearchTree) InOrderTraverse() {
	if bst.root == nil {
		fmt.Println("BST empty")
		return
	}
	bst.root.inOrderTraverseNode()
}

func (bst *binarySearchTree) PreOrderTraverse() {
	if bst.root == nil {
		fmt.Println("BST empty")
		return
	}
	bst.root.preOrderTraverseNode()
}

func (bst *binarySearchTree) PostOrderTraverse() {
	if bst.root == nil {
		fmt.Println("BST empty")
		return
	}
	bst.root.postOrderTraverseNode()
}

func (bst *binarySearchTree) Min() key {
	current := bst.root
	if current == nil {
		return 0
	}
	for current.left != nil {
		current = current.left
	}
	return current.key
}

func (bst *binarySearchTree) Max() key {
	current := bst.root
	if current == nil {
		return 0
	}
	for current.right != nil {
		current = current.right
	}
	return current.key
}

func (bst *binarySearchTree) Search(k key) bool {
	current := bst.root
	if current == nil {
		return false
	}
	for current != nil {
		if current.key > k {
			current = current.left
		} else if current.key < k {
			current = current.right
		} else {
			return true
		}
	}
	return false
}
