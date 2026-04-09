/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	secondHalf := slow.Next
	slow.Next = nil
	var prev *ListNode
	for secondHalf != nil {
		next := secondHalf.Next
		secondHalf.Next = prev
		prev = secondHalf
		secondHalf = next
	}

	firstPart, secondHalf := head, prev
	for secondHalf != nil {
		tmp := firstPart.Next
		firstPart.Next = secondHalf
		tmp2 := secondHalf.Next
		secondHalf.Next = tmp
		firstPart = tmp
		secondHalf = tmp2
	}
}
