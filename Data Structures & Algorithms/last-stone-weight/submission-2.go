import "slices"
func lastStoneWeight(stones []int) int {
	for len(stones) > 1 {
		slices.Sort(stones)
		n := len(stones)
		last, beforeLast := n-1, n-2
		if stones[last] == stones[beforeLast] {
			stones = stones[:beforeLast]
		} else {
			stones[beforeLast] = stones[last] - stones[beforeLast]
			stones = stones[:n-1]
		}
	}
	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
