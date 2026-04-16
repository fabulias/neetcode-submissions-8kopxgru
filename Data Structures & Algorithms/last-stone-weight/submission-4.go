type MaxHeap []int

func (m MaxHeap) Less(i, j int)bool{return m[i] > m[j]}
func (m MaxHeap) Len()int {return len(m)}
func (m MaxHeap) Swap(i,j int){m[i], m[j] = m[j], m[i]}

func (m *MaxHeap) Push(x any) {
	*m = append(*m, x.(int))
}

func (m *MaxHeap) Pop() any {
	old := *m
	n := len(old)
	x := old[n-1]
	*m = old[:n-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := MaxHeap(stones)
	heap.Init(&h)
	fmt.Println(h)
	for h.Len() > 1 {
		x := heap.Pop(&h).(int)
		y := heap.Pop(&h).(int)
		if x > y {
			heap.Push(&h, x-y)
		} 
	}
	if h.Len() == 1 {
		return h[0]
	}
	return 0
}
