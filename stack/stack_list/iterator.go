package stack_list

import (
	"gontainers/container"
)

// assert Queue[T] to be a ForwardIterable[T]
var _ container.ForwardIterable[int] = (*Stack[int])(nil)

func (q *Stack[T]) Begin() container.Iterator[T] {
	return q.list.RBegin()
}
