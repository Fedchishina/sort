<a href="https://pkg.go.dev/github.com/fedchishina/sort"><img src="https://pkg.go.dev/badge/github.com/fedchishina/sort.svg" alt="Go Reference"></a>

Sort data
=======================

Library for sorting data

You can sort any data using these algorithms.
In this library for help and as example we added `type Slice` (slice for sorting ordered elements) and `type Person` with implementation `Element` interface


### BubbleSort
#### ordered slice
Use function `Bubble` for sorting slice with ordered elements:
   - first param is a slice of constraints.Ordered elements (string, int, float64...)
   - second param should be `type direction` (`sort.Asc` or `sort.Desc`)

#### example with ordered slice: 
```
sort.Bubble[int]([]int{4, 5, 8, 7}, sort.Asc)
```

#### slice of structures
Use function `BubbleStruct` for sorting slice with of structures:
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
sort.BubbleStruct[int](persons, sort.Asc)
```

### InsertionSort
#### ordered slice
Use function `Insertion` for sorting slice with ordered elements:
- first param is a slice of constraints.Ordered elements (string, int, float64...)
- second param should be `type direction` (`sort.Asc` or `sort.Desc`)

#### example with ordered slice:
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

## Performance
```
$ GOMAXPROCS=1 go test -bench=. -benchmem -benchtime=10s
BenchmarkBubbleInt              285158943               41.68 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleString           377145843               31.86 ns/op            0 B/op          0 allocs/op
BenchmarkBubbleStructure        295597588               41.47 ns/op            0 B/op          0 allocs/op
BenchmarkInsertionInt           394700175               31.07 ns/op            0 B/op          0 allocs/op
BenchmarkInsertionString        459822885               26.30 ns/op            0 B/op          0 allocs/op
BenchmarkInsertionStructure     400968597               30.22 ns/op            0 B/op          0 allocs/op
BenchmarkMergeInt               58191954               204.5 ns/op           104 B/op          7 allocs/op
BenchmarkMergeString            42165343               276.3 ns/op           160 B/op          5 allocs/op
```
