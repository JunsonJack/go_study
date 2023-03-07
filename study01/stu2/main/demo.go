package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// 定义一个整数类型
	var num1 int8 = 23
	fmt.Println("num1",num1)

	var num2 uint8 = 20
	fmt.Println("num2",num2)

	var num3 = 28
	fmt.Printf("num3对应的类型为:%T",num3)

	fmt.Println()
	fmt.Println(unsafe.Sizeof(num3))

	/*
		num1 23
		num2 20
		num3对应的类型为:int
		8
	*/

	/*
		在确保程序正常运行的情况下，尽量选择占据内存空间小的类型
	*/

	// 年龄  注* byte等价于uint8
	var age uint8 = 18 //uint8  一字节 0~255
	var age1 uint8 = 20
	fmt.Println("age:",age)
	fmt.Println(unsafe.Sizeof(age1))



}