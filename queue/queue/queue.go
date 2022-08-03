package _queue

import (
	"gontainers/list/linkedlist"
	"gontainers/queue"
)

// assert Queue[T] to be a Queue[T]
var _ queue.Queue[int] = (*Queue[int])(nil)

type Queue[T any] struct {
	list *linkedlist.LinkedList[T]
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		list: linkedlist.New[T](),
	}
}

func (q *Queue[T]) Push(values ...T) {
	q.list.PushBack(values...)
}

func (q *Queue[T]) Pop() (T, bool) {
	return q.list.PopFront()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.list.Front()
}

func (q *Queue[T]) Len() int {
	return q.list.Len()
}

func (q *Queue[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Queue[T]) Clear() {
	q.list.Clear()
}

func (q *Queue[T]) Values() []T {
	return q.list.Values()
}

func (q *Queue[T]) String() string {
	return q.list.String()
}
