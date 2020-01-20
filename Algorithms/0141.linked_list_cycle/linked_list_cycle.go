package _141_linked_list_cycle

/*
 * Definition for singly-linked list.
*/


/*
Approach 1: Hash Table

Time complexity : O(n)
Space complexity: O(n)
*/
type ListNode struct {
   Val int
   Next *ListNode
}
//
//type NodesSeen []*ListNode
//
//func hasCycle(head *ListNode) bool {
//	nodesSeen := NodesSeen{}
//	for head != nil{
//		if nodesSeen.contain(head) {
//			return true
//		}
//		nodesSeen = append(nodesSeen,head)
//		head = head.Next
//	}
//	return false
//}
//
//func(n NodesSeen) contain(head *ListNode) bool{
//	for _,val := range n{
//		if *val == *head{
//			return true
//		}
//	}
//	return false
//}

/*
Approach 2: Two Pointers

time complexity is O(N+K), which is O(n)
Space complexity : O(1).
*/
func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}
	slow := head
	fast := head.Next
	//Check whether the fast runner will eventually meet the slow runner or not
	//if not, head has no cycle
	for slow != nil && fast != nil && fast.Next != nil{
		if *slow == *fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}