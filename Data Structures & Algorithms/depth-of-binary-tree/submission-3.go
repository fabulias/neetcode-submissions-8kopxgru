/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	counter := 0
	for len(queue) > 0 {
		for range len(queue) {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		counter++
	}
	return counter
}
*/

type Pair struct {
	Node *TreeNode
	Depth int
}

// Stack represents a LIFO (Last-In-First-Out) data structure.
type Stack[T any] struct {
	items []T
}

// Push adds an element to the top of the stack.
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element. 
// It returns the zero value and false if the stack is empty.
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	index := len(s.items) - 1
	item := s.items[index]
	
	// Prevent memory leaks: zero out the reference before resizing
	var zero T
	s.items[index] = zero 
	
	s.items = s.items[:index]
	return item, true
}

// IsEmpty returns true if the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

func maxDepth(root *TreeNode) int {
    if root == nil {
		return 0
	}
	stack := &Stack[*Pair]{}
	stack.Push(&Pair{root, 1})
	
	resp := 1
	for !stack.IsEmpty() {
		node, _ := stack.Pop()
		
		if node.Node != nil {
			resp = max(resp, node.Depth)
			stack.Push(&Pair{node.Node.Left, node.Depth+1})
			stack.Push(&Pair{node.Node.Right, node.Depth+1})
		}
	}
	return resp
}
