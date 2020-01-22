package _703_kth_largest_element_in_a_stream

/*  Heap, PriorityQueue */

import (
	"container/heap"
)

// KthLargest object will be instantiated and called as such:
// obj := Constructor(k, nums);
// param_1 := obj.Add(val);
type KthLargest struct {
	k    int
	heap intHeap
}


func Constructor(k int, nums []int) KthLargest {
	h := intHeap(nums)
	//Convert to heap
	heap.Init(&h)

	return KthLargest{
		k:    k,
		heap: h,
	}
}

func (kl *KthLargest) Add(val int) int {
	heap.Push(&kl.heap, val)

	//Pop minimum element until len(h) = k and kthlargest = h[0]
	for len(kl.heap) > kl.k {
		heap.Pop(&kl.heap)
	}

	return kl.heap[0]
}

type intHeap []int

func (h intHeap) Len() int {
	return len(h)
}

func (h intHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h intHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *intHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	res := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return res
}