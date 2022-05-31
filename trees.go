// A linked list is essentially a unidirectional tree
// the difference is the way trees usually branch out
// e.g. a binary tree can have 2, 1 or 0 children, a linked list has 1 child
// in a "perfect binary tree", each node has 2 children, that way, the bottom of the tree always has twice as many nodes as the previous level and the sum of the rest of the nodes + 1 node
// this means more than around half of the data is on the bottom level of the tree

package main

import "fmt"

type TreeNode struct {
	left  *TreeNode
	right *TreeNode
	data  int32
}

type Tree struct {
	root *TreeNode
}

func (t *Tree) Insert(data int32) *Tree {
	if t.root == nil {
		t.root = &TreeNode{data: data}
	} else {
		t.root.Insert(data)
	}

	return t
}

func (n *TreeNode) Insert(data int32) {
	if n == nil {
		return
	} else if data <= n.data {
		if n.left == nil {
			n.left = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.left.Insert(data)
		}
	} else {
		if n.right == nil {
			n.right = &TreeNode{data: data, left: nil, right: nil}
		} else {
			n.right.Insert(data)
		}
	}
}

func (t *Tree) Lookup(value int32) *TreeNode {
	if t.root == nil {
		return nil
	}
	currentNode := t.root
	for currentNode != nil {
		if value < currentNode.data {
			currentNode = currentNode.left
		} else if value > currentNode.data {
			currentNode = currentNode.right
		} else if value == currentNode.data {
			return currentNode
		}
	}

	return currentNode
}

// I hate this implementation and want to kill myself for it
func (t *Tree) RemoveNode(value int32) bool {
	if t.root == nil {
		return false
	}

	currentNode := t.root
	parentNode := t.root

	for currentNode != nil {
		if value < currentNode.data {
			currentNode = currentNode.left
		} else if value > currentNode.data {
			currentNode = currentNode.right
		} else if value == currentNode.data {
			if currentNode.right == nil {
				if parentNode == nil {
					t.root = currentNode.left
				} else {
					if currentNode.data < parentNode.data {
						parentNode.left = currentNode.left
					} else {
						parentNode.right = currentNode.left
					}
				}
			} else if currentNode.right.left == nil {
				currentNode.right.left = currentNode.left
				if parentNode == nil {
					t.root = currentNode.right
				} else {
					if currentNode.data < parentNode.data {
						parentNode.left = currentNode.right
					} else if currentNode.data > parentNode.data {
						parentNode.right = currentNode.right
					}
				}
			} else {
				leftmost := currentNode.right.left
				leftmostParent := currentNode.right
				for leftmost.left != nil {
					leftmostParent = leftmost
					leftmost = leftmost.left
				}

				leftmostParent.left = leftmost.right
				leftmost.left = currentNode.left
				leftmost.right = currentNode.right

				if parentNode == nil {
					t.root = leftmost
				} else {
					if currentNode.data < parentNode.data {
						parentNode.left = leftmost
					} else if currentNode.data > parentNode.data {
						parentNode.right = leftmost
					}
				}
			}

			return true
		}
	}

	return false
}

func main() {
	tree := &Tree{}

	tree.Insert(9)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(20)
	tree.Insert(170)
	tree.Insert(15)
	tree.Insert(1)
	fmt.Println(tree.Lookup(15))
	fmt.Println(tree.Lookup(4))

	nodeRemoved := tree.RemoveNode(15)

	fmt.Println(nodeRemoved)
	fmt.Println(tree.RemoveNode(22))
}
