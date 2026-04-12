import "slices"
func combinationSum(nums []int, target int) [][]int {
	resp := [][]int{}
    var dfs func(index int, curr []int, total int)
	dfs = func(index int, curr []int, total int) {
		if total == target {
			resp = append(resp, slices.Clone(curr))
			return
		}
		if index >= len(nums) || total > target {
			return
		}

		curr = append(curr, nums[index])
		dfs(index, curr, total + nums[index])

		curr = curr[:len(curr)-1]
		dfs(index+1, curr, total)
	}
	dfs(0,  []int{}, 0)
	return resp

}
