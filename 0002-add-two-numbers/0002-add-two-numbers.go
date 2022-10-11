/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode
	var first *ListNode
	remaining := 0

	for {
		if l1 == nil && l2 == nil {
			break
		}

		var (
			val1 int
			val2 int
		)

		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		var next ListNode
		next.Val = val1 + val2 + remaining
		remaining = 0

		if next.Val > 9 {
			remaining = next.Val / 10
			next.Val = next.Val % 10
		}

		if result == nil {
			result = &next
			first = &next
		} else {
			result.Next = &next
			result = result.Next
		}

	}
    
    if remaining > 0 {
		var next ListNode
		next.Val = remaining
		result.Next = &next
	}

	return first
}