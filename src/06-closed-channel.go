package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})

	for i := 0; i < 10; i++ {
		go func() {
			<-ch
			fmt.Println("do something..", i)
		}()
	}

	fmt.Println("waiting 5 sec")
	time.Sleep(5 * time.Second)
	fmt.Println("channel will be closed")
	close(ch)

	// スリープしとかないと他のgoroutineの処理を行う前に終了してしまうので
	time.Sleep(2 * time.Second)

}