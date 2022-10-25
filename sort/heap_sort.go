package sort


type Heap struct {
	HeapSize int
	List     []int
}

func HeapSort(heap *Heap) {

	BuildMaxHeap(heap)

	for i := len(heap.List) - 1; i >= 1; i-- {
		firstValue := heap.List[0]
		indexValue := heap.List[i]

		heap.List[i] = firstValue
		heap.List[0] = indexValue

		heap.HeapSize -= 1

		MaxHeapify(heap, 0)

	}

}

func BuildMaxHeap(heap *Heap) {
	heap.HeapSize = len(heap.List) - 1

	mid := (heap.HeapSize / 2)
	for i := mid; i >= 0; i-- {
		MaxHeapify(heap, i)
	}

}

func MaxHeapify(heap *Heap, index int) {
	left := Left(index)
	right := Right(index)
	
	leftSmallerThanHeapSize := left <= heap.HeapSize
	rightSmallerThanHeapSize := right <= heap.HeapSize


	var largest int

	if leftSmallerThanHeapSize && heap.List[left] > heap.List[index] {
		largest = left
	} else {
		largest = index
	}

	if rightSmallerThanHeapSize && heap.List[right] > heap.List[largest] {
		largest = right
	}

	if largest != index {

		indexValue := heap.List[index]
		largestValue := heap.List[largest]

		heap.List[index] = largestValue
		heap.List[largest] = indexValue
		MaxHeapify(heap, largest)
	}

}

func Left(index int) int {
	return 2 * index +1
}

func Right(index int) int {
	return (2 * index) + 2
}

func Parent(index int) int {
	return index / 2
}


