Sort data
=======================

Library for sorting data

You can sort arrays with these types of elements

```
int | int32 | int64 | float32 | float64 | string
```

## BubbleSort 
use function Bubble from sort package
- first param is array with data
- second param is bool: true - ascending sort, false - descending sort

Example: 
```
var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
bubbleSorted := sort.Bubble(data, true)
fmt.Println("Bubble sort: ", bubbleSorted)
```

## InsertionSort
use function Insertion from sort package
- first param is array with data
- second param is bool: true - ascending sort, false - descending sort

Example:
```
var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
insertionSorted := sort.Insertion(data, false)
fmt.Println("Insertion sort: ", insertionSorted)
```

## MergeSort
use function Merge from sort package
- first param is array with data
- second param is bool: true - ascending sort, false - descending sort

Example:
```
var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
mergeSorted := sort.Merge(data, false)
fmt.Println("Merge sort: ", mergeSorted)
```
