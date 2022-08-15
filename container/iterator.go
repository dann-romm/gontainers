package container

// Iterator is an interface that defines the methods of an iterator
type Iterator[T any] interface {
	// Next returns the next iterator element
	Next() Iterator[T]

	// HasNext returns true if there are more elements in the iterator
	HasNext() bool

	// Value returns the current element in the iterator
	Value() T
}

// MapIterator is an interface that defines the methods of an iterator for maps
type MapIterator[K comparable, V any] interface {
	// Next returns the next iterator element
	Next() MapIterator[K, V]

	// HasNext returns true if there are more elements in the iterator
	HasNext() bool

	// Value returns the current element in the iterator
	Value() V

	// Key returns the current key in the iterator
	Key() K
}
