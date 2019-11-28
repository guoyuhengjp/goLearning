package encap_test

import (
	"fmt"
	"testing"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0","Bob",20}

	e1 := Employee{Name:"Mike",Age:30}

	e2 :=new(Employee)
	//e2 是一个指向对象的一个指针

	e2.Id = "2"
	e2.Name = "Rose"
	e2.Age = 22

	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Log(e2.Name)
}


func (e Employee) String() string{
	return fmt.Sprintf("Id:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}



func TestStructOperations(t *testing.T) {
	e:=Employee{"0","Bob",20}
	t.Log(e.String())
}
