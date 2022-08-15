package arraylist

import "gontainers/container"

// assert Iterator[T] to be a container.Iterator[T]
var _ container.Iterator[int] = (*Iterator[int])(nil)

type Iterator[T any] struct {
	list      *ArrayList[T]
	index     int
	isReverse bool
}

func (it *Iterator[T]) Next() container.Iterator[T] {
	if it.isReverse {
		it.index--
		return it
	} else {
		it.index++
		return it
	}
}

func (it *Iterator[T]) HasNext() bool {
	if it.isReverse {
		return it.index >= 0
	} else {
		return it.index < it.list.Len()
	}
}

func (it *Iterator[T]) Value() T {
	return it.list.At(it.index)
}

// assert ArrayList[T] to be an Iterable[T]
var _ container.Iterable[int] = (*ArrayList[int])(nil)

func (l *ArrayList[T]) Begin() container.Iterator[T] {
	return &Iterator[T]{
		list:      l,
		index:     0,
		isReverse: false,
	}
}

func (l *ArrayList[T]) RBegin() container.Iterator[T] {
	return &Iterator[T]{
		list:      l,
		index:     l.Len() - 1,
		isReverse: true,
	}
}
