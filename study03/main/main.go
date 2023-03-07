package main

import "fmt"

func main() {
	test()
	fmt.Println("上面的触发操作执行成功")
	fmt.Println("正常执行下面的操作.....")
}
func test (){
	// 使用defer+recover来捕获错误 ：defer加上后面的匿名函数的调用
	defer func(){
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获")
			fmt.Println("err是：",err)
		}
	}() //加上括号可以自动调用
	num1 := 10
	num2 := 0
	result := num1/num2
	fmt.Println("res",result)
}

/*
	错误已经捕获
	err是： runtime error: integer divide by zero
	上面的触发操作执行成功
	正常执行下面的操作.....
*/