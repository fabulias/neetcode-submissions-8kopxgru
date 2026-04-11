/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
    return valid(root, -1001, 1001)
}

func valid(node *TreeNode, left, right int) bool {
	if node == nil {
		return true
	}
	if !(left < node.Val && node.Val < right) {
		return false
	}
	return valid(node.Left, left, node.Val) && valid(node.Right, node.Val, right)
}
