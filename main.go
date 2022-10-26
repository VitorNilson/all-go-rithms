package main

import "github.com/vitornilson/all-go-rithms/sort"

func main() {
	arr := []int{7, 5, 8, 4, 1, 3, 6, 2, 9, 0}

	resp := sort.QuickSort(arr, 0, len(arr)-1)

	for i := 0; i < len(resp); i++ {
		print(resp[i], ",")
	}
}
