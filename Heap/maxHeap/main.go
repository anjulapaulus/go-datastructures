package main

import "fmt"

type MaxHeap struct {
	heap []int
}

func (mh *MaxHeap) Insert(key int) {
	mh.heap = append(mh.heap, key)
	// arrange heap - bring last element of array to root
	mh.maxHeapifyUp(len(mh.heap) - 1)
}

func (mh *MaxHeap) Extract() int {
	extract := mh.heap[0]
	l := len(mh.heap) - 1

	if len(mh.heap) == 0 {
		fmt.Println("cannot extract because array length is zero")
		return -1
	}
	mh.heap[0] = mh.heap[l]
	mh.heap = mh.heap[:l]

	mh.maxHeapifyDown(0)
	return extract
}

// heapify from top to bottom
func (mh *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(mh.heap) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	//loop while index has atleast one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is only child
			childToCompare = l
		} else if mh.heap[l] > mh.heap[r] { // when left child is larger
			childToCompare = l
		} else { // when right child is larger
			childToCompare = r
		}

		//compare array value of current index to larger child and swap if smaller
		if mh.heap[index] < mh.heap[childToCompare] {
			mh.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// heapify from bottom top
func (mh *MaxHeap) maxHeapifyUp(index int) {
	for mh.heap[parent(index)] < mh.heap[index] {
		mh.swap(parent(index), index)
		index = parent(index)
	}
}

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return i*2 + 1
}

func right(i int) int {
	return i*2 + 2
}

func (mh *MaxHeap) swap(i, j int) {
	mh.heap[i], mh.heap[j] = mh.heap[j], mh.heap[i]
}

func main() {
	m := &MaxHeap{}
	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}
	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
