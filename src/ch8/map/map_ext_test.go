package map_ext

import "testing"

//map的函数带入的方法
func TestMapWithFunValue(t *testing.T) {
	// map[数据类型]数据类型{数据类型:数据类型}
	//数据类型以外 也可以接函数
	m:= map[int]func(op int)int{}
	m[1] = func(op int)int{
		return op
	}
	m[2] = func(op int)int{
		return op*op
	}
	m[3] = func(op int)int{
		return op*op*op
	}
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T){
	mySet := map[int]bool{}
	mySet[1] = true

	n :=3
	if mySet[n]{
		t.Logf("%d is exxisting",n)
	}else {
		t.Logf("%d is not exxisting",n)
	}

	mySet[3] = true
	t.Log(len(mySet))
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
}
