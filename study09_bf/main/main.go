package main

import (
	"fmt"
	"sync"
	"time"
)

// 等待协程执行完毕
var wg sync.WaitGroup

func rundemo (){
	for i := 1; i <= 20; i++ {
		fmt.Printf("协程打印了%v\n",i)
		time.Sleep(time.Microsecond * 100)
	}
	wg.Done() //2、协程计数器减一
}
func main() {
	wg.Add(1) //1、协程计数器加一
	// goroutine 开启协程
	go rundemo()
	for i := 1; i <= 10; i++ {
		fmt.Printf("普通打印%v\n", i)
		time.Sleep(time.Microsecond * 100)	
	}
	wg.Wait() //3、等待协程执行结束

}