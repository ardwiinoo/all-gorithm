package main

import (
	"ardwiinoo/all-gorithm/sort"
	"fmt"
)

func main() {
	arr := []int{4, 6, 3, 33, 9, 100}

	sorter := sort.NewSortAlgorithm()

	bubbleSort := sorter.BubbleSort(arr)
	fmt.Println("Result Bubble Sort: ", bubbleSort)

	/*
		[4, 6, 3, 33, 9, 100] -> Compare: 4, 6
		[4, 6, 3, 33, 9, 100] -> Compare: 6, 3
		[4, 3, 6, 33, 9, 100] -> Compare: 6, 33
		[4, 3, 6, 33, 9, 100] -> Compare: 33, 9
		[4, 3, 6, 9, 33, 100] -> Compare: 33, 100
		[4, 3, 6, 9, 33, 100] -> End of pass 1
		[4, 3, 6, 9, 33, 100] -> Compare: 4, 3
		[3, 4, 6, 9, 33, 100] -> Compare: 4, 6
		[3, 4, 6, 9, 33, 100] -> Compare: 6, 9
		[3, 4, 6, 9, 33, 100] -> Compare: 9, 33
		[3, 4, 6, 9, 33, 100] -> Compare: 33, 100
		[3, 4, 6, 9, 33, 100] -> End of pass 2
		[3, 4, 6, 9, 33, 100] -> Sorted
	*/

	quickSort := sorter.QuickSort(arr)
	fmt.Println("Result Quick Sort: ", quickSort)

	/*
		[4, 6, 3, 33, 9, 100] -> Pivot: 100
		[4, 6, 3, 9, 33, 100] -> Pivot: 33
		[4, 6, 3, 9, 33, 100] -> Pivot: 9
		[3, 4, 6, 9, 33, 100] -> Sorted
	*/

}