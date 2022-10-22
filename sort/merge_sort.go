package sort

func Merge(left []int, right []int) []int {

	final := []int{}

	i := 0
	j := 0

	for i < len(left) && j < len(right) {

		if left[i] < right[j] {
			final = append(final, left[i])
			i++
		} else {
			final = append(final, right[j])
			j++
		}
	}

	for ;i < len(left); i++ {
		final = append(final, left[i])
	}

	for ;j < len(right);j++ {
		final = append(final, right[j])
	}

	return final

}

func MergeSort(array []int) []int {

	if len(array) < 2 {
		return array
	}

	left := MergeSort(array[:len(array)/2])
	right := MergeSort(array[len(array)/2:])

	return Merge(left, right)
}
