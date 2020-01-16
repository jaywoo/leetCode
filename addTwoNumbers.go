package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

//(2 -> 4 -> 3) + (5 -> 6 -> 4)
func main() {
	l1 := new(ListNode)

	l1.Val = 2
	l1.Next = new(ListNode)
	l1.Next.Val = 3
	l1.Next.Next = new(ListNode)
	l1.Next.Next.Val = 4

	l2 := new(ListNode)
	l2.Val = 5
	l2.Next = new(ListNode)
	l2.Next.Val = 6
	l2.Next.Next = new(ListNode)
	l2.Next.Next.Val = 4

	node := addTwoNumbers1(l1, l2)
	for ln := node; ln != nil; ln = ln.Next {
		fmt.Println(ln.Val)
	}
}

func appendLists(ln *ListNode, newLn *ListNode) {
	if ln.Next == nil {
		ln.Next = newLn
		return
	} else {
		appendLists(ln.Next, newLn)
	}
}

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {
	var pre *ListNode = &ListNode{}
	carry := 0
	curr := pre
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
		}
		if l2 != nil {
			y = l2.Val
		}

		sum := x + y + carry
		val := sum % 10
		carry = sum / 10
		curr.Next = &ListNode{val, nil}
		curr = curr.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry > 0 {
		curr.Next = &ListNode{carry, nil}
	}
	return pre.Next
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	sum := 0
	node := new(ListNode)
	node.Val = -1

	for l1 != nil || l2 != nil {
		tmp_sum := l1.Val + l2.Val + sum
		sum = 0
		sum = tmp_sum / 10
		val := tmp_sum % 10
		if node.Val == -1 {
			node.Val = val
		} else {
			tmp_node := new(ListNode)
			tmp_node.Val = val
			appendLists(node, tmp_node)
		}

		if l1.Next == nil && l2.Next == nil {
			break
		}

		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

	}
	if sum > 0 {
		tmp := new(ListNode)
		tmp.Val = sum
		appendLists(node, tmp)
	}

	return node
}
