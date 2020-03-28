package leetCodeArray

import (
	"fmt"
	"testing"
)

func pivotIndex(nums []int) int {
	for i ,_:= range nums{
		var result = 0
		var temp  = 0

		for k := 0;k < i ;k++  {
			result += nums[k]
		}

		for j := i+1; j< len(nums);j++  {
			temp += nums[j]
		}

		if temp == result {
			return i
		}
	}

	return -1
}

type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val{
		head.Next = head.Next.Next
	}
	return head
}
func ShowNode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func TestDeleteDuplicates(t *testing.T) {

	var head = new(ListNode)
	head.Val = 1
	var node1 = new(ListNode)
	node1.Val = 1
	head.Next = node1
	var node2 = new(ListNode)
	node2.Val = 2
	node1.Next = node2
	head = deleteDuplicates(head)
	ShowNode(head)

}


func TestPivotIndex(t *testing.T) {
	var hoge = []int{1, 7, 3, 6, 5, 6}
	result :=pivotIndex(hoge)
	fmt.Println(result)
}

