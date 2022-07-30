package container

type Container[T any] interface {
	// Len gets the number of elements in the container
	Len() int

	// IsEmpty check if the container is empty
	IsEmpty() bool

	// Clear removes all elements from the container
	Clear()

	// Values returns slice of all elements in the container
	Values() []T

	// String returns string representation of the container
	String() string
}
