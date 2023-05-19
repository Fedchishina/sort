package sort

import "sort"

// Insertion sorts data using insertion sort algorithm
//   - first param is any structure where was implemented `sort.Interface`
//   - second param should be `type Direction` (`sort.Asc` or `sort.Desc`)
func Insertion(data sort.Interface, direction Direction) {
	for i := 1; i < data.Len(); i++ {
		j := i
		for j > 0 {
			if direction == Asc && !data.Less(j-1, j) || (direction == Desc && data.Less(j-1, j)) {
				data.Swap(j-1, j)
			}
			j = j - 1
		}
	}
}
