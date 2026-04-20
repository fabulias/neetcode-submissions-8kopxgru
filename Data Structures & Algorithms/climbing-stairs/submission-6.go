// top-down approach
/*
func climbStairs(n int) int {
	dp := make([]int, n)
	for ix := range n {
		dp[ix] = -1
	}
    var dfs func(step int) int
	dfs = func(step int) int {
		if step >= n {
			if step == n {
				return 1
			}
			return 0
		}

		if dp[step] != -1 {
			return dp[step]
		}

		dp[step] = dfs(step+1) + dfs(step+2)
		return dp[step]
	}
	
	return dfs(0)
}
*/

// bottom-up approach
func climbStairs(n int) int {
	dp := [2]int{1,1}
	n -= 2
	for n >= 0 {
		dp[0], dp[1] = dp[0]+dp[1], dp[0]
		n--
	}
	return dp[0]
}
