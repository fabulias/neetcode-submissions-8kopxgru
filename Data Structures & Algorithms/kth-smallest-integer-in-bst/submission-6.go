/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Recursive DFS
func kthSmallest(root *TreeNode, k int) int {
	counter, smallest := k, 0
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
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


// Iterative DFS
/*
func kthSmallest(root *TreeNode, k int) int {
	n := 0
	stack := []*TreeNode{}
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		n++
		if n == k {
			return curr.Val
		}
		curr = curr.Right
	}
	return -1
}
*/