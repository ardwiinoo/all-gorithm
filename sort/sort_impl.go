package sort

type SortAlgorithmImpl struct {
}

func NewSortAlgorithm() SortAlgorithm {
	return &SortAlgorithmImpl{}
}

func (sortAlgorithm *SortAlgorithmImpl) BubbleSort(arr []int) []int {
	n := len(arr)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func (sortAlgorithm *SortAlgorithmImpl) QuickSort(arr []int) []int {
	return quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, low int, high int) []int {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
	return arr
}

func partition(arr []int, low int, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func (sortAlgorithm *SortAlgorithmImpl) MergeSort(arr []int) []int {
	panic("implement me!")
}

func (sortAlgorithm *SortAlgorithmImpl) InsertionSort(arr []int) []int {
	panic("implement me!")
}