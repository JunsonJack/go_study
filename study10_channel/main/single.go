package main

func main() {
	// 默认是双向管道
	ch2 := make(chan int, 4)
	ch2 <- 12
	<-ch2
	// 只写管道
	ch3 := make(chan<- int, 2)
	ch3 <- 10
	ch3 <- 12
	// <-ch3 //invalid operation: cannot receive from send-only channel ch3 (variable of type chan<- int)

	// 只读管道
	ch4 := make(<-chan int, 4)
	// ch4 <- 11 ////invalid operation: cannot receive from send-only channel ch4 (variable of type chan<- int)
}