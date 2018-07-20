/**
* Definition for singly-linked list.
* type ListNode struct {
	*     Val int
	*     Next *ListNode
	* }
	*/
	func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
		forward := 0
		now := &ListNode{}
		first := now
		i := 0
		for {
			//当两个都已经为nil
			if l1 == nil && l2 == nil {
				if forward > 0 {
					//当存在进位时 
					now.Next = &ListNode{Val:forward}
				} 
				break
			}
			tmpNode := &ListNode{}
			if l1 == nil {
				tmpNode.Val = (l2.Val + forward) % 10
				forward = (l2.Val + forward) / 10
			} else if l2 == nil {
				tmpNode.Val = (l1.Val + forward) % 10
				forward = (l1.Val + forward) / 10
			} else {
				sum := l1.Val + l2.Val
				tmpNode.Val = (sum + forward) % 10
				forward = (sum + forward) / 10
			}
			if i == 0 {
				now = tmpNode
				first = now
			} else {
				now.Next = tmpNode
				now = tmpNode
			}
			i++
			if l1 != nil {
				l1 = l1.Next 
			}
			if l2 != nil {
				l2 = l2.Next 
			}
		}
		return first
	}
