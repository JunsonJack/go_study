package main

import (
	"fmt"
	"gocode/study02/utils"
)

// 自定义函数
func sum(a int , b int )(int){
	var total int = 0
	total = (a + b)*2
	return total
}
/*
	函数首字母大写时，该函数可以在本包和其他包使用 类似java的public
	函数首字母小写时，该函数只能在本包使用  类似java的private
*/

func main() {
	
	fmt.Println("sum",sum(2,3)) // sum  10

	var num int= utils.AddNumber(3,5)
	fmt.Println("num",num) //num 8
}