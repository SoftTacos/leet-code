package findmedianfromdatastream

import (
	"container/heap"
)

// import "github.com/fufuok/utils/generic/heap"

// https://leetcode.com/problems/find-median-from-data-stream/

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h MinHeap) Peek() (v int) {
	// v = h.Pop().(int)
	// h.Push(v)
	return h[0]
}

// embed minheap to reuse code, override Less function to make it a MaxHeap
type MaxHeap struct {
	MinHeap
}

func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j] }

type MedianFinder struct {
	less    *MaxHeap
	greater *MinHeap
}

func Constructor() MedianFinder {
	maxHeap := MaxHeap{}
	heap.Init(&maxHeap)
	minHeap := MinHeap{}
	heap.Init(&minHeap)

	return MedianFinder{
		less:    &maxHeap,
		greater: &minHeap,
	}
}

func (mf *MedianFinder) AddNum(num int) {
	if mf.less.Len() == 0 {
		heap.Push(mf.less, num)
		return
	}

	// if the two heaps are not equal in size:
	if mf.greater.Len() != mf.less.Len() {
		var popped int
		if mf.greater.Len() > mf.less.Len() {
			popped = heap.Pop(mf.greater).(int)
		} else if mf.greater.Len() < mf.less.Len() {
			popped = heap.Pop(mf.less).(int)
		}
		if num < popped {
			heap.Push(mf.less, num)
			heap.Push(mf.greater, popped)
		} else {
			heap.Push(mf.greater, num)
			heap.Push(mf.less, popped)
		}
	} else {
		if mf.less.Peek() > num {
			heap.Push(mf.less, num)
		} else {
			heap.Push(mf.greater, num)
		}
	}
}

func (mf *MedianFinder) FindMedian() (median float64) {
	len := mf.greater.Len() + mf.less.Len()
	if len%2 == 0 {
		median = float64(mf.greater.Peek()+mf.less.Peek()) / 2
	} else {
		if mf.greater.Len() > mf.less.Len() {
			median = float64(mf.greater.Peek())
		} else {
			median = float64(mf.less.Peek())
		}
	}

	return
}
