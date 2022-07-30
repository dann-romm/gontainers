package container

// Iterator is an interface that defines the methods of a bidirectional iterator
type Iterator[T any] interface {
	// Next returns the next iterator element
	Next() Iterator[T]

	// HasNext returns true if there are more elements in the iterator
	HasNext() bool

	// Value returns the current element in the iterator
	Value() T
}

// usage:
// for it := container.Begin(); it.HasNext(); it = it.Next() {
//     value := it.Value()
// }
