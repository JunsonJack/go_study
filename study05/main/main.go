package main

import "fmt"

// 定义一个接口
type Phone interface {
	call()
}
// 定义一个结构体--小米
type XiaoMi struct {

}
// 实现接口方法
func (redmi XiaoMi) call(){
	fmt.Println("are you ok?")
}
// 定义一个结构体 -- 苹果
type IPhone struct{

}
// 实现接口方法
func (iphone12 IPhone) call(){
	fmt.Println("iphone12 pro max .....")
}

func main() {
	// 定义Phone类型的变量
	var myphone Phone

	myphone = new(XiaoMi)
	myphone.call()

	myphone = new(IPhone)
	myphone.call()
/*
	are you ok?	
	iphone12 pro max .....
*/


}