/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root *TreeNode, p *TreeNode, q *TreeNode) *TreeNode {
	for root != nil {
		if p.Val > root.Val && q.Val > root.Val {
			// right subtree
			root = root.Right
		} else if p.Val < root.Val && q.Val < root.Val {
			// left subtree
			root = root.Left
		} else {
			return root
		}
	}
	return nil
}
