package main

import (
	"fmt"
)

/**
Given the head of a linked list, remove the nth node from the end of the list and return its head.



Example 1:


Input: head = [1,2,3,4,5], n = 2
Output: [1,2,3,5]
Example 2:

Input: head = [1], n = 1
Output: []
Example 3:

Input: head = [1,2], n = 1
Output: [1]


Constraints:

The number of nodes in the list is sz.
1 <= sz <= 30
0 <= Node.val <= 100
1 <= n <= sz


Follow up: Could you do this in one pass?

*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func SliceToLinkedList(arr []int) *ListNode {
	if len(arr) < 1 {
		return nil
	}

	root := &ListNode{
		Val: arr[0],
	}

	prev := root

	for _, v := range arr[1:] {
		cur := &ListNode{
			Val: v,
		}

		prev.Next = cur

		prev = cur
	}

	return root
}

func LinkedListFromArgs(v ...int) *ListNode {
	return SliceToLinkedList(v)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	list := SliceToLinkedList(arr)

	cur := list
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := getListLength(head)

	count := l - n - 1

	fake := &ListNode{}
	fake.Next = head

	cur := fake

	for ; count >= 0; count-- {
		cur = cur.Next
	}

	cur.Next = cur.Next.Next

	return fake.Next
}

func getListLength(head *ListNode) int {
	if head == nil {
		return 0
	}

	l := 1

	cur := head

	for cur.Next != nil {
		cur = cur.Next
		l++
	}

	return l
}
