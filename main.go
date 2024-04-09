package main

import (
	"ardwiinoo/all-gorithm/sort"
	"fmt"
)

func main() {
	arr := []int{4, 6, 3, 33, 9, 100}

	sorter := sort.NewSortAlgorithm()

	sortedArr := sorter.BubbleSort(arr)

	fmt.Println("Result: ", sortedArr)
}