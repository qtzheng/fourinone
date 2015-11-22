package utils

func QuickSort(sortArray []int, left, right int) {
	if left < right {
		pos := partition(sortArray, left, right)
		QuickSort(sortArray, left, pos-1)
		QuickSort(sortArray, pos+1, right)
	}
}
func partition(sortArray []int, left, right int) int {
	key := sortArray[right]
	i := left - 1

	for j := left; j < right; j++ {
		if sortArray[j] <= key {
			i++
			swap(sortArray, i, j)
		}
	}

	swap(sortArray, i+1, right)

	return i + 1
}
func swap(sortArray []int, a, b int) {
	sortArray[a], sortArray[b] = sortArray[b], sortArray[a]
}
