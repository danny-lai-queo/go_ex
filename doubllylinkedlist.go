package main

import (
	"errors"
	"fmt"
)

type Node[T any] struct {
	value      T
	next, prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

// 0 <= index <= q.size
func (q *DoublyLinkedList[T]) Add_Dan(index int, v T) error {
	if index < 0 || index > q.size {
		return errors.New("invalid index. index must satisfy  0 <= index < list.size")
	}

	vnode := Node[T]{
		value: v,
		next:  nil,
		prev:  nil,
	}

	// add to empty list
	if q.size == 0 {
		q.head = &vnode
		q.tail = &vnode
		q.size = 1
		return nil
	}

	// add before head
	if index == 0 {
		vnode.next = q.head
		q.head.prev = &vnode
		q.head = &vnode
		q.size += 1
		return nil
	}

	// add after head
	if index == q.size {
		q.tail.next = &vnode
		vnode.prev = q.tail
		q.tail = &vnode
		q.size += 1
		return nil
	}

	// 0 < index < q.size
	if 0 < index && index < q.size {
		curr := q.head
		for i := 0; i < q.size && i < index && curr != nil; i++ {
			curr = curr.next
		}
		parent := curr.prev

		parent.next = &vnode
		vnode.prev = parent
		vnode.next = curr
		curr.prev = &vnode

		q.size += 1
		//return nil
	}

	return nil
}

// 0 <= index <= q.size
func (q *DoublyLinkedList[T]) Add(index int, v T) error {
	if index > q.size {
		return errors.New("invalid index. index must satisfy  index < list.size")
	}

	currSize := q.size

	vnode := &Node[T]{
		value: v,
	}

	q.size += 1

	// add to empty list
	if q.head == nil {
		q.head, q.tail = vnode, vnode
		return nil
	}

	// add before head
	if index == 0 {
		vnode.next = q.head
		q.head.prev, q.head = vnode, vnode
		return nil
	}

	// add after head
	if index == currSize {
		vnode.prev = q.tail
		q.tail.next, q.tail = vnode, vnode
		return nil
	}

	// 0 < index < q.size
	curr := q.head
	for i := 0; i < q.size && i < index && curr != nil; i++ {
		curr = curr.next
	}
	vnode.prev = curr.prev
	vnode.next = curr
	curr.prev.next, curr.prev = vnode, vnode

	return nil
}

// 0 <= index <= q.size
func (q *DoublyLinkedList[T]) Add_old(index int, v T) error {
	if index > q.size {
		return errors.New("index exceeded list size")
	}
	curr := q.head
	for i := 0; i < q.size && i < index && curr != nil; i++ {
		curr = curr.next
	}
	vnode := Node[T]{
		value: v,
		next:  nil,
		prev:  nil,
	}

	if q.size == 0 {
		q.head = &vnode
		q.tail = &vnode
	} else if index == q.size {
		q.tail.next = &vnode
		vnode.prev = q.tail
		q.tail = &vnode
	} else {
		parent := curr.prev
		if parent != nil {
			parent.next = &vnode
			vnode.prev = parent
			vnode.next = curr
			curr.prev = &vnode
		} else {
			vnode.next = q.head
			q.head.prev = &vnode
			q.head = &vnode
		}
	}
	q.size += 1

	return nil
}

func (q *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		if err := q.Add(e.index, e.value); err != nil {
			return err
		}
		// q.Printd()
	}

	return nil
}

func (q *DoublyLinkedList[T]) Printd() {
	for n := q.head; n != nil; n = n.next {
		fmt.Printf("%v ", n)
	}
	fmt.Println()
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func (l *DoublyLinkedList[T]) PrintRR() {
	fmt.Printf("%s\n", l.PrintReverse())
}

func (l *DoublyLinkedList[T]) PrintFF() {
	fmt.Printf("%s\n", l.PrintForward())
}
