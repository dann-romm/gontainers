package queue_slice

import (
	"bytes"
	"fmt"
	"gontainers/queue"
)

const (
	defaultCapacity = 16
)

var _ queue.Queue[int] = (*Queue[int])(nil)

type Queue[T any] struct {
	values []T
}

func New[T any]() *Queue[T] {
	return &Queue[T]{
		values: make([]T, 0, defaultCapacity),
	}
}

func (q *Queue[T]) Push(values ...T) {
	q.values = append(q.values, values...)
}

func (q *Queue[T]) Pop() (T, bool) {
	if len(q.values) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := q.values[0]
	q.values = q.values[1:]
	return value, true
}

func (q *Queue[T]) Peek() (T, bool) {
	if len(q.values) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return q.values[0], true
}

func (q *Queue[T]) Len() int {
	return len(q.values)
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.values) == 0
}

func (q *Queue[T]) Clear() {
	q.values = make([]T, 0, defaultCapacity)
}

func (q *Queue[T]) Values() []T {
	return q.values
}

func (q *Queue[T]) String() string {
	buf := bytes.NewBufferString("[")
	if len(q.values) > 0 {
		buf.WriteString(fmt.Sprintf("%v", q.values[0]))
	}
	for i := 1; i < len(q.values); i++ {
		buf.WriteString(fmt.Sprintf(", %v", q.values[i]))
	}
	buf.WriteString("]")
	return buf.String()
}
