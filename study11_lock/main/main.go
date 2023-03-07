package main

import (
	"fmt"
	"sync"
	"time"
)

var count = 0
var wg sync.WaitGroup
// 互斥锁
var mutex sync.Mutex

func test (){
	mutex.Lock() //当有协程访问这块资源就会锁上
	count++
	fmt.Println("the count is:",count)
	time.Sleep(time.Millisecond)
	mutex.Unlock()
	wg.Done()
}

func main() {
	for r := 0; r < 20; r++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
}