package test

import (
	"testing"

	"github.com/vitornilson/all-go-rithms/sort"
)

func Test_GivenQuickSort_WhenThereIsAValidInputArray_ThenResultShouldContainsTheInitialArrayInOrder(t *testing.T) {

	arr := []int{7, 5, 8, 4, 1, 3, 6, 2, 9, 0}

	got := sort.QuickSort(arr, 0, len(arr)-1)
	want := []int{0, 1, 2, 3, 4, 5, 6, 8, 7, 9}

	for i := 0; i < 10; i++ {

		if want[i] != got[i] {
			t.Error("got", got[i], "wanted", want[i])
		}
	}
}

func Test_GivenQuickSort_WhenThereIsAValidInputArrayOf0To9_ThenItemOnIndice5ShouldBeNumber5(t *testing.T) {

	arr := []int{7, 5, 8, 4, 1, 3, 6, 2, 9, 0}

	got := sort.QuickSort(arr, 0, len(arr)-1)
	want := 5

	if got[5] != want {
		t.Error("got", got[5], "wanted", want)
	}
}
