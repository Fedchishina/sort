package sort

import (
	"golang.org/x/exp/constraints"
)

type Direction string

const (
	// Desc specifies the sort direction to be descending.
	Desc Direction = "desc"
	// Asc specifies the sort direction to be ascending.
	Asc Direction = "asc"
)

type Element interface {
	constraints.Ordered
}

type OrderedSlice[V Element] []V

func (x OrderedSlice[V]) Len() int           { return len(x) }
func (x OrderedSlice[V]) Less(i, j int) bool { return x[i] < x[j] }
func (x OrderedSlice[V]) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type Person struct {
	Name string
	Age  int
}
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
