package main

import (
	"fmt"
	sort "github.com/Fedchishina/sort/app"
)

func main() {
	var data = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	//var data = []string{"", "Hello", "foo", "bar", "foo", "f00", "%*&^*&^&", "***"}

	bubbleSorted := sort.Bubble(data, false)
	fmt.Println("Bubble sort: ", bubbleSorted)
	insertionSorted := sort.Insertion(data, false)
	fmt.Println("Insertion sort: ", insertionSorted)
	mergeSorted := sort.Merge(data, false)
	fmt.Println("Merge sort: ", mergeSorted)
}
