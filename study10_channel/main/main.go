package main

import "fmt"

func main() {
	//1、创建channel
	ch := make(chan int, 3)

	// 2、给管道里面存储数据

	ch <- 3
	ch <- 21
	ch <- 11

	a := <-ch
	fmt.Printf("a: %v\n", a)  //a: 3
	<-ch
	c := <-ch
	fmt.Printf("c: %v\n", c) //c: 11
	ch <- 55
	ch <- 66
	// 3、管道的类型(引用数据类型)
	ch1 := make(chan int,4)
	ch1 <- 11
	ch1 <- 21
	ch1 <- 31

	ch2 := ch1

	ch2 <- 22

	<-ch1
	<-ch1
	<-ch1

	d:=<-ch1
	fmt.Printf("d: %v\n", d) //d: 22 说明ch2改变了ch1属于是引用数据类型

	// 4、管道的容量和长度
	fmt.Printf("值：%v 容量：%v 长度：%v", ch,cap(ch),len(ch)) //值：0x1180e320 容量：3 长度：2
	// 管道的值是一个地址、容量就是定义的时候的值，长度是未被取出的长度

	// 5、管道阻塞
	ch4 := make(chan int,1)
	ch4 <- 22
	// ch4 <- 23
	//fatal error: all goroutines are asleep - deadlock! 管道长度为1，但是放进去两个数据，造成阻塞

	ch5 := make(chan string,2)
	ch5 <-"数据1"
	ch5 <- "数据2"
	<-ch5
	<-ch5
	// <-ch5
	//fatal error: all goroutines are asleep - deadlock! 当管道里面没有值了，再继续取也属于阻塞
}
