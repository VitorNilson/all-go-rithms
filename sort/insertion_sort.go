package sort

func InsertionSort(keys []int) []int {
	
	for i := 0; i < len(keys); i++ {
		key := keys[i]
		j := i - 1
		for j > 0 && keys[j] > key {
			keys[j+1] = keys[j]
			j = j - 1
		}
		keys[j+1] = key
	}
	
	return keys
}
