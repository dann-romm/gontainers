package linkedlist

import "gontainers/container"

var _ container.Iterator[int] = (*Iterator[int])(nil)

type Iterator[T any] struct {
	node      *Node[T]
	isReverse bool
}

func (it *Iterator[T]) Next() container.Iterator[T] {
	if it.isReverse {
		it.node = it.node.prev
	} else {
		it.node = it.node.next
	}
	return it
}

func (it *Iterator[T]) HasNext() bool {
	if it.isReverse {
		return it.node.prev != nil
	}
	return it.node.next != nil
}

func (it *Iterator[T]) Value() T {
	return it.node.value
}

// assert LinkedList[T] to be an Iterable[T]
var _ container.Iterable[int] = (*LinkedList[int])(nil)

func (l *LinkedList[T]) Begin() container.Iterator[T] {
	return &Iterator[T]{node: l.head, isReverse: false}
}

func (l *LinkedList[T]) RBegin() container.Iterator[T] {
	return &Iterator[T]{node: l.tail, isReverse: true}
}
