/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
    inOrderPositions := make(map[int]int, 0)
	for ix, n := range inorder {
		inOrderPositions[n] = ix
	}

	root := &TreeNode{Val: preorder[0]}
	mid := inOrderPositions[preorder[0]]
	root.Left = buildTree(preorder[1:mid+1], inorder[:mid])
	root.Right = buildTree(preorder[mid+1:], inorder[mid+1:])

	return root
}
