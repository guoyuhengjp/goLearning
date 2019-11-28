package extension_test

import (
	"fmt"
	"testing"
)

type pet struct {

}

func (p *pet) Speak22(){
		fmt.Print("...")
}

func (p *pet) SpeakTo(host string){
	fmt.Println(" ",host)
}

type Dog struct {
	pet
}

func (d *Dog) Speak(){
	fmt.Print("wang!")
}

func TestDog(t *testing.T) {
	dog :=new(Dog)

	dog.SpeakTo("chao")

	dog.Speak22()
}


