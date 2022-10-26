package model

type Collection struct {
	HeapSize int
	Array []int
}

func (c Collection) GetKey() int {
	return 0
}

func (c *Collection) SetSize(size int) {
	c.HeapSize = size
}


func (c Collection) Size() int {
	return c.HeapSize
}

func (c Collection) List() []int {
	return c.Array
}