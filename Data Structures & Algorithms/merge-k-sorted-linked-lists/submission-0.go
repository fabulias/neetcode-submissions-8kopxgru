/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "slices"


func mergeKLists(lists []*ListNode) *ListNode {
	iters := make([]*ListNode, 0)
    for _, list := range lists {
		if list != nil {
			iters = append(iters, list)
		}
	}

	dummy := &ListNode{}
	curr := dummy
	for len(iters) > 0 {
		minIndex := 0
		for ix, iter := range iters {
			if iter.Val < iters[minIndex].Val {
				minIndex = ix
			}
		}
		
		curr.Next = &ListNode{Val:iters[minIndex].Val}
		curr = curr.Next

		iters[minIndex] = iters[minIndex].Next
		
		if iters[minIndex] == nil {
			iters = slices.Delete(iters, minIndex, minIndex+1)
		}
	}

	return dummy.Next
}
