package binaryTree

import (
	"fmt"
	"testing"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p== nil && q ==nil{
		return true
	}

	if p==nil || q==nil {
		return false
	}

	if p.Val !=q.Val {
		return false
	}

	return isSameTree(p.Right,q.Right)&& isSameTree(p.Left,q.Left)
}

func TestIsSameTree(t *testing.T) {
	a := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	
	b:=TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	c :=isSameTree(&a,&b)

	fmt.Println(c)
}
