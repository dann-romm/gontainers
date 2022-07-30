package linkedlist

import "gontainers/container"

var _ container.Iterator[int] = (*Node[int])(nil)

type Node[T any] struct {
	value T
	next  *Node[T]
	prev  *Node[T]
}

func newNode[T any](value T, next, prev *Node[T]) *Node[T] {
	return &Node[T]{
		value: value,
		prev:  prev,
		next:  next,
	}
}

func newNodeValue[T any](value T) *Node[T] {
	return &Node[T]{value: value}
}
