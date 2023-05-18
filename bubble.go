package sort

import "sort"

func Bubble(data sort.Interface, direction Direction) {
	for i := 0; i < data.Len()-1; i++ {
		for j := 0; j < data.Len()-1-i; j++ {
			if (direction == Asc && !data.Less(j, j+1)) || (direction == Desc && data.Less(j, j+1)) {
				data.Swap(j, j+1)
			}
		}
	}
}
