func countBits(n int) []int {
	answer := make([]int, n+1)
	for ix := 0; ix <= n; ix++ {
		answer[ix] = bits.OnesCount(uint(ix))
	}
	return answer
}
