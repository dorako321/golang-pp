package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	//go func() {
	//	ch1<-5
	//}()

	for {
		select {
			case value := <- ch1:
				fmt.Println(value)

			case value := <- ch2:
				fmt.Println(value)

			case value := <- ch3:
				fmt.Println(value)

			default:
				//fmt.Println("nothing to do !")
		}
	}

}