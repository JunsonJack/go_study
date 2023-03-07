package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 获取当前计算机上的CPU个数
	cpuNum := runtime.NumCPU()
	fmt.Println("cpunum=",cpuNum)
	//cpunum= 8

	// 自己设置使用多少个CPU
	runtime.GOMAXPROCS(cpuNum - 1)
	fmt.Println("ok")
}