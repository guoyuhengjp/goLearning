package twosum

import (
	"fmt"
	"testing"
)

func twoSum1(nums []int, target int)[]int{
	lookup := make(map[int]int)

	for i, v := range nums {
		j, ok := lookup[-v]
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


func reverse(x int) int {

	var result int = 0

	for ; x !=0; x/=10  {
		result = result*10 + x % 10

		if result > 2147483647|| result < -2147483648{
			return 0
		}
	}
	return result
}

func TestReverse(t *testing.T) {
	var x int = 1534236469

	result := reverse(x)

	println(result)
}