package main

import "fmt"

// 管道遍历
func main() {
	ch := make(chan int,5)
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	
	// range 遍历管道  管道没有key
	// 使用range遍历的时候要关闭管道，要不然会报死锁
	close(ch)
	for val := range ch {
		fmt.Println("管道的值：",val)
	}
	/* 
	管道的值： 1
	管道的值： 2
	管道的值： 3
	管道的值： 4
	管道的值： 5 */

	for j := 1; j < 6; j++ {
		fmt.Println("for循环管道：",<-ch)
	}
/* 	
  for循环管道： 1
	for循环管道： 2
	for循环管道： 3
	for循环管道： 4
	for循环管道： 5 */
}