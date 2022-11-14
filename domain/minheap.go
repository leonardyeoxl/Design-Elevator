package domain

// A PriorityQueue implements heap.Interface and holds Items.
type MinHeap []*Request

func (minHeap MinHeap) Len() int { return len(minHeap) }

func (minHeap MinHeap) Less(i, j int) bool {
	return minHeap[i].desiredFloor < minHeap[j].desiredFloor
}

func (minHeap MinHeap) Swap(i, j int) {
	minHeap[i], minHeap[j] = minHeap[j], minHeap[i]
}

func (minHeap *MinHeap) Push(x any) {
	request := x.(*Request)
	*minHeap = append(*minHeap, request)
}

func (minHeap *MinHeap) Pop() any {
	old := *minHeap
	n := len(old)
	request := old[n-1]
	old[n-1] = nil // avoid memory leak
	*minHeap = old[0 : n-1]
	return request
}
