package main

import (
	"fmt"
	"time"
)

// はじめてのgoroutine

func main() {

	// 関数をgoroutineで実行
	go HelloWorld()

	// 無名関数をgoroutineで実行
	go func() {
		fmt.Println("Hello World! from anonymous function")
	}()

	// main関数をスリープで止めておかないと
	// main関数終了時に問答無用で他のgoroutineも終了する
	time.Sleep(time.Second)

}

func HelloWorld() {
	fmt.Println("Hello World! from function")
}