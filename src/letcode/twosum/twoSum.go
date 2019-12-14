package twosum


import (
	"fmt"
	"testing"
)

func twoSum(nums []int, target int) []int {

	var result []int

	for i:= 0;i < len(nums); i++{
		for j := i + 1;j < len(nums); j++{
			fmt.Println(i)
			fmt.Println(j)
			if nums[i] + nums[j] == target{
				result = append(result,i,j)
			}
		}
		fmt.Println(i)
	}

	return result
}

func TestSum(t *testing.T) {
	nums := []int{3,2,4}

	target := 6

	result :=twoSum(nums,target)

	fmt.Println(result)
}