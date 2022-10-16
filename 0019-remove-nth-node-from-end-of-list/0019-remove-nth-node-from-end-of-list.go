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
    behind := head
    for ; next != nil; next = next.Next {
        if length > n {
            behind = behind.Next
        }
        
        length++
    }
    
    nth := length-n-1
    if nth == -1 {
        head = head.Next
    } else if behind.Next != nil {
        behind.Next = behind.Next.Next
    } else {
        head = nil
    }
    
    return head
}