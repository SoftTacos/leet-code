package main

import "fmt"

func main() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		output  = &ListNode{}
		current = output
		v1, v2  int
		carry   = 0
	)

	// if next is not nil, get the value and add it
	// if next is nil, just use 0
	for {
		v1 = 0
		v2 = 0
		if l1 != nil {
			v1 = l1.Val
		}
		if l2 != nil {
			v2 = l2.Val
		}
		sum := (v1 + v2 + carry)
		current.Val = sum % 10
		carry = sum / 10

		// advance the nodes
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}

		// if both are nil, return
		if l1 == nil && l2 == nil {
			if carry > 0 {
				current.Next = &ListNode{
					Val: 1,
				}
			}
			break
		}

		current.Next = &ListNode{}
		current = current.Next
	}
	for lt := output; lt != nil; lt = lt.Next {
		fmt.Print(lt.Val)
	}
	return output
}
