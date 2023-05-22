package sort

import (
	"golang.org/x/exp/constraints"
)

// Insertion sorts data using insertion sort algorithm
//   - first param is a slice of constraints.Ordered elements (string, int, float64...)
//   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
func Insertion[T constraints.Ordered](data Slice[T], direction direction) {
	l := len(data)
	for i := 1; i < l; i++ {
		j := i
		for j > 0 {
			if (direction == Asc && !(data[j-1] < data[j])) || (direction == Desc && data[j-1] < data[j]) {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
}

// InsertionStruct sorts data of structures using insertion sort algorithm
//   - first param is a slice of structure elements. This slice should implement Element interface
//   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
func InsertionStruct[T constraints.Ordered, V Element[T]](data Elements[T, V], direction direction) {
	l := len(data)
	for i := 1; i < l; i++ {
		j := i
		for j > 0 {
			if (direction == Asc && !(data[j-1].Element() < data[j].Element())) ||
				(direction == Desc && data[j-1].Element() < data[j].Element()) {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
}
