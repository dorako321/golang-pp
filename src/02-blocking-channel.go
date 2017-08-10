package main

//func main() {
//	//fatal error: all goroutines are asleep - deadlock!
//	ch := make(chan struct{})
//	<-ch
//}


//func main() {
//	//fatal error: all goroutines are asleep - deadlock!
//	ch := make(chan struct{})
//	ch<-struct{}{}
//}


func main() {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	<-ch
}