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
	panic("implement me!")
}

func (sortAlgorithm *SortAlgorithmImpl) MergeSort(arr []int) []int {
	panic("implement me!")
}

func (sortAlgorithm *SortAlgorithmImpl) InsertionSort(arr []int) []int {
	panic("implement me!")
}