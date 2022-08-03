package queue

import "gontainers/container"

type Queue[T any] interface {
	container.Container[T]
	container.ForwardIterable[T]

	// Push adds a bunch of elements to the end of the queue
	Push(values ...T)

	// Pop removes the first element from the queue
	Pop() (T, bool)

	// Peek returns the first element from the queue
	Peek() (T, bool)
}
