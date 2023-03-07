package main

import (
	"fmt"
	"time"
)

func main() {
	// 从多个管道接受数据
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	stringChan := make(chan string, 5)
	for i := 0; i < 5; i++ {
		stringChan <- "hello" + fmt.Sprintf("%d",i)
	}

	for {
		select {
		case v:= <-intChan:
			fmt.Println("从intchan读取数据",v)
			time.Sleep(time.Microsecond * 50)
		case v:= <-stringChan:
			fmt.Println("从stringchan读取数据",v)
			time.Sleep(time.Microsecond * 50)
		default:
			fmt.Println("获取完毕")
			return
		}
	}
/* 	
从stringchan读取数据 hello0
从stringchan读取数据 hello1
从intchan读取数据 0
从intchan读取数据 1
从intchan读取数据 2
从stringchan读取数据 hello2
从stringchan读取数据 hello3
从stringchan读取数据 hello4
从intchan读取数据 3
从intchan读取数据 4
从intchan读取数据 5
从intchan读取数据 6
从intchan读取数据 7
从intchan读取数据 8
从intchan读取数据 9
获取完毕 */
}