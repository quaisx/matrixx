// -----------------------------------------------------------
// Technical assessment
// Implement a method that reverses the order of a singly
// linked list
// -----------------------------------------------------------
package main

import (
	"fmt"
	"strings"
)

// Structure that supports singly linked list
type Node[T any] struct {
	value T
	next  *Node[T]
}

// Reverse method reverses the order of a singly linked list
func (n *Node[T]) Reverse() *Node[T] {
	if n == nil {
		return nil
	}
	head := n
	tmp := n.next
	head.next = nil

	for tmp != nil {
		forward := tmp.next
		tmp.next = head
		head = tmp
		tmp = forward
	}
	return head
}

// Implement Stringer interface to print the contents of the singly linked list
func (n *Node[T]) String() string {
	if n == nil {
		return "nil"
	}
	p := n
	var buf []string
	for p != nil {
		buf = append(buf, fmt.Sprintf("%+v", p.value))
		p = p.next
	}
	return strings.Join(buf, ", ")
}
