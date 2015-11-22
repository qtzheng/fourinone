package fourinone

import (
	"github.com/qtzheng/fourinone/utils"
)

type ListInt interface {
	Add(int)
	AddArray([]int)
	Size() int
	Set(int, int)
	Get(int) int
	Sort()
	SortArray([]int) []int
	ToArray() []int
	Order([]int)
}
type ArrayInt struct {
	arrsize      int
	arrayAdapter *ArrayAdapter
	arrint       []int
	arrindex     int
}

func NewArrayInt() *ArrayInt {
	array := &ArrayInt{
		arrsize:  0x20000,
		arrint:   make([]int, 0x20000),
		arrindex: 0,
	}
	ad := NewArrayAdapter()
	ad.Add(array.arrint)
	return array
}
func (a *ArrayInt) Add(i int) {
	if a.arrindex == len(a.arrint) {
		a.arrint = make([]int, a.arrsize)
		a.arrayAdapter.auto()
		a.arrayAdapter.Add(a.arrint)
		a.arrindex = 0
	}
	a.arrint[a.arrindex] = i
	a.arrindex++
}
func (a *ArrayInt) AddArray(arr []int) {
	for _, v := range arr {
		a.Add(v)
	}
}
func (a *ArrayInt) Size() int {
	return (a.arrayAdapter.objinit-1)*a.arrsize + a.arrindex
}
func (a *ArrayInt) Set(index, i int) {
	if arr, ok := a.arrayAdapter.abjarr[index/a.arrsize].([]int); ok {
		arr[index%a.arrsize] = i
	}
}
func (a *ArrayInt) Get(index int) int {
	if arr, ok := a.arrayAdapter.abjarr[index/a.arrsize].([]int); ok {
		return arr[index%a.arrsize]
	}
	return 0
}
func (a *ArrayInt) Sort() {
	a.arrsort(0, a.Size())
}
func (a *ArrayInt) arrsort(k, m int) {
	i, j := k, m
	for ; i < j; i++ {
		vai := a.Get(i)
		for vai <= a.Get(i+1) {
			if j == i+1 {
				break
			}
			vai1 := a.Get(i + 1)
			a.Set(i+1, a.Get(j))
			a.Set(j, vai1)
			j--
		}
		if vai > a.Get(i+1) {
			a.Set(i, a.Get(i+1))
			a.Set(i+1, vai)
		}
	}
	if i-1 > k {
		a.arrsort(k, i-1)
	}
	if m > i {
		a.arrsort(i, m)
	}
}
func (a *ArrayInt) SortArray(arr []int) []int {
	a.intsort(0, len(arr)-1, arr)
	return arr
}
func (a *ArrayInt) intsort(k, m int, arr []int) {
	j, i := m, k
	for ; i < k; i++ {
		vai := a.Get(i)
		for vai <= arr[i+1] {
			if j == i+1 {
				break
			}
			vai1 := arr[i+1]
			arr[i+1] = arr[j]
			arr[j] = vai1
			j--
		}
		if vai > arr[i+1] {
			arr[i] = arr[i+1]
			arr[i+1] = vai
		}
	}
	if i-i > k {
		a.intsort(k, i-1, arr)
	}
	if m > i {
		a.intsort(i, m, arr)
	}
}
func (a *ArrayInt) ToArray() []int {
	s := a.Size()
	m := make([]int, s)
	for i := 0; i < s; i++ {
		m[i] = a.Get(i)
	}
	return m
}
func (a *ArrayInt) Order(arr []int) {
	utils.QuickSort(arr, 0, len(arr)-1)
}
