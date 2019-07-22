package map_test

import "testing"

func TestInitMap(t *testing.T){
	m1:=map[string]int{"1":1,"2":4,"3":9}
	t.Log(m1["2"])
}