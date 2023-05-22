package sort

import (
	"golang.org/x/exp/constraints"
)

// Bubble sorts data using bubble sort algorithm
//   - first param is a slice of constraints.Ordered elements (string, int, float64...)
//   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
func Bubble[T constraints.Ordered](data Slice[T], direction direction) {
	l := len(data)
	for i := 0; i < l-1; i++ {
		for j := 0; j < l-1-i; j++ {
			if (direction == Asc && !(data[j] < data[j+1])) || (direction == Desc && data[j] < data[j+1]) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// BubbleStruct sorts data of structures using bubble sort algorithm
//   - first param is a slice of structure elements. This slice should implement Element interface
//   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
func BubbleStruct[T constraints.Ordered, V Element[T]](data Elements[T, V], direction direction) {
	for i := 0; i < len(data)-1; i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if (direction == Asc && !(data[j].Element() < data[j+1].Element())) ||
				(direction == Desc && data[j].Element() < data[j+1].Element()) {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}
