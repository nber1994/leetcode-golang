package main

import "fmt"

type ListNode struct {
    Val int
    Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    res := &ListNode{}
    tmp := res
    carry := 0
    var val, sum int
    for nil != l1 || nil != l2 || 0 != carry {
        sum = carry
        if nil != l1 {
            sum += l1.Val
            l1 = l1.Next
        }
        if nil != l2 {
            sum += l2.Val
            l2 = l2.Next
        }

        val = sum % 10
        carry = sum / 10

        tmp.Next = &ListNode{Val: val}
        tmp = tmp.Next;
    }
    return res.Next;
}

// [2,4,3]
// [5,6,4]
func main() {
    l1 := &ListNode{Val: 2}
    l1.Next = &ListNode{Val: 4}
    l1.Next.Next = &ListNode{Val: 3}

    l2 := &ListNode{Val: 5}
    l2.Next = &ListNode{Val: 6}
    l2.Next.Next = &ListNode{Val: 4}

    res := addTwoNumbers(l1, l2);
    for nil != res {
        fmt.Println(res.Val)
        res = res.Next
    }
}
