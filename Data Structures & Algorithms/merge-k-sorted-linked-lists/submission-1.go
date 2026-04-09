/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//import "slices"

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	for len(lists) > 1 {
		mergedLists := make([]*ListNode, 0)

		for ix := 0; ix < len(lists); ix+=2 {
			l1 := lists[ix]
			if ix+1 < len(lists) {
				l2 := lists[ix+1]
				mergedLists = append(mergedLists, mergeSortedLists(l1, l2))
			} else {
				mergedLists = append(mergedLists, l1)
			}
		}
		lists = mergedLists
	}
	return lists[0]
}

func mergeSortedLists(l1, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			curr.Next = l1
			l1 = l1.Next
		} else {
			curr.Next = l2
			l2 = l2.Next
		}
		curr = curr.Next
	}
	if l1 != nil {
		curr.Next = l1
	}
	if l2 != nil {
		curr.Next = l2
	}
	return dummy.Next
}

/*
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
*/