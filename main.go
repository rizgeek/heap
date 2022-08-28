package main

import "fmt"

// MaxHeap struct has a slice that hold the array
type MaxHeap struct {
	array []int
}

// insert add a element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// extract return the largest key, and move it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	// when the array is empty
	if len(h.array) == 0 {
		fmt.Println("Cannot extract because array length is 0")
		return -1
	}

	// take out the last index and put it in the root
	h.array[0] = h.array[l]
	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.array[parent(index)] < h.array[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {

	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	// loop while index has at least one child
	for l <= lastIndex {

		if l == lastIndex { // when left is the only child
			childToCompare = l
		} else if h.array[l] > h.array[r] { // whe le child is large
			childToCompare = l
		} else { // when right child is large
			childToCompare = r
		}

		// compare array value current index to larger child and swap if smaller
		if h.array[r] < h.array[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}

	}
}

// get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get left child index
func left(i int) int {
	return 2*i + 1
}

// get right child index
func right(i int) int {
	return 2*i + 2
}

// swap key in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.array[i1], h.array[i2] = h.array[i2], h.array[i1]
}

func main() {

	m := &MaxHeap{}
	buildHelp := []int{9, 5, 1, 7, 20, 34, 50, 8, 14, 16, 48}

	for _, v := range buildHelp {
		m.Insert(v)
		fmt.Println(m)
	}

	for i := 0; i < 7; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
