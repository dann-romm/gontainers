package circular_queue

import (
	"bytes"
	"fmt"
	"gontainers/queue"
)

const (
	defaultCapacity = 16
)

// assert CircularQueue[T] to be a Queue[T]
var _ queue.Queue[int] = (*CircularQueue[int])(nil)

type CircularQueue[T any] struct {
	data []T
	cap  int
	head int
	tail int
}

func New[T any]() *CircularQueue[T] {
	return &CircularQueue[T]{
		data: make([]T, defaultCapacity),
		cap:  defaultCapacity,
		head: 0,
		tail: 0,
	}
}

func (q *CircularQueue[T]) Push(values ...T) {
	if q.Len()+len(values) > q.cap {
		q.resize()
	}
	for _, value := range values {
		q.data[q.tail] = value
		q.tail = (q.tail + 1) % q.cap
	}
}

func (q *CircularQueue[T]) Pop() (T, bool) {
	if q.Len() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := q.data[q.head]
	q.head = (q.head + 1) % q.cap
	return value, true
}

func (q *CircularQueue[T]) Peek() (T, bool) {
	if q.Len() == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return q.data[q.head], true
}

func (q *CircularQueue[T]) Len() int {
	if q.tail >= q.head {
		return q.tail - q.head
	}
	return q.cap - q.head + q.tail
}

func (q *CircularQueue[T]) resize() {
	oldLength := q.Len()
	newData := make([]T, q.cap*2)
	for i := 0; i < oldLength; i++ {
		newData[i] = q.data[(q.head+i)%q.cap]
	}
	q.data = newData
	q.cap *= 2
	q.head = 0
	q.tail = oldLength
}

func (q *CircularQueue[T]) IsEmpty() bool {
	return q.Len() == 0
}

func (q *CircularQueue[T]) Clear() {
	q.head = 0
	q.tail = 0
}

func (q *CircularQueue[T]) Values() []T {
	values := make([]T, q.Len())
	for i := 0; i < q.Len(); i++ {
		values[i] = q.data[(q.head+i)%q.cap]
	}
	return values
}

func (q *CircularQueue[T]) String() string {
	buf := bytes.NewBufferString("[")
	for i := 0; i < q.Len()-1; i++ {
		buf.WriteString(fmt.Sprintf("%v ", q.data[(q.head+i)%q.cap]))
	}
	buf.WriteString(fmt.Sprintf("%v]", q.data[(q.tail-1+q.cap)%q.cap]))
	return buf.String()
}
