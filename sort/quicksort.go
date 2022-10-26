package sort

func QuickSort(arr []int, init int, end int) []int {

	if init < end {
		arr, mid := Partition(arr, init, end)
		arr = QuickSort(arr, init, mid -1)
		arr = QuickSort(arr, mid +1, end)
	}

	return arr

}

func Partition(arr []int, init int, end int) ([]int, int) {
	x := arr[end]
	i := init

	for j := init; j < end; j++ {

		if arr[j] < x {
			

			arrJ := arr[j]
			arrI := arr[i]

			arr[j] = arrI
			arr[i] = arrJ

			i++

		}

		arrEnd := arr[end]
		arrI := arr[i]

		arr[end] = arrI
		arr[i] = arrEnd

	}

	return arr, i
}
