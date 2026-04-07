func binarySearch(nums []int, left, right int) int {
	res := nums[0]
	for left <= right {
		if nums[left] < nums[right] {
			res = min(res, nums[left])
			break
		}
		mid := left + (right - left) / 2
		res = min(res, nums[mid])
		if nums[mid] >= nums[left] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
	
func findMin(nums []int) int {
	/*
	minValue := nums[0]
	for _, n := range nums {
		if n < minValue {
			minValue = n
		}
	}
	fmt.Println("minValue", minValue)
	*/

	val := binarySearch(nums, 0, len(nums)-1)
	return val

	//return minValue
}