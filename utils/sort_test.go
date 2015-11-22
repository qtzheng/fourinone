package utils

import (
	"fmt"
	"sort"
	"testing"
)

func TestA(t *testing.T) {
	arr := []int{1, 4, 26, 67, 45, 347, 32456, 86, 2354, 56, 23452, 45, 7}
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
func BenchmarkA(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 4, 26, 67, 45, 347, 32456, 86, 2354, 56, 23452, 45, 7}
		QuickSort(arr, 0, len(arr)-1)
	}
}
func BenchmarkB(b *testing.B) {
	for i := 0; i < b.N; i++ {
		arr := []int{1, 4, 26, 67, 45, 347, 32456, 86, 2354, 56, 23452, 45, 7}
		sort.Ints(arr)
	}
}
