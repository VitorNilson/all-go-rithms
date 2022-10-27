package sort

import "github.com/vitornilson/all-go-rithms/model"

func CountingSort(arr []int) []int {

	heap := model.Collection{Array: arr}
	BuildMaxHeap(&heap)

	max := heap.Array[0] + 1

	var tempArr []int = make([]int, max)
	var output []int = make([]int, len(arr))

	for i := 0; i < max; i++ {
		tempArr[i] = 0
	}

	for j := 0; j < len(arr); j++ {
		tempArr[arr[j]] += 1
	}

	for i := 1; i < max; i++ {
		tempArr[i] += tempArr[i-1]
	}

	for i := 0; i < len(arr); i++ {
		output[tempArr[arr[i]]-1] = arr[i]
		tempArr[arr[i]] -= 1
	}

	return output
}
