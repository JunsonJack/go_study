package main

import (
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup
// 写数据
func writeData (ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i
		fmt.Printf("写入数据%v成功\n", i)
		time.Sleep(time.Microsecond * 100)
	}
	close(ch)
	wg.Done()
}

// 读数据
func readData (ch chan int){
	for v := range ch {
		fmt.Printf("读取数据%v成功\n",v)
		time.Sleep(time.Microsecond * 50)
	}
	wg.Done()
}

func main() {
	var ch = make(chan int,10)
	wg.Add(1)
	go writeData(ch)
	wg.Add(1)
	go readData(ch)

	wg.Wait()
	fmt.Println("执行完毕...")

	// 管道是线程安全的，会等待写入后才读取数据
/* 	
写入数据1成功
读取数据1成功
写入数据2成功
写入数据3成功
读取数据2成功
写入数据4成功
读取数据3成功
写入数据5成功
读取数据4成功
读取数据5成功
写入数据6成功
写入数据7成功
读取数据6成功
写入数据8成功
读取数据7成功
写入数据9成功
读取数据8成功
写入数据10成功
读取数据9成功
读取数据10成功
执行完毕... */
}