package sort

type SortAlgorithm interface {
	BubbleSort(arr []int) []int
	QuickSort(arr []int) []int
	MergeSort(arr []int) []int
	InsertionSort(arr []int) []int
}