package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test (n int) {
	
	if n > 1{
		for i := (n-1)*30000 + 1; i < n*30000; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i % j == 0 {
				flag = false
				break
			}
		}
		if flag {
			// fmt.Println(i,"是素数")
		}
	}
	}
	
	wg.Done()
	
	 
}

func main() {
	start := time.Now().UnixMilli()
	for i := 1; i <=4 ; i++ {
		wg.Add(1)
		test(i)
	}
	end := time.Now().UnixMilli()
	fmt.Println("花费的时间是：",end - start,"ms") //花费的时间是： 2192 ms
}