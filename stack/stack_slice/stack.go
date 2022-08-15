package stack_slice

import (
	"bytes"
	"fmt"
	"gontainers/stack"
)

const (
	defaultCapacity = 16
)

var _ stack.Stack[int] = (*Stack[int])(nil)

type Stack[T any] struct {
	values []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		values: make([]T, 0, defaultCapacity),
	}
}

func (s *Stack[T]) Push(values ...T) {
	s.values = append(s.values, values...)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.values) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	value := s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return value, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.values) == 0 {
		var zeroValue T
		return zeroValue, false
	}
	return s.values[len(s.values)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.values)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Stack[T]) Clear() {
	s.values = make([]T, 0, defaultCapacity)
}

func (s *Stack[T]) Values() []T {
	return s.values
}

func (s *Stack[T]) String() string {
	buf := bytes.NewBufferString("[")
	if len(s.values) > 0 {
		buf.WriteString(fmt.Sprintf("%v", s.values[0]))
	}
	for i := 1; i < len(s.values); i++ {
		buf.WriteString(fmt.Sprintf(", %v", s.values[i]))
	}
	buf.WriteString("]")
	return buf.String()
}
