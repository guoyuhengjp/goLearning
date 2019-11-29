package polymorphsim

//Go语言的多态

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {

}

func (p *GoProgrammer) WriteHelloWorld() Code{
	return "fmt.Println(\"hello world\")"
}

type JavaProgrammer struct {

}

func (p *JavaProgrammer) WriteHelloWorld() Code{
	return "System.out.Println(\"hello world\")"
}

func writeFirstProgrammer(p Programmer){
	fmt.Printf("%T %v\n",p,p.WriteHelloWorld())
}

func TestPolymorphsim(t *testing.T) {
	goProg :=&GoProgrammer{}

	JavaProgrammer := new(JavaProgrammer)

	writeFirstProgrammer(goProg)

	writeFirstProgrammer(JavaProgrammer)
}