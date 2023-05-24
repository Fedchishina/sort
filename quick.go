package sort

import "golang.org/x/exp/constraints"

func Quick[T constraints.Ordered](data Slice[T], direction direction) {
	quickSort(data, direction, 0, len(data)-1)
}

func quickSort[T constraints.Ordered](data Slice[T], direction direction, minIndex int, maxIndex int) {
	if minIndex >= maxIndex || len(data) == 0 {
		return
	}

	pivotIndex := getPivotIndex(data, direction, minIndex, maxIndex)
	quickSort(data, direction, minIndex, pivotIndex-1)
	quickSort(data, direction, pivotIndex+1, maxIndex)
}

func getPivotIndex[T constraints.Ordered](data Slice[T], direction direction, minIndex int, maxIndex int) int {
	pivot := minIndex - 1

	for i := minIndex; i <= maxIndex; i++ {
		if (data[i] < data[maxIndex] && direction == Asc) || (data[i] > data[maxIndex] && direction == Desc) {
			pivot++
			data[pivot], data[i] = data[i], data[pivot]
		}
	}

	pivot++
	data[pivot], data[maxIndex] = data[maxIndex], data[pivot]

	return pivot
}

func QuickStruct[T constraints.Ordered, V Element[T]](data Elements[T, V], direction direction) {
	quickSortStruct(data, direction, 0, len(data)-1)
}

func quickSortStruct[T constraints.Ordered, V Element[T]](data Elements[T, V], direction direction, minIndex int, maxIndex int) {
	if minIndex >= maxIndex || len(data) == 0 {
		return
	}

	pivotIndex := getPivotIndexStruct(data, direction, minIndex, maxIndex)
	quickSortStruct(data, direction, minIndex, pivotIndex-1)
	quickSortStruct(data, direction, pivotIndex+1, maxIndex)
}

func getPivotIndexStruct[T constraints.Ordered, V Element[T]](data Elements[T, V], direction direction, minIndex int, maxIndex int) int {
	pivot := minIndex - 1

	for i := minIndex; i <= maxIndex; i++ {
		if (data[i].Element() < data[maxIndex].Element() && direction == Asc) ||
			(data[i].Element() > data[maxIndex].Element() && direction == Desc) {
			pivot++
			data[pivot], data[i] = data[i], data[pivot]
		}
	}

	pivot++
	data[pivot], data[maxIndex] = data[maxIndex], data[pivot]

	return pivot
}
