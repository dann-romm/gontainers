package set

import "gontainers/container"

type Set[T any] interface {
	container.Container[T]

	// Add adds an element to the set
	Add(values ...T)

	// Remove removes the element from the set
	Remove(values ...T)

	// Contains returns true if the set contains the given value
	Contains(value T) bool
}
