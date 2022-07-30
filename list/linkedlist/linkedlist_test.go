package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// func createLinkedList() *LinkedList[int] {
//
// }

func TestNew(t *testing.T) {
	l := New[int]()

	assert.Nil(t, l.head, "head of new list should be nil")
	assert.Nil(t, l.tail, "tail of new list should be nil")
	assert.Equal(t, 0, l.len, "len of new list should be 0")
}

func AssertCorrectList[T any](t *testing.T, l *LinkedList[T]) {
	if l.head == nil {
		assert.Nil(t, l.tail, "tail of empty list should be nil")
		assert.Equal(t, 0, l.len, "len of empty list should be 0")
	} else {
		assert.NotNil(t, l.tail, "tail of non-empty list should not be nil")
		assert.Equal(t, l.head.prev, l.tail, "head.prev should be tail")
		assert.Nil(t, l.tail.next, "tail.next should be nil")
	}
}

func AssertListsEqual[T any](t *testing.T, l1, l2 *LinkedList[T]) {
	assert.Equal(t, l1.len, l2.len)
	node1, node2 := l1.head, l2.head
	for node1 != nil && node2 != nil {
		assert.Equal(t, node1.value, node2.value, "values of lists should be equal")
		node1 = node1.next
		node2 = node2.next
	}
	assert.Nil(t, node1, "node1 should be nil")
	assert.Nil(t, node2, "node2 should be nil")
}

func TestLinkedList_PushBack(t *testing.T) {
	l := New[int]()

	l.PushBack(1)
	AssertCorrectList(t, l)

	assert.Equal(t, 1, l.len, "len of list should be 1")
	assert.Equal(t, 1, l.head.value, "value of head should be 1")
	assert.Equal(t, 1, l.tail.value, "value of tail should be 1")

	l.PushBack(2)
	AssertCorrectList(t, l)

	assert.Equal(t, 2, l.len, "len of list should be 2")
	assert.Equal(t, 1, l.head.value, "value of head should be 1")
	assert.Equal(t, 2, l.tail.value, "value of tail should be 2")

	l.PushBack(3, 4, 5, 6)
	AssertCorrectList(t, l)

	assert.Equal(t, 6, l.len, "len of list should be 6")
	assert.Equal(t, 1, l.head.value, "value of head should be 1")
	assert.Equal(t, 6, l.tail.value, "value of tail should be 6")

	l2 := New[int]()
	l2.PushBack(1, 2, 3, 4, 5, 6)
	AssertCorrectList(t, l2)

	assert.Equal(t, 6, l2.len, "len of list should be 6")
	assert.Equal(t, 1, l2.head.value, "value of head should be 1")
	assert.Equal(t, 6, l2.tail.value, "value of tail should be 6")
}

func TestLinkedList_PushFront(t *testing.T) {
	l := New[int]()

	l.PushFront(1)
	AssertCorrectList(t, l)

	assert.Equal(t, 1, l.len, "len of list should be 1")
	assert.Equal(t, 1, l.head.value, "value of head should be 1")
	assert.Equal(t, 1, l.tail.value, "value of tail should be 1")

	l.PushFront(2)
	AssertCorrectList(t, l)

	assert.Equal(t, 2, l.len, "len of list should be 2")
	assert.Equal(t, 2, l.head.value, "value of head should be 2")
	assert.Equal(t, 1, l.tail.value, "value of tail should be 1")

	l.PushFront(3, 4, 5, 6)
	AssertCorrectList(t, l)

	assert.Equal(t, 6, l.len, "len of list should be 6")
	assert.Equal(t, 3, l.head.value, "value of head should be 3")
	assert.Equal(t, 1, l.tail.value, "value of tail should be 1")

	l2 := New[int]()
	l2.PushFront(1, 2, 3, 4, 5, 6)
	AssertCorrectList(t, l2)

	assert.Equal(t, 6, l2.len, "len of list should be 6")
	assert.Equal(t, 1, l2.head.value, "value of head should be 1")
	assert.Equal(t, 6, l2.tail.value, "value of tail should be 6")

	l3 := New[int]()
	l3.PushBack(1, 2, 3, 4, 5, 6)
	AssertListsEqual(t, l2, l3)
}
