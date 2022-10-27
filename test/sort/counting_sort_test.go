package test

import (
	"testing"

	"github.com/vitornilson/all-go-rithms/sort"
)

func Test_GivenCountingSort_WhenThereIsAValidInputArray_ThenResultShouldContainsTheInitialArrayInOrder(t *testing.T) {
	arr := []int{7, 5, 8, 4, 1, 3, 70, 2, 9, 0}
	
	got := sort.CountingSort(arr)
	want := []int{0,1,2,3,4,5,7,8,9,70}

	for i:=0; i < 10; i++{
		if want[i] != got[i] {
			t.Error("got",got[i], "wanted", want[i] )
		}
	}
}

func Test_GivenCountingSort_WhenThereIsAValidInputArrayOf0To9_ThenItemOnIndice5ShouldBeNumber5(t *testing.T) {
	arr := []int{7, 5, 8, 4, 1, 3, 6, 2, 9, 0}
	
	got := sort.CountingSort(arr)
	want := 5

	if got[5] != want {
		t.Error("got", got[5], "wanted", want)
	}
}

func Test_GivenCountingSort_WhenThereIsAValidInputArray_ThenResponseLenghtShouldBeSameAsInput(t *testing.T) {
	arr := []int{7, 5, 8, 4, 1, 3, 6, 2, 9, 0}
	
	got := len(sort.CountingSort(arr))
	want := len(arr)

	if got != want {
		t.Error("got", got, "wanted", want)
	}
}