/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

/* Naive solution
func isBalanced(root *TreeNode) bool {
    if root == nil {
		return true
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)

	if math.Abs(float64(left - right)) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return 1 + max(getDepth(node.Left), getDepth(node.Right))
}

*/

func isBalanced(root *TreeNode) bool {
	var dfs func(node *TreeNode) (bool, int)
	dfs = func(node *TreeNode) (bool, int) {
		if node == nil {
			return true, 0
		}
		leftBalanced, leftHeight := dfs(node.Left)
		rightBalanced, rightHeight := dfs(node.Right)

		balanced := leftBalanced && rightBalanced && abs(leftHeight - rightHeight) <= 1

		return balanced, 1 + max(leftHeight, rightHeight)
	}
	if isBalanced, _ := dfs(root); isBalanced {
		return true
	}
	return false
}

func abs(x int ) int {
	if x < 0 {
		return -1*x
	}
	return x
}