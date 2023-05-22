// Package sort is a package for sorting data.
package sort

import (
	"golang.org/x/exp/constraints"
)

// direction is a type which uses to set the sort direction.
type direction string

const (
	// Desc specifies the sort direction to be descending.
	Desc direction = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc direction = "asc"
)

// Element is an interface which needs to implement for slice sorting
type Element[V constraints.Ordered] interface {
	// Element returns value of element.
	// Type of this element should be constraints.Ordered (int, string, float64 etc)
	// This function will be used for comparing elements
	Element() V
}

type Elements[T constraints.Ordered, V Element[T]] []V

// Person is an example of structure which you want to sort
type Person struct {
	Name string
	Age  int
}

// Element - example of implementation Element() function for sorting of structures
func (p Person) Element() int {
	return p.Age
}

// Slice is a slice for sorting ordered elements
type Slice[V constraints.Ordered] []V
