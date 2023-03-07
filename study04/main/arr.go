package main

import "fmt"
func main() {
	var arr [3] int16
	fmt.Println(len(arr))
	fmt.Println(arr)
	fmt.Printf("arr地址为：%p",&arr) //&取数组的地址
	fmt.Println()
	fmt.Printf("arr[0]地址为：%p",&arr[0])
	/*
			3
			[0 0 0]
			arr地址为：0xc0000120a0
			arr[0]地址为：0xc0000120a0
	*/

}