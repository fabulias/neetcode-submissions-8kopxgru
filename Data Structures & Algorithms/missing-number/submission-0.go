func missingNumber(nums []int) int {
	xorr := len(nums)
	for ix := 0; ix < len(nums); ix++ {
		xorr ^= nums[ix]
		xorr ^= ix
	}
	return xorr
}
