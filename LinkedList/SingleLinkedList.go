package LinkedList

import (
	"errors"
	"fmt"
	"strconv"
)

type Node struct {
	Value      int
	Next, Prev *Node
}

type SingleLinkedList struct {
	Size int
	Head *Node
}

func Constractor() *SingleLinkedList {
	return new(SingleLinkedList)
}

func (l *SingleLinkedList) Append(i int) {

	item := &Node{
		Value: i,
	}

	if l.Head == nil {
		l.Head = item
	} else {
		last := l.Head
		for {
			if last.Next == nil {
				break
			}

			last = last.Next
		}
		last.Next = item
	}

	l.Size++
}

func (l *SingleLinkedList) Insert(i int, n int) error {
	if i < 0 || i > l.Size {
		return fmt.Errorf("index out of bound")
	}

	newNode := &Node{
		Value: n,
	}
	if i == 0 {
		newNode.Next = l.Head
		l.Head = newNode
		l.Size++
		return nil
	}

	j := 0
	node := l.Head
	for j < i-2 {
		j++
		node = node.Next
	}

	newNode.Next = node.Next
	node.Next = newNode

	l.Size++
	return nil
}

func (l *SingleLinkedList) RemoveAt(i int) (n *Node, err error) {
	if i > l.Size || i < 0 {
		return nil, fmt.Errorf("out of bound")
	}

	node := l.Head
	j := 0

	for j < i-1 {
		j++
		node = node.Next
	}

	remove := node.Next
	node.Next = remove.Next
	l.Size--

	return remove, nil
}

func (l *SingleLinkedList) Remove() {
	node := l.Head
	var prev *Node

	if l.Size == 1 {
		l.Head = nil
		l.Size--
		return
	}

	for {
		if node == nil {
			break
		}

		if node.Next == nil {
			prev.Next = nil
			break
		}

		prev = node
		node = node.Next
	}

	l.Size--
}

func (l *SingleLinkedList) IndexOf(value int) int {
	node := l.Head

	j := 0
	for {

		if node == nil {
			return -1
		}

		if node.Value == value {
			return j
		}

		node = node.Next
		j++
	}
}
func (l *SingleLinkedList) IsEmpty() bool {
	return l.Head == nil
}

// Reverse represents making reverse single list
// 1. step
// current = 1
// prev = nil
// next = 2
//
// 2. step
// current = 2
// prev = 1
// next = 3
//
// 3. step
// current = 3
// prev = 2
// next = 4
//
// 4. step
// current = 4
// prev = 3
// next = 4
//
// Time Complexity = O(n) linear
// Space Complexity = O(1) constant
func (l *SingleLinkedList) Reverse() {
	current := l.Head
	var prev *Node
	var next *Node

	for current != nil {
		next = current.Next

		current.Next = prev

		prev = current
		current = next
	}
	l.Head = prev
}

// isSortedAsc
func isSortedAsc(node *Node) bool {

	if node.Next == nil {
		return true
	}
	return node.Value < node.Next.Value && isSortedAsc(node.Next)
}

// MergeSorted
// fnc is a sequence checker . Can be asc or desc or something another.
// Example: MergeSorted(l.Head, l2.Head, isSortedAsc)
// TODO : in progress
func MergeSorted(first, second *Node, fnc func(node *Node) bool) (dummy *Node, err error) {

	dummy = new(Node)

	if fnc(second) && fnc(first) {

		for second != nil && first != nil {
			if first == nil {
				dummy.Next = second
				break
			}

			if second == nil {
				dummy.Next = first
				break
			}

			if first.Value <= second.Value {
				//MoveNode(dummy, first)
				dummy.Next = first

				first = first.Next

			} else {
				//MoveNode(dummy, second)
				dummy.Next = second
				second = second.Next

			}
		}

		return dummy, nil
	}

	return nil, errors.New("some of from there lists is unsorted")
}

func MoveNode(destRef, sourceRef *Node) {
	newNode := *sourceRef

	sourceRef = newNode.Next

	newNode.Next = destRef

	destRef = &newNode
}

func (l SingleLinkedList) String() (str string) {
	node := l.Head
	str = DisplayRecursion(node)
	return
}

func DisplayRecursion(n *Node) (str string) {
	if n == nil {
		return ""
	}

	if n.Next != nil {
		str += strconv.Itoa(n.Value) + "-->"
		str += DisplayRecursion(n.Next)
	} else {
		str += strconv.Itoa(n.Value)
	}

	return
}
