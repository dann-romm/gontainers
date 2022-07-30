package list

import "gontainers/container"

type List[T any] interface {
	container.Container[T]
	container.Iterable[T]

	// PushBack adds an elements to the end of the list
	PushBack(values ...T)

	// PushFront adds an elements to the front of the list
	PushFront(values ...T)

	// PopBack removes the last element from the list and
	// returns it and true if the list is not empty
	// if the list is empty, return T zero-value and false
	PopBack() (T, bool)

	// PopFront removes the first element from the list
	// returns it and true if the list is not empty
	// if the list is empty, return T zero-value and false
	PopFront() (T, bool)

	// Front returns the first element of the list
	// returns T zero-value and false if the list is empty
	Front() (T, bool)

	// Back returns the last element of the list
	// returns T zero-value and false if the list is empty
	Back() (T, bool)

	// Remove removes the first element from the list that equals to the given value
	Remove(value T)

	// RemoveAt removes the element at the given index
	RemoveAt(index int)

	// // InsertBefore inserts the given value before the marker element
	// InsertBefore(value T, marker T)
	//
	// // InsertAfter inserts the given value after the marker element
	// InsertAfter(value T, marker T)

	// Contains returns true if the list contains the given value
	Contains(value T) bool

	// // Swap swaps the elements at the given indexes
	// Swap(index1, index2 int)
}
