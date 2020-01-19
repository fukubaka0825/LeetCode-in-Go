package _141_linked_list_cycle

/*
 * Definition for singly-linked list.
*/
//type ListNode struct {
//     Val int
//     Next *ListNode
//}

type NodesSeen []*ListNode

func hasCycle(head *ListNode) bool {
	nodesSeen := NodesSeen{}
	for head != nil{
		if nodesSeen.contain(head) {
			return true
		}
		nodesSeen = append(nodesSeen,head)
		head = head.Next
	}
	return false
}

func(n NodesSeen) contain(head *ListNode) bool{
	for _,val := range n{
		if *val == *head{
			return true
		}
	}
	return false
}