package sort

import "sort"

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
