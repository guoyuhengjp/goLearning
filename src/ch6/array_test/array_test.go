package array_test

import "testing"

func TestArrayInit(t *testing.T){
	var arr [3]int
	arr1 := [4]int{1,2,3,4}
	arr2 := [...]int{1,3,4,5}
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr2)
}

func TestArrayTravel(t *testing.T){
	arr3:=[...]int{1,3,4,5}
	arr3Sec := arr3[1:4]
	for i:=0;i<len(arr3Sec) ;i++  {
		t.Log(arr3Sec[i])
	}

}