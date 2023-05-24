<a href="https://pkg.go.dev/github.com/fedchishina/sort"><img src="https://pkg.go.dev/badge/github.com/fedchishina/sort.svg" alt="Go Reference"></a>

Data sorting
=======================

Library for sorting data

You can sort any data using these algorithms.
In this library for help and as example we added `type Slice` (slice for sorting ordered elements) and `type Person` with implementation `Element` interface


<!-- TOC -->
* [Data sorting](#data-sorting)
    * [BubbleSort](#bubblesort)
      * [ordered slice](#ordered-slice)
      * [example with ordered slice](#example-with-ordered-slice-)
      * [slice of structures](#slice-of-structures)
      * [example with slice of structures](#example-with-slice-of-structures)
    * [InsertionSort](#insertionsort)
      * [ordered slice](#ordered-slice-1)
      * [example with ordered slice](#example-with-ordered-slice)
      * [slice of structures](#slice-of-structures-1)
      * [example with slice of structures:](#example-with-slice-of-structures-1)
    * [MergeSort](#mergesort)
    * [QuickSort](#quicksort)
      * [ordered slice](#ordered-slice-2)
      * [example with ordered slice](#example-with-ordered-slice-1)
      * [slice of structures](#slice-of-structures-2)
      * [example with slice of structures](#example-with-slice-of-structures-2)
  * [Performance](#performance)
<!-- TOC -->

### BubbleSort
#### ordered slice
Use function `Bubble` for sorting slice with ordered elements:
   - first param is a slice of constraints.Ordered elements (string, int, float64...)
   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)

#### example with ordered slice 
```
sort.Bubble[int]([]int{4, 5, 8, 7}, sort.Asc)
```

#### slice of structures
Use function `BubbleStruct` for sorting slice with of structures:
   - first param is a slice of structure elements. This slice should implement `Element` interface
   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)
#### example with slice of structures
```
persons := []sort.Person{
    {"John", 42},
    {"Bob", 31},
    {"Jenny", 26},
    {"Michael", 17},
}
sort.BubbleStruct[int](persons, sort.Asc)
```

### InsertionSort
#### ordered slice
Use function `Insertion` for sorting slice with ordered elements:
- first param is a slice of constraints.Ordered elements (string, int, float64...)
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)

#### example with ordered slice
```
sort.Insertion[int]([]int{4, 5, 8, 7}, sort.Asc)
```

#### slice of structures
Use function `InsertionStruct` for sorting slice with of structures:
- first param is a slice of structure elements. This slice should implement `Element` interface
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)
#### example with slice of structures:
```
persons := []sort.Person{
    {"John", 42},
    {"Bob", 31},
    {"Jenny", 26},
    {"Michael", 17},
}
sort.InsertionStruct[int](persons, sort.Asc)
```

### MergeSort
You can sort Slice data using this algorithm.
Use function Merge for sort package:
- first param is array (`type Slice`).
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)

Example:
```
data := sort.Slice[int]([]int{4, 5, 8, 7})
result := sort.Merge(data, sort.Asc)
fmt.Println("Merge sort: ", result)
```

### QuickSort
#### ordered slice
Use function `Quick` for sorting slice with ordered elements:
- first param is a slice of constraints.Ordered elements (string, int, float64...)
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)

#### example with ordered slice
```
sort.Quick[int]([]int{4, 5, 8, 7}, sort.Asc)
```

#### slice of structures
Use function `QuickStruct` for sorting slice with of structures:
- first param is a slice of structure elements. This slice should implement `Element` interface
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)
#### example with slice of structures
```
persons := []sort.Person{
    {"John", 42},
    {"Bob", 31},
    {"Jenny", 26},
    {"Michael", 17},
}
sort.QuickStruct[int](persons, sort.Asc)
```

## Performance
```
$ go test -bench=. -benchmem
BenchmarkBubbleInt-8                            84884545                14.13 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleString-8                         58206878                20.93 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleStruct/structure-8               33827671                36.03 ns/op            0 B/op          0 allocs/op
BenchmarkInsertionInt-8                         120945201                9.541 ns/op           0 B/op          0 allocs/op
BenchmarkInsertionString-8                      68478008                17.75 ns/op            0 B/op          0 allocs/op
BenchmarkInsertionStruct/structure-8            34708350                36.13 ns/op            0 B/op          0 allocs/op
BenchmarkMergeInt-8                             11594480                99.62 ns/op           64 B/op          3 allocs/op
BenchmarkMergeString-8                           9863658               116.4 ns/op            80 B/op          2 allocs/op
BenchmarkQuickInt-8                             37386302                32.31 ns/op            0 B/op          0 allocs/op
BenchmarkQuickString-8                          26831978                42.77 ns/op            0 B/op          0 allocs/op
BenchmarkQuickStruct/structure-8                13882881                88.05 ns/op            0 B/op          0 allocs/op
```
