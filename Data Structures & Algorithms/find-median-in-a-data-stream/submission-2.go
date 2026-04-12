type MinHeap []int
func (h MinHeap) Len() int {return len(h)}
func (h MinHeap) Less(i, j int) bool {return h[i] < h[j]}
func (h MinHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0:n-1]
	return x
}

type MaxHeap struct{ MinHeap }
func (h MaxHeap) Less(i, j int) bool {return h.MinHeap[i] > h.MinHeap[j]}

type MedianFinder struct {
	low  *MaxHeap
	high *MinHeap
}

func Constructor() MedianFinder {
    return MedianFinder{
		low: &MaxHeap{MinHeap: make(MinHeap, 0)},
		high: &MinHeap{},
	}
}


func (this *MedianFinder) AddNum(num int)  {
	heap.Push(this.low, num)
	heap.Push(this.high, heap.Pop(this.low))
	
	if this.low.Len() < this.high.Len() {
		heap.Push(this.low, heap.Pop(this.high))
	}
}


func (this *MedianFinder) FindMedian() float64 {
	if this.low.Len() > this.high.Len() {
		return float64(this.low.MinHeap[0])
	}
	return float64(this.low.MinHeap[0] + (*this.high)[0]) / 2
}
