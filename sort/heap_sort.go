package sort

import "github.com/vitornilson/all-go-rithms/model"



func HeapSort(heap model.Heap) {

	BuildMaxHeap(heap)

	for i := len(heap.List()) - 1; i >= 1; i-- {

		firstValue := heap.List()[0]
		indexValue := heap.List()[i]

		heap.List()[i] = firstValue
		heap.List()[0] = indexValue

		heap.SetSize(heap.Size() - 1)

		MaxHeapify(heap, 0)

	}

}

func BuildMaxHeap(heap model.Heap) {
	heap.SetSize(len(heap.List()) - 1)

	mid := (heap.Size() / 2)
	for i := mid; i >= 0; i-- {
		MaxHeapify(heap, i)
	}

}

func MaxHeapify(heap model.Heap, index int) {
	left := Left(index)
	right := Right(index)
	
	leftSmallerThanSize := left <= heap.Size()
	rightSmallerThanSize := right <= heap.Size()


	var largest int

	if leftSmallerThanSize && heap.List()[left] > heap.List()[index] {
		largest = left
	} else {
		largest = index
	}

	if rightSmallerThanSize && heap.List()[right] > heap.List()[largest] {
		largest = right
	}

	if largest != index {

		indexValue := heap.List()[index]
		largestValue := heap.List()[largest]

		heap.List()[index] = largestValue
		heap.List()[largest] = indexValue
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


