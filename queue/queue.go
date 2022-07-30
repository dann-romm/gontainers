package queue

import "gontainers/container"

type Queue[T any] interface {
	container.Container[T]

	// Enqueue adds an element to the end of the queue
	Enqueue(values ...T)

	// Dequeue removes the first element from the queue
	Dequeue() (T, bool)

	// Peek returns the first element from the queue
	Peek() (T, bool)
}
