/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
	counter, smallest := k, 0
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		if counter == 0 {
			return
		}
		counter--
		if counter == 0 {
			smallest = node.Val
			return
		}
		inOrder(node.Right)
	}
	inOrder(root)
	return smallest
}