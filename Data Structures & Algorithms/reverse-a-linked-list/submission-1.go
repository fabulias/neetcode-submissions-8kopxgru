/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// iterative approach
/*
func reverseList(head *ListNode) *ListNode {
 	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
*/
// recursive approach
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}	
	newHead := head
	if head.Next != nil {
		newHead = reverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil
	return newHead
}