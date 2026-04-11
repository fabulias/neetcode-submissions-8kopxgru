/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
    queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	
	res := [][]int{}
	
	for len(queue) > 0 {
		levelSize := len(queue)
		currentLevelValues := []int{}

		for ix := 0; ix < levelSize; ix++ {
			el := queue[0]
			queue = queue[1:]

			currentLevelValues = append(currentLevelValues, el.Val)
			if el.Left != nil {
				queue = append(queue, el.Left)
			}	
			if el.Right != nil {
				queue = append(queue, el.Right)
			}	
		}
		res = append(res, currentLevelValues)
	}
	return res
}
