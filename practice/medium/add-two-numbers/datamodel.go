package add_two_numbers

import "fmt"

// ListNode definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func (n *ListNode) GetNumber() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%d%s", n.Val, n.Next.GetNumber())
}
