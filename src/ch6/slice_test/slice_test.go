package slice_test

import "testing"

func TestSliceInit(t *testing.T){
	var s0 []int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1)
	t.Log(len(s0),cap(s0))

	var s1 =[...]int{1,2,3,4}
	t.Log(s1)
}

func TestSliceShareMemory(t *testing.T){
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep",
		"Oct", "Nov", "Dec"}
	t.Log(year)
}
