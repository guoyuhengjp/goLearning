package channel_close_test

import (
	"fmt"
	"sync"
	"testing"
)

func dataProducer(ch chan int,wg *sync.WaitGroup){
	go func() {
		for i :=0; i<10; i++ {
			ch <- i
		}
		close(ch) //关闭这个通道
		wg.Done()
	}()
}

func dataReceiver(ch chan int ,wg *sync.WaitGroup){
	go func() {
		for{
			if date,ok:=<-ch;ok{
				fmt.Println(date)
			}else {
				break
			}
		}
		wg.Done()
	}()
}

func TestCloseChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch:= make(chan int)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}
