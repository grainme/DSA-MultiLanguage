package main

import "fmt"

// for now we'll focus on integers without getting into generics yet.
type DynamicArray struct {
	// Max number of elements of a vector.
	capacity int
	// Current number of elements of a vector.
	size int
	// underlying data structure that holds the elements of a vector.
	data []int
}

func NewDynamicArray(capacity int) *DynamicArray {
	if capacity < 0 {
		panic("Capacity should be positive")
	}

	fmt.Println("Capacity: ", capacity)
	return &DynamicArray{
		capacity: capacity,
		size:     0,
		data:     make([]int, capacity),
	}
}

func (da *DynamicArray) Get(i int) int {
	if i >= da.size || i < 0 {
		panic("Index is out of range")
	}
	return da.data[i]
}

func (da *DynamicArray) Set(i int, n int) {
	if i >= da.size {
		panic("Index is out of range")
	}
	da.data[i] = n
}

func (da *DynamicArray) Print() {
	for i := range da.size {
		fmt.Println(da.data[i])
	}
}

func (da *DynamicArray) Pushback(n int) {
	// expansion strategy
	if da.size >= da.capacity {
		da.resize()
	}
	da.data[da.size] = n
	da.size++
}

func (da *DynamicArray) Popback() int {
	// hysteresis (provide a stable band to avoid repeatedly growing and shrinking)
	if da.size == 0 {
		panic("Array is empty")
	}

	lastElement := da.data[da.size-1]
	da.size--

	// Shrink if size is 1/4 of capacity to avoid thrashing (aka insufficient memory for workload)
	if da.size > 0 && da.size <= da.capacity/4 {
		da.capacity /= 2
		newData := make([]int, da.capacity)
		copy(newData, da.data[:da.size])
		da.data = newData
	}

	return lastElement
}

func (da *DynamicArray) resize() {
	// growth factor of 2 => bad memory usage but better amortized !
	da.capacity *= 2
	newData := make([]int, da.capacity)
	copy(newData, da.data[:da.size])
	da.data = newData
}

func (da *DynamicArray) GetSize() int {
	return da.size
}

func (da *DynamicArray) GetCapacity() int {
	return da.capacity
}

func main() {
}
