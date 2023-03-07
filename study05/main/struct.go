package main

import "fmt"

// 定义一个结构体
type Person struct {
	name , uname , sex string
	age int

}
func main() {
	var person1 Person

	person1.name = "张三"
	person1.uname = "帅哥"
	person1.sex = "男"
	person1.age = 18
	
	fmt.Printf("Person:%#v",person1)
	// Person:main.Person{name:"张三", uname:"帅哥", sex:"男", age:18}
	fmt.Printf("%v的昵称是%s\n",person1.name,person1.uname) //张三的昵称是帅哥
}