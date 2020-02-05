package select_test

import (
	"fmt"
	"testing"
	"time"
)

func otherTask(){
	fmt.Println("---otherTask---working on something else")
	fmt.Println("---otherTask---立即执行 100 毫秒延时")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("---otherTask---Task is Done")
}
func service() string{
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncService() chan string{
	retCh := make(chan string)

	go func() {
		ret := service()
		fmt.Println("return result")
		retCh <- ret
		fmt.Println("service exited")
	}()

	return retCh
}


func TestSelect(t *testing.T) {
	ret := AsyncService()
	otherTask()
	fmt.Println(<-ret)
}

