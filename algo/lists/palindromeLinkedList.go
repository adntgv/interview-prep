package main

/**
Given the head of a singly linked list, return true if it is a
palindrome
 or false otherwise.



Example 1:


Input: head = [1,2,2,1]
Output: true
Example 2:


Input: head = [1,2]
Output: false


Constraints:

The number of nodes in the list is in the range [1, 105].
0 <= Node.val <= 9


Follow up: Could you do it in O(n) time and O(1) space?
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	center := middleNode(head)
	reversed := reverseList(center)

	for head != nil && reversed != nil {
		if reversed.Val != head.Val {
			return false
		}

		reversed = reversed.Next
		head = head.Next
	}

	return true
}
