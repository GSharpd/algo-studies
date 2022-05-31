// stacks and queues usually only access the first or last element of the collection
// in a stack you have objects being passed on top of the stack, in such a manner that you can only "touch" the top element of the stack
// in a single call to the stack, you can only touch the top element
// in a queue, you can only "touch" the first element of the queue

// stacks usually have 3 methods: pop, push and peek, all of them with a bigO of O(1)
// stacks -> LIFO (last in first out)
// works well with arrays and linked lists

// queues -> FIFO (first in first out)
// methods: enqueue, dequeue, peek
// queues works better with linked lists, since arrays need to be shifted when you remove the element at the first index and move every other element to the index of i-1
// with a linked list, all you need to do is move the head and the tail pointers, making it O(1), or a constant time operation

package main

import "fmt"

type Node struct {
	Value int
}

func (n *Node) String() string {
	return fmt.Sprint(n.Value)
}

// Returns the top(last) element of the stack
func (s *Stack) Peek() *Node {
	return s.nodes[len(s.nodes)-1]
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Stack is a basic LIFO stack that resizes as needed.
type Stack struct {
	nodes []*Node
	count int
}

// Push adds a node to the stack.
func (s *Stack) Push(n *Node) {
	s.nodes = append(s.nodes[:s.count], n)
	s.count++
}

// Pop removes and returns a node from the stack in last to first order.
func (s *Stack) Pop() *Node {
	if s.count == 0 {
		return nil
	}
	s.count--
	return s.nodes[s.count]
}

// NewQueue returns a new queue with the given initial size.
func NewQueue(size int) *Queue {
	return &Queue{
		nodes: make([]*Node, size),
		size:  size,
	}
}

// Queue is a basic FIFO queue based on a circular list that resizes as needed.
type Queue struct {
	nodes []*Node
	size  int
	head  int
	tail  int
	count int
}

// Push adds a node to the queue.
func (q *Queue) Push(n *Node) {
	if q.head == q.tail && q.count > 0 {
		nodes := make([]*Node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.head:])
		copy(nodes[len(q.nodes)-q.head:], q.nodes[:q.head])
		q.head = 0
		q.tail = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.tail] = n
	q.tail = (q.tail + 1) % len(q.nodes)
	q.count++
}

// Pop removes and returns a node from the queue in first to last order.
func (q *Queue) Pop() *Node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.head]
	q.head = (q.head + 1) % len(q.nodes)
	q.count--
	return node
}

func (q *Queue) Peek() *Node {
	return q.nodes[q.head]
}

func main() {
	s := NewStack()
	s.Push(&Node{3})
	s.Push(&Node{5})
	s.Push(&Node{7})
	s.Push(&Node{9})
	fmt.Println(s.Pop(), s.Pop(), s.Pop(), s.Pop())
	fmt.Println(s.Peek())

	q := NewQueue(1)
	q.Push(&Node{2})
	q.Push(&Node{4})
	q.Push(&Node{6})
	q.Push(&Node{8})
	fmt.Println(q.Pop(), q.Pop(), q.Pop(), q.Pop())
	fmt.Println(q.Peek())
}
