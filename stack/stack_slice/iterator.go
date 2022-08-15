package stack_slice

import "gontainers/container"

// assert Iterator[T] to be a container.Iterator[T]
var _ container.Iterator[int] = (*Iterator[int])(nil)

type Iterator[T any] struct {
	stack *Stack[T]
	index int
}

func (it *Iterator[T]) Next() container.Iterator[T] {
	it.index--
	return it
}

func (it *Iterator[T]) Value() T {
	if it.index > -1 {
		return it.stack.values[it.index]
	}
	var zeroValue T
	return zeroValue
}

func (it *Iterator[T]) HasNext() bool {
	return it.index > 0
}

// assert Queue[T] to be a ForwardIterable[T]
var _ container.ForwardIterable[int] = (*Stack[int])(nil)

func (s *Stack[T]) Begin() container.Iterator[T] {
	return &Iterator[T]{
		stack: s,
		index: s.Len() - 1,
	}
}
