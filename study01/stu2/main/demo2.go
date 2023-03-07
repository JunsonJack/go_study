package main

import (
	"fmt"
)
func main() {
	// 浮点型
	var num float32 = 0.66
	/*
		float32 4字节
		float64 8字节
	*/
	fmt.Println("num",num) // num 0.66
	
	// 科学记数法
	var num1 float32 = 314E-2
	fmt.Println("num1",num1) // num1 3.14
	
	var num2 float32 = 314e+2
	fmt.Println("num2",num2)  //num2 31400

	var num3 float64 = 314e+2
	fmt.Println("num3",num3)  //num3 31400



}