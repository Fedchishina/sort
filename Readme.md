<a href="https://pkg.go.dev/github.com/fedchishina/sort"><img src="https://pkg.go.dev/badge/github.com/fedchishina/sort.svg" alt="Go Reference"></a>

Sort data
=======================

Library for sorting data

### BubbleSort and InsertionSort
You can sort any data using these algorithms.
You need to implement `sort.Interface` in your data structure

In this library for help and as example we added `type OrderedSlice` and `type ByAge []Person` with implementation `sort.Interface`

Use function Bubble or Insertion from sort package:
- first param is any structure where was implemented `sort.Interface`
- second param should be `type Direction` (`sort.Asc` or `sort.Desc`)

#### Example with ordered slice: 
```
data := sort.OrderedSlice[int]([]int{4, 5, 8, 7})
sort.Bubble(data, sort.Asc)
sort.Insertion(data, sort.Asc)
```

#### Example with slice of structures:
```
people := []sort.Person{
    {"Bob", 31},
    {"John", 42},
    {"Michael", 17},
    {"Jenny", 26},
}

sort.Bubble(sort.ByAge(people), sort.Asc)
sort.Insertion(sort.ByAge(people), sort.Asc)
```

### MergeSort
You can sort OrderedSlice data using this algorithm.
Use function Merge for sort package:
- first param is array (`type OrderedSlice`).
- second param should be `type Direction` (`sort.Asc` or `sort.Desc`)

Example:
```
data := sort.OrderedSlice[int]([]int{4, 5, 8, 7})
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
