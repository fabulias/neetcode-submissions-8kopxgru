/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxPathSum(root *TreeNode) int {
	res := root.Val
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		leftMax := dfs(root.Left)
		rightMax := dfs(root.Right)
		leftMax = max(leftMax, 0)
		rightMax = max(rightMax, 0)

		// Get the max without split
		res = max(res, root.Val + leftMax + rightMax)

		// Return spliting the result
		return root.Val + max(leftMax, rightMax)
	}
	dfs(root)
	return res
}
