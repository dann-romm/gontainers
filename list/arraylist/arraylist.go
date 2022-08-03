package arraylist

import (
	"gontainers/list"
)

// assert ArrayList[T] to be a List[T]
var _ list.List[int] = (*ArrayList[int])(nil)

const (
	defaultCapacity = 16
)

type ArrayList[T any] struct {
	len    int
	values []T
}

func New[T any]() *ArrayList[T] {
	return &ArrayList[T]{
		len:    0,
		values: make([]T, 0, defaultCapacity),
	}
}

func (l *ArrayList[T]) PushBack(values ...T) {
	l.values = append(l.values, values...)
	l.len += len(values)
}

func (l *ArrayList[T]) PushFront(values ...T) {
	l.values = append(values, l.values...)
	l.len += len(values)
}

func (l *ArrayList[T]) PopBack() (T, bool) {
	if l.len == 0 {
		var zeroValue T
		return zeroValue, false
	}

	l.len--
	value := l.values[l.len]
	l.values = l.values[:l.len]
	return value, true
}

func (l *ArrayList[T]) PopFront() (T, bool) {
	if l.len == 0 {
		var zeroValue T
		return zeroValue, false
	}

	value := l.values[0]
	l.values = l.values[1:]
	l.len--
	return value, true
}

func (l *ArrayList[T]) Front() (T, bool) {
	if l.len == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return l.values[0], true
}

func (l *ArrayList[T]) Back() (T, bool) {
	if l.len == 0 {
		var zeroValue T
		return zeroValue, false
	}

	return l.values[l.len-1], true
}

func (l *ArrayList[T]) Remove(value T) {
	for i, v := range l.values {
		if any(v) == any(value) {
			l.values = append(l.values[:i], l.values[i+1:]...)
			l.len--
			return
		}
	}
}

func (l *ArrayList[T]) RemoveAt(index int) {
	l.values = append(l.values[:index], l.values[index+1:]...)
	l.len--
}

func (l *ArrayList[T]) At(index int) T {
	if index < 0 || index >= l.len {
		var zeroValue T
		return zeroValue
	}
	return l.values[index]
}

func (l *ArrayList[T]) Contains(value T) bool {
	for _, v := range l.values {
		if any(v) == any(value) {
			return true
		}
	}
	return false
}

func (l *ArrayList[T]) Len() int {
	return l.len
}

func (l *ArrayList[T]) IsEmpty() bool {
	return l.len == 0
}

func (l *ArrayList[T]) Clear() {
	l.values = make([]T, 0, defaultCapacity)
	l.len = 0
}

func (l *ArrayList[T]) Values() []T {
	return l.values
}

func (l *ArrayList[T]) String() string {
	panic("not implemented")
}
