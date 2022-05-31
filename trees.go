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
}
