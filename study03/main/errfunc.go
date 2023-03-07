package main

import (
	"errors"
	"fmt"
)

func main() {
	err := test()
	if err != nil {
		fmt.Println("自定义错误",err)
		panic(err) //遇见错误，终止go程序
	}
	fmt.Println("上面的触发操作执行成功")
	fmt.Println("正常执行下面的逻辑")
}
// 自定义异常处理
func test() (err error){
	num1 := 10
	num2 := 0
	if num2 == 0 {
		// 抛出自定义错误
		return errors.New("除数不能为0")
	}else {
		result := num1 / num2
		fmt.Println(result)
		// 如果没有错误，返回零值
		return nil
	}
}

/*
	goroutine 1 [running]:
	main.main()
					D:/goProject/src/gocode/study03/main/errfunc.go:12 +0x1bd
	exit status 2
*/