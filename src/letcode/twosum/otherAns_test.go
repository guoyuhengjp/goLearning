package twosum

import (
	"fmt"
	"testing"
)

func twoSum1(nums []int, target int)[]int{
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[-v]
		fmt.Println(ok)
		lookup[v - target] = i
		if ok {
			return []int{j, i}
		}
	}
	return []int{}
}


func addTwo3(nums []int, target int) []int{
	lookup := make(map[int]int)

	for i,v := range nums{
		j, ok := lookup[-v]
		fmt.Println(ok)
		lookup[v - target] = i
		if ok {
			return []int{j, i}
		}
	}

	return []int{}
}


func TestTwo(t *testing.T) {

	nums := []int{3,2,4}

	target := 6

	result :=twoSum1(nums,target)

	fmt.Println(result)
}