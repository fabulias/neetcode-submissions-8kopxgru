func climbStairs(n int) int {
	dp := make([]int, n)
    var dfs func(step int) int
	dfs = func(step int) int {
		if step >= n {
			if step == n {

				return 1
			}
			return 0
		}

		if dp[step] != 0 {
			return dp[step]
		}

		dp[step] = dfs(step+1) + dfs(step+2)
		return dp[step]
	}
	
	return dfs(0)
}
