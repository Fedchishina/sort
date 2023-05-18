package sort

func Merge[V Element](data OrderedSlice[V], direction Direction) OrderedSlice[V] {
	if len(data) < 2 {
		return data
	}
	L := Merge(data[:len(data)/2], direction)
	R := Merge(data[len(data)/2:], direction)

	var (
		i, j  int
		final OrderedSlice[V]
	)

	lenL := len(L)
	lenR := len(R)
	for i < lenL && j < lenR {
		if (direction == Asc && L[i] < R[j]) || (direction == Desc && L[i] > R[j]) {
			final = append(final, L[i])
			i++
		} else {
			final = append(final, R[j])
			j++
		}
	}

	for ; i < lenL; i++ {
		final = append(final, L[i])
	}

	for ; j < lenR; j++ {
		final = append(final, R[j])
	}

	return final
}
