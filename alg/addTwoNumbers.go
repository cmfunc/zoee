package alg

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表遍历
	var sum int
	index := 1
	index2:=1
	for l1.Next!=nil{
		sum = sum + l1.Val*index	
		index = index *10
	}

	for l2.Next!=nil{
		sum = sum + l1.Val*index	
		index2 = index2 *10
	}

	return nil
}
