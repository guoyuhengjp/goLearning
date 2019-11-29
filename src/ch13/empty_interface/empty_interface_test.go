package empty_interface

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}){
	if i,ok :=p.(int); ok{
		fmt.Println("Integer",i)
		return
	}

	if s,ok :=p.(string); ok{
		fmt.Println("String",s)
		return
	}

	fmt.Println("Unknown Type")
}

func SwitchDoSometing(p interface{}){
	switch v:=p.(type) {
	case int:
		fmt.Println("Inter",v)
	case string:
		fmt.Println("string",v)
	default:
		fmt.Println("Unknown Type")
	}
}

func TestSwitch(t *testing.T) {
	SwitchDoSometing(100)
	SwitchDoSometing("hello world")
}


func TestEmptyInterface(t *testing.T) {
	DoSomething(10)
	DoSomething("i loce")
}