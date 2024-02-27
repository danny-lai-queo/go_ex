package main

import (
	"errors"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (*T, error) {
	if !s.IsEmpty() {
		popped := s.elements[len(s.elements)-1]
		s.elements = s.elements[:len(s.elements)-1]
		return &popped, nil
	}

	return nil, errors.New("stack: empty stack pop error")
}
