package stack_list

import (
	"gontainers/list/linkedlist"
	"gontainers/stack"
)

// assert Queue[T] to be a Queue[T]
var _ stack.Stack[int] = (*Stack[int])(nil)

type Stack[T any] struct {
	list *linkedlist.LinkedList[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		list: linkedlist.New[T](),
	}
}

func (q *Stack[T]) Push(values ...T) {
	q.list.PushBack(values...)
}

func (q *Stack[T]) Pop() (T, bool) {
	return q.list.PopBack()
}

func (q *Stack[T]) Peek() (T, bool) {
	return q.list.Back()
}

func (q *Stack[T]) Len() int {
	return q.list.Len()
}

func (q *Stack[T]) IsEmpty() bool {
	return q.list.IsEmpty()
}

func (q *Stack[T]) Clear() {
	q.list.Clear()
}

func (q *Stack[T]) Values() []T {
	return q.list.Values()
}

func (q *Stack[T]) String() string {
	return q.list.String()
}
