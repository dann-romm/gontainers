package queue_slice

import "gontainers/container"

// assert Iterator[T] to be a container.Iterator[T]
var _ container.Iterator[int] = (*Iterator[int])(nil)

type Iterator[T any] struct {
	queue *Queue[T]
	index int
}

func (it *Iterator[T]) Next() container.Iterator[T] {
	it.index++
	return it
}

func (it *Iterator[T]) Value() T {
	if it.index < it.queue.Len() {
		return it.queue.values[it.index]
	}
	var zeroValue T
	return zeroValue
}

func (it *Iterator[T]) HasNext() bool {
	return it.index < it.queue.Len()-1
}

// assert Queue[T] to be a ForwardIterable[T]
var _ container.ForwardIterable[int] = (*Queue[int])(nil)

func (q *Queue[T]) Begin() container.Iterator[T] {
	return &Iterator[T]{
		queue: q,
		index: 0,
	}
}
