func reverseBits(n int) int {
	reverseN := 0
	for ix := range 32 {
		bit := (n >> ix) & 1
		reverseN = reverseN | (bit << (31 - ix))
	}
	return reverseN
}
