package sort

import "sort"

// Bubble sorts data using bubble sort algorithm
//   - first param is any structure where was implemented `sort.Interface`
//   - second param should be `type Direction` (`sort.Asc` or `sort.Desc`)
func Bubble(data sort.Interface, direction Direction) {
	for i := 0; i < data.Len()-1; i++ {
		for j := 0; j < data.Len()-1-i; j++ {
			if (direction == Asc && !data.Less(j, j+1)) || (direction == Desc && data.Less(j, j+1)) {
				data.Swap(j, j+1)
			}
		}
	}
}
