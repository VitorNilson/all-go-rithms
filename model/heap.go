package model

type Heap interface {
	Size() int
	List() []int
	SetSize(size int)
	GetKey() int
}
