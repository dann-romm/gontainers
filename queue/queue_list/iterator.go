package queue_list

import (
	"gontainers/container"
)

// assert Queue[T] to be a ForwardIterable[T]
var _ container.ForwardIterable[int] = (*Queue[int])(nil)

func (q *Queue[T]) Begin() container.Iterator[T] {
	return q.list.Begin()
}
