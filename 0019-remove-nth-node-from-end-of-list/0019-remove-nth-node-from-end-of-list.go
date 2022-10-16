/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    length := 0
    next := head
    
    for ; next != nil; next = next.Next {
        length++
    }
    
    nth := length-n-1
    next = head
    for i:=0;i<nth;i++ {
        next = next.Next
    }
    
    if nth == -1 {
        head = head.Next
    } else if next.Next != nil {
        next.Next = next.Next.Next
    } else {
        head = nil
    }
    
    return head
}