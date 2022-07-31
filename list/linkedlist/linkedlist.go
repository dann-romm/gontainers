package linkedlist

import (
	"bytes"
	"fmt"
	"gontainers/container"
	"gontainers/list"
)

// assert LinkedList[T] to be a Container[T]
var _ container.Container[int] = (*LinkedList[int])(nil)

// assert LinkedList[T] to be a List[T]
var _ list.List[int] = (*LinkedList[int])(nil)

type LinkedList[T any] struct {
	len  int
	head *Node[T]
	tail *Node[T]
}

func New[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) initWithValue(value T) {
	l.head = newNodeValue(value)
	l.head.prev = l.head
	l.tail = l.head
	l.len = 1
}

func (l *LinkedList[T]) PushFront(values ...T) {
	if len(values) == 0 {
		return
	}
	if l.head == nil {
		l.initWithValue(values[len(values)-1])
		values = values[:len(values)-1] // TODO: is this operation takes O(1) time?
	}
	for i := len(values) - 1; i >= 0; i-- {
		l.head.prev = newNode(values[i], l.head, l.tail)
		l.head = l.head.prev
	}
	l.len += len(values)
}

func (l *LinkedList[T]) PushBack(values ...T) {
	if len(values) == 0 {
		return
	}
	if l.head == nil {
		l.initWithValue(values[0])
		values = values[1:] // TODO: is this operation takes O(1) time?
	}
	for _, value := range values {
		l.tail.next = newNode(value, nil, l.tail)
		l.tail = l.tail.next
	}
	l.head.prev = l.tail
	l.len += len(values)
}

func (l *LinkedList[T]) PopFront() (T, bool) {
	if l.head == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := l.head.value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = l.tail
	}
	l.len--
	return value, true
}

func (l *LinkedList[T]) PopBack() (T, bool) {
	if l.tail == nil {
		var zeroValue T
		return zeroValue, false
	}
	value := l.tail.value
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	l.len--
	return value, true
}

func (l *LinkedList[T]) Front() (T, bool) {
	if l.head == nil {
		var zeroValue T
		return zeroValue, false
	}
	return l.head.value, true
}

func (l *LinkedList[T]) Back() (T, bool) {
	if l.tail == nil {
		var zeroValue T
		return zeroValue, false
	}
	return l.tail.value, true
}

// Remove removes the first occurrence of the value
func (l *LinkedList[T]) Remove(value T) {
	if l.head == nil {
		return
	}
	if any(l.head.value) == any(value) {
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.head.prev = l.tail
		}
		l.len--
		return
	}
	node := l.head
	for node.next != nil {
		if any(node.next.value) == any(value) {
			node.next = node.next.next
			if node.next != nil {
				node.next.prev = node
			}
			l.len--
			return
		}
		node = node.next
	}
}

func (l *LinkedList[T]) RemoveAt(index int) {
	if l.head == nil {
		return
	}
	if index == 0 {
		if l.head == l.tail {
			l.head = nil
			l.tail = nil
		} else {
			l.head = l.head.next
			l.head.prev = l.tail
		}
		l.len--
		return
	}
	node := l.head
	for i := 0; i < index-1; i++ {
		node = node.next
		if node == nil {
			return
		}
	}
	node.next = node.next.next
	if node.next != nil {
		node.next.prev = node
	}
	l.len--
}

func (l *LinkedList[T]) Contains(value T) bool {
	if l.head == nil {
		return false
	}
	node := l.head
	for node != nil {
		if any(node.value) == any(value) {
			return true
		}
		node = node.next
	}
	return false
}

func (l *LinkedList[T]) Len() int {
	return l.len
}

func (l *LinkedList[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *LinkedList[T]) Clear() {
	l.len = 0
	l.head = nil
	l.tail = nil
}

func (l *LinkedList[T]) Values() []T {
	if l.head == nil {
		var zeroValue []T
		return zeroValue
	}
	values := make([]T, l.len)
	node := l.head
	for i := 0; i < l.len; i++ {
		values[i] = node.value
		node = node.next
	}
	return values
}

func (l *LinkedList[T]) String() string {
	buf := bytes.NewBufferString("[")
	node := l.head
	for i := 0; i < l.len-1; i++ {
		buf.WriteString(fmt.Sprintf("%v", node.value))
		node = node.next
	}
	buf.WriteString(fmt.Sprintf("%v]", node.value))
	return buf.String()
}
