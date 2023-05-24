package sort

import (
	_ "net/http/pprof"

	"golang.org/x/exp/constraints"
)

// Merge sorts data using merge sort algorithm
//   - first param is array (`type Slice`).
//   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
func Merge[V constraints.Ordered](data Slice[V], direction direction) Slice[V] {
	if len(data) < 2 {
		return data
	}
	L := Merge(data[:len(data)/2], direction)
	R := Merge(data[len(data)/2:], direction)

	var (
		i, j int
	)

	final := make(Slice[V], len(data))
	fIndex := 0
	lenL := len(L)
	lenR := len(R)
	for i < lenL && j < lenR {
		if (direction == Asc && L[i] < R[j]) || (direction == Desc && L[i] > R[j]) {
			final[fIndex] = L[i]
			i++
		} else {
			final[fIndex] = R[j]
			j++
		}

		fIndex++
	}

	for ; i < lenL; i++ {
		final[fIndex] = L[i]
		fIndex++
	}

	for ; j < lenR; j++ {
		final[fIndex] = R[j]
		fIndex++
	}

	for k, _ := range data {
		data[k] = final[k]
	}

	return data
}
