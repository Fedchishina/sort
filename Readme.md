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
