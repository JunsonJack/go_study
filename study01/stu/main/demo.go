package main

import "fmt"

// 全局变量
var n7 = 77
var account = "zhang"
// 简写
var (
	n8 = 88
	nn = "junson"
)
	
func main() {
	// 局部变量

	// 指定变量类型，并且赋值
	var num1 int = 19
	fmt.Println("num1",num1)

	// 指定变量类型，不赋值  int的默认值是0
	var num2 int 
	fmt.Println("num2",num2)

	// 没有类型，赋值 --可以类型推断
	var num3 = "tom"
	fmt.Println("num3",num3)

	// 省略var 
	name := "zhangsan"
	fmt.Println("name",name)

	/*
	num1 19
	num2 0
	num3 tom
	name zhangsan
	*/

	fmt.Println("---------------------------------------")
	// 声明多个变量
	var n1,n2,n3 int
	fmt.Println("n1",n1)
	fmt.Println("n2",n2)
	fmt.Println("n3",n3)

	var n4,name1,n5 = 10,"jack",7.7
	fmt.Println("n4",n4,"-","name",name1,"-","n5",n5)

	n6,str := 666,"666"
	fmt.Println("n6,str",n6,str)

	/*
		n1 0
		n2 0
		n3 0
		n4 10 - name jack - n5 7.7
	*/
	
	fmt.Println(n7,account)
	fmt.Println("var()",n8,nn)

}