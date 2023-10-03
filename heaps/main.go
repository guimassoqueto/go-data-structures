package main

import "fmt"

// Return the index of the parent
func Parent(index int) int {
	return (index - 1) / 2
}

// get the left child index
func Left(index int) int {
	return 2 * index + 1
}

// get the right child index
func Right(index int) int {
	return 2 * index + 2
}

// swap will swap the position of the Parent/child
func (heap *MaxHeap) Swap(index1, index2 int) {
	heap.array[index1], heap.array[index2] = heap.array[index2], heap.array[index1]
}

// Max Heap
type MaxHeap struct{
	array []int
}

// Insert adds an element to the heap
func (heap *MaxHeap) Insert(value int) {
	heap.array = append(heap.array, value)
	heap.MaxHeapifyUp(len(heap.array)-1)
}

// maxHeapifyDown will heapify from top to bottom
func (heap *MaxHeap) MaxHeapifyUp(index int) {
	for heap.array[Parent(index)] < heap.array[index] {
		heap.Swap(Parent(index), index)
		index = Parent(index)
	}
}

// return the largest key, and removes it from the heap
func (heap *MaxHeap) Extract() int {
	if len(heap.array) == 0 {
		fmt.Println("Cannot extract because the length of the array is 0")
		return -1
	}

	extracted := heap.array[0]
	lastElementIndex := len(heap.array) - 1
	heap.array[0] = heap.array[lastElementIndex]
	heap.array = heap.array[:lastElementIndex]

	heap.MaxHeapifyDown(0)
	return extracted
}

// maxHeapifyUp will heapify from bottom to top
func (heap *MaxHeap) MaxHeapifyDown(index int) {
	lastIndex := len(heap.array) - 1
	leftChildIndex, rightChildIndex := Left(index), Right(index)
	childToCompare := 0

	// loop while it has at least one child
	for leftChildIndex <= lastIndex{
		if leftChildIndex == lastIndex { // when left child is the only child
			childToCompare = leftChildIndex
		} else if heap.array[leftChildIndex] > heap.array[rightChildIndex] { // when left child is larger
			childToCompare = leftChildIndex
		} else {// when right child is larger
			childToCompare = rightChildIndex
		}

		// compare array value of the current index to larger child and swap if it is smaller
		if heap.array[index] < heap.array[childToCompare] {
			heap.Swap(index, childToCompare)
			index = childToCompare
			leftChildIndex, rightChildIndex = Left(index), Right(index)
		} else {
			return
		}
	}
}

func main() {
	maxHeap := &MaxHeap{}
	buildHeap := []int{10,20,30,15,9,72,18,96}

	for _, value := range buildHeap {
		maxHeap.Insert(value)
		fmt.Println(*maxHeap)
	}

	for i := 0; i < 5; i++ {
		maxHeap.Extract()
		fmt.Println(*maxHeap)
	}
}
