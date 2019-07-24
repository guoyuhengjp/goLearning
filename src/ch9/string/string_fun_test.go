package string

import (
	"strconv"
	"strings"
	"testing"
)

//字符串的拼接
//string.Split(s,"")
//string.Join(s,"")
func TestStringFn(t *testing.T){
	s := "A,B,C"

	parts := strings.Split(s,",")
	for _,part :=range parts{
		t.Log(part)
	}

	t.Log(strings.Join(parts,"-"))
}

func TestConv(t *testing.T){

	// 从数值变换到字符串
	s := strconv.Itoa(10)

	t.Log("str"+s)

	// 从字符串变换到数值
	if i,err:=strconv.Atoi("10");err == nil{
		t.Log(10+i)
	}
}