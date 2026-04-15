import "slices"
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		slices.Sort(stones)
		n := len(stones)
		fmt.Println(0, stones)
		last, beforeLast := n-1, n-2
		fmt.Println("2 stones", stones[last], stones[beforeLast])
		if stones[last] == stones[beforeLast] {
			stones = stones[:beforeLast]
			fmt.Println(1, stones)
		} else {
			stones[beforeLast] = stones[last] - stones[beforeLast]
			stones = stones[:n-1]
			fmt.Println(2, stones)
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
