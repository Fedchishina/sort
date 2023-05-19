// Package sort is a package for sorting data.
package sort

import (
	"golang.org/x/exp/constraints"
)

// Direction is a type which uses to set the sort direction.
type Direction string

const (
	// Desc specifies the sort direction to be descending.
	Desc Direction = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc Direction = "asc"
)

// Element is a type elements of Ordered slice (int, string, float64, etc)
type Element interface {
	constraints.Ordered
}

// OrderedSlice is a slice for sorting ordered elements
type OrderedSlice[V Element] []V

// Len returns a len of Ordered slice
func (x OrderedSlice[V]) Len() int { return len(x) }

// Less compares two elements. If i-element less j-element function returns true.
func (x OrderedSlice[V]) Less(i, j int) bool { return x[i] < x[j] }

// Swap swaps i and j elements
func (x OrderedSlice[V]) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

// Person is an example of structure which you want to sort
type Person struct {
	Name string
	Age  int
}

// ByAge is slice of Person elements
type ByAge []Person

// Len returns a len of ByAge slice
func (a ByAge) Len() int { return len(a) }

// Swap swaps i and j elements
func (a ByAge) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

// Less compares two elements. If i-element less j-element function returns true.
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
