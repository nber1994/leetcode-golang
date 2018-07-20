package kds

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	bit := sum
	carry := 0
	var aBit, bBit, bitSum int
	for nil != l1 || nil != l2 || 0 != carry {
		if nil == l1 {
			aBit = 0
		} else {
			aBit = l1.Val
			l1 = l1.Next
		}
		if nil == l2 {
			bBit = 0
		} else {
			bBit = l2.Val
			l2 = l2.Next
		}
		bitSum = aBit + bBit + carry
		carry = bitSum / 10
		bit.Next = &ListNode{Val: bitSum % 10}
		bit = bit.Next
	}

	return sum.Next
}

func AddTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := &ListNode{}
	bit := sum
	carry := 0
	var bitSum int
	for nil != l1 || nil != l2 || 0 != carry {
		bitSum = carry
		if nil != l1 {
			bitSum += l1.Val
			l1 = l1.Next
		}
		if nil != l2 {
			bitSum += l2.Val
			l2 = l2.Next
		}
		carry = bitSum / 10
		bitSum = bitSum % 10
		bit.Next = &ListNode{Val: bitSum}
		bit = bit.Next
	}

	return sum.Next
}
