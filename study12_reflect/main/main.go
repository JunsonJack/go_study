package main

import (
	"fmt"
	"reflect"
)

type myInt int

type Person struct {
	Name string
	Age int
}

func reflectFn(x interface{}){
	v := reflect.TypeOf(x)
	//类型名称
	  //种类
	fmt.Printf("类型：%v，类型名称：%v，类型种类: %v\n",v,v.Name(), v.Kind())
}
/* 
类型：int，类型名称：int，类型种类: int
类型：float64，类型名称：float64，类型种类: float64
类型：bool，类型名称：bool，类型种类: bool
类型：string，类型名称：string，类型种类: string
类型：main.myInt，类型名称：myInt，类型种类: int
类型：main.Person，类型名称：Person，类型种类: struct */

func main() {
	a := 10
	b := 11.2
	c := true
	d := "你好golang"
	reflectFn(a)
	reflectFn(b)
	reflectFn(c)
	reflectFn(d)

/* 	
	int
	float64
	bool
	string */

	var e myInt = 34
	var f = Person{
		Name : "张三",
		Age : 22,
	}
	reflectFn(e)
	reflectFn(f)
	
/* 	
main.myInt
main.Person */

}