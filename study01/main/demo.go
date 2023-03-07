package main

import "fmt"

func main() {
	// 1、变量声明
	var fristname string
	var lastname string
	var username string 
	// 2、变量赋值
	fristname = "zhang"
	lastname = "san"
	// 3、变量使用
	username = fristname + lastname
	fmt.Println("姓氏为：" + fristname + "名字为：" + lastname)
	fmt.Println("全名为：" + username)

}