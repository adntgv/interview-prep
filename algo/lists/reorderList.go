package main

import (
	"fmt"
)

/**
You are given the head of a singly linked-list. The list can be represented as:

L0 → L1 → … → Ln - 1 → Ln
Reorder the list to be on the following form:

L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
You may not modify the values in the list's nodes. Only nodes themselves may be changed.



Example 1:


Input: head = [1,2,3,4]
Output: [1,4,2,3]
Example 2:


Input: head = [1,2,3,4,5]
Output: [1,5,2,4,3]


Constraints:

The number of nodes in the list is in the range [1, 5 * 104].
1 <= Node.val <= 1000

*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	preCenter := getPreMid(head)
	reversed := reverseList(preCenter.Next)
	// preCenter.Next = nil

	p1 := head
	p2 := reversed
	tmp := head

	for p2 != nil {
		p1n := p1.Next
		p1.Next = p2
		p1 = p2
		p2 = p1n
	}

	head = tmp
}

func getPreMid(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}

func linkedListToString(head *ListNode) string {
	res := ""
	for head != nil {
		res += fmt.Sprintf("%v -> ", head.Val)
		head = head.Next
	}

	return res
}

func listsEq(first, second *ListNode) bool {
	for first != nil && second != nil {
		fmt.Println(first.Val, second.Val)
		if first.Val != second.Val {
			return false
		}
		first = first.Next
		second = second.Next
	}
	fmt.Println("exot", first == nil && second == nil)

	return first == nil && second == nil
}
