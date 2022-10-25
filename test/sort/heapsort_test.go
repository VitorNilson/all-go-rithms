package test

import (
	"testing"

	"github.com/vitornilson/all-go-rithms/sort"
)

func TestHeapSort(t *testing.T) {

	arr := []int{70, 1, 3, 2, 500, 9, 0, 10, 14, 8, 7, -2, -1}
	heap := sort.Heap{List: arr}

	sort.HeapSort(&heap)

	
	got := (*[13]int)(heap.List)
	want := &[13]int{-2,-1,0,1,2,3,7,8,9,10,14,70,500}
	
	t.Log(got)
	t.Log(want)

	for i:=0; i < 13; i++{
		if want[i] != got[i] {
			t.Error("got",got[i], "wanted", want[i] )
		}
	}
	
	
}