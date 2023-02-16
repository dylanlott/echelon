package main

import "fmt"

// Echelon is a first-in first-out generic queue in Go.

// Queue defines a first-in first-out queue implementation
type Queue[T any] struct {
	items []T
}

// Peek returns a boolean if T exists and T if it exists
func (q *Queue[T]) Peek() (bool, T) {
	var result T
	if len(q.items) > 0 {
		result = q.items[0]
		return true, result
	}
	return false, result
}

// Pop removes the 0th element of the queue
func (q *Queue[T]) Pop() (T, error) {
	var result T
	if len(q.items) > 0 {
		result := q.items[0]
		q.items = q.items[1:]
		return result, nil
	}
	return result, fmt.Errorf("ErrNoExist")
}

// List returns the list of sorted entries
func (q *Queue[T]) List() []T {
	return q.items
}

// Add inserts an order into the queue
func (q *Queue[T]) Add(item T) ([]T, error) {
	q.items = append(q.items, item)
	return q.items, nil
}
