package animale_test

import (
	"fmt"
	"testing"
)

// interface只需要关注调用接口的定义，不需要关注具体怎么实现
type Animal interface {
	Say(msg string) string
}

//数据构造
type Duck struct {
	Name string
}

type Dog struct {
	Name string
}

//接口中的方法定义，
func(duck *Duck) Say(msg string) string{
	return fmt.Sprintf("duck duck my name is %s%s",duck.Name,msg)
}

func(dog *Dog) Say(world string) string{
	return fmt.Sprintf("wang wang my name is %s%s",dog.Name,world)
}

func TestAnimal(t *testing.T) {
	dog1 := Dog{Name:"kitty"}
	t.Log(dog1.Say("hello"))
}
