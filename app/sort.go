package sort

type Sort interface {
}

type Element interface {
	int | int32 | int64 | float32 | float64 | string
}

func Bubble[V Element](data []V, asc bool) []V {
	var i, j int
	var swap bool
	for i = 0; i < len(data)-1; i++ {
		for j = 0; j < len(data)-1; j++ {
			switch asc {
			case true:
				swap = data[j] > data[j+1]
			case false:
				swap = data[j] < data[j+1]
			}
			if swap {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}

	return data
}

func Insertion[V Element](data []V, asc bool) []V {
	var n = len(data)
	var swap bool
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			switch asc {
			case true:
				swap = data[j-1] > data[j]
			case false:
				swap = data[j-1] < data[j]
			}
			if swap {
				data[j-1], data[j] = data[j], data[j-1]
			}
			j = j - 1
		}
	}
	return data
}

func Merge[V Element](data []V, asc bool) []V {
	if len(data) < 2 {
		return data
	}
	first := Merge(data[:len(data)/2], asc)
	second := Merge(data[len(data)/2:], asc)

	return MergeElements(first, second, asc)
}

func MergeElements[V Element](a, b []V, asc bool) []V {
	var final []V
	var i, j int
	var swap bool
	for i < len(a) && j < len(b) {
		switch asc {
		case true:
			swap = a[i] < b[j]
		case false:
			swap = a[i] > b[j]
		}
		if swap {
			final = append(final, a[i])
			i++
		} else {
			final = append(final, b[j])
			j++
		}
	}

	for ; i < len(a); i++ {
		final = append(final, a[i])
	}

	for ; j < len(b); j++ {
		final = append(final, b[j])
	}
	return final
}
