/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var length = 0
    
    next := head
    for ;next != nil; next = next.Next {
        length++
    }
    
    middle := length/2 + 1
    middleNode := head
    next = head
    for i:=1; i<middle;i++ {
        if i == middle-1 {
            middleNode = next.Next
        }
        
        next = next.Next
    }
    
    return middleNode
}