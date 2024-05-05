package main

import "fmt"

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

type Node struct {
	Value int
	Next  *Node
}

func SliceToLinkedList(arr []int) *Node {
	root := &Node{
		Value: arr[0],
	}

	prev := root

	for _, v := range arr[1:] {
		cur := &Node{
			Value: v,
		}

		prev.Next = cur

		prev = cur
	}

	return root
}

func main() {
	arr := []int{1, 2, 3, 4, 5}

	list := SliceToLinkedList(arr)

	cur := list
	for cur != nil {
		fmt.Println(cur.Value)
	}
}
