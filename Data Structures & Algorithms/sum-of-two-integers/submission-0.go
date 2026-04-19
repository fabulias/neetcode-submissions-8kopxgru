func getSum(a int, b int) int {
	// xor: ^
	for b != 0 {
		tmp := a
		a = a ^ b
		b = (tmp & b) << 1
	}
	return a
}
