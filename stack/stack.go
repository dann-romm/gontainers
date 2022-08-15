package stack

import "gontainers/container"

type Stack[T any] interface {
	container.Container[T]
	container.ForwardIterable[T]

	// Push adds a bunch of elements to the top of the stack
	Push(values ...T)

	// Pop removes the first element from the stack
	Pop() (T, bool)

	// Peek returns the first element from the stack
	Peek() (T, bool)
}
