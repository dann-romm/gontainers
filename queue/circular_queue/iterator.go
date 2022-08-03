package circular_queue

import (
	"gontainers/container"
)

// assert Iterator[T] to be a container.Iterator[T]
var _ container.Iterator[int] = (*Iterator[int])(nil)

type Iterator[T any] struct {
	q     *CircularQueue[T]
	shift int
}

func (it *Iterator[T]) Next() container.Iterator[T] {
	if it.HasNext() {
		it.shift++
	}
	return it
}

func (it *Iterator[T]) HasNext() bool {
	return it.shift < it.q.Len()
}

func (it *Iterator[T]) Value() T {
	return it.q.data[(it.q.head+it.shift)%it.q.cap]
}

func (q *CircularQueue[T]) Begin() container.Iterator[T] {
	return &Iterator[T]{
		q:     q,
		shift: 0,
	}
}
