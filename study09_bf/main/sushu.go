package main

import (
	"fmt"
	"time"
)

// 需求：统计1-100有多少个素数？
//方式一：for循环
func rangeNum (){
	start := time.Now().UnixMilli()
	for i := 2; i < 120000; i++ {
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
	end := time.Now().UnixMilli()
	fmt.Println("花费的时间是：",end - start,"ms") //花费的时间是： 2140
}

func main() {
	
	rangeNum()

}