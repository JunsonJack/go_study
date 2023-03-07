package main

import "fmt"

// 数组的切片
func test (){
	//数组切片同样是前包含后不包含
	var arr = [...]int{1,2,3,4,5,6}
	a1 := arr[:3]
	fmt.Printf("a1: %v\n", a1)
	a2 := arr[3:]
	fmt.Printf("a2: %v\n", a2)
	a3 := arr[:]
	fmt.Printf("a3: %v\n", a3)

}

func main() {
	// 定义切片 slice
	var a []int                      //去除数组长度就是切片
	a = append(a, 1)                 //追加一个元素
	a = append(a, 2, 3, 3)           //追加多个元素
	a = append(a, []int{1, 2, 3}...) //追加一个切片，切片需要解包
	fmt.Printf("a: %v\n", a) //a: [1 2 3 3 1 2 3]

	var s1 = []int{1,2,3,4,5,6}
	s2 := s1[0:3]   //起始下标算，终止下标不算
	fmt.Printf("s2: %v\n", s2) //s2: [1 2 3]
	s3 := s1[3:]
	fmt.Printf("s3: %v\n", s3)  //s3: [4 5 6]

	// 数组切片
	test()
	/* 
	a1: [1 2 3]
	a2: [4 5 6]
	a3: [1 2 3 4 5 6] 
	*/

}