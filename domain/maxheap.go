package domain

// A PriorityQueue implements heap.Interface and holds Items.
type MaxHeap []*Request

func (maxHeap MaxHeap) Len() int { return len(maxHeap) }

func (maxHeap MaxHeap) Less(i, j int) bool {
	return maxHeap[i].desiredFloor > maxHeap[j].desiredFloor
}

func (maxHeap MaxHeap) Swap(i, j int) {
	maxHeap[i], maxHeap[j] = maxHeap[j], maxHeap[i]
}

func (maxHeap *MaxHeap) Push(x any) {
	request := x.(*Request)
	*maxHeap = append(*maxHeap, request)
}

func (maxHeap *MaxHeap) Pop() any {
	old := *maxHeap
	n := len(old)
	request := old[n-1]
	old[n-1] = nil // avoid memory leak
	*maxHeap = old[0 : n-1]
	return request
}
